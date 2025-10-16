package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

// ==== Константы конфигурации ====

const (
	// Сетевые настройки
	DefaultAddr       = ":8080"
	ReadHeaderTimeout = 5 * time.Second
	ReadBodyTimeout   = 10 * time.Second
	WriteTimeout      = 20 * time.Second
	IdleTimeout       = 60 * time.Second
	MaxHeaderBytes    = 1 << 20
	MaxBodyBytes      = 10 << 20

	// Безопасность
	AllowedHosts         = "localhost:8080,127.0.0.1:8080" // CSV
	AllowedOrigins       = "http://localhost:8080"         // CSV
	RateLimitMaxRequests = 100
	RateLimitWindow      = 1 * time.Minute
	SessionTokenLength   = 32
	SessionMaxAge        = 3600 // 1 час

	// TLS
	UseTLS      = false // true для продакшена
	TLSCertFile = ""    // /path/to/cert.pem
	TLSKeyFile  = ""    // /path/to/key.pem

	// Файлы
	MaxUploadFileMB         = 10
	AllowedUploadExtensions = ".png,.jpg,.jpeg,.gif"
)

// ==== Конфигурация ====

type Config struct {
	Addr              string
	AllowedHosts      []string
	AllowedOrigins    []string
	MaxHeaderBytes    int
	MaxBodyBytes      int64
	ReadHeaderTimeout time.Duration
	ReadBodyTimeout   time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	UseTLS            bool
	CertFile          string
	KeyFile           string
	RateLimitMax      int
	RateLimitWindow   time.Duration
}

func LoadConfig() Config {
	return Config{
		Addr:              DefaultAddr,
		AllowedHosts:      splitCSV(AllowedHosts),
		AllowedOrigins:    splitCSV(AllowedOrigins),
		MaxHeaderBytes:    MaxHeaderBytes,
		MaxBodyBytes:      MaxBodyBytes,
		ReadHeaderTimeout: ReadHeaderTimeout,
		ReadBodyTimeout:   ReadBodyTimeout,
		WriteTimeout:      WriteTimeout,
		IdleTimeout:       IdleTimeout,
		UseTLS:            UseTLS,
		CertFile:          TLSCertFile,
		KeyFile:           TLSKeyFile,
		RateLimitMax:      RateLimitMaxRequests,
		RateLimitWindow:   RateLimitWindow,
	}
}

func splitCSV(s string) []string {
	var out []string
	for _, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// ==== Rate Limiter (in-memory, per IP) ====

type rateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	max      int
	window   time.Duration
}

func newRateLimiter(max int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		requests: make(map[string][]time.Time),
		max:      max,
		window:   window,
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	reqs := rl.requests[ip]

	// Очистка старых записей
	reqs = filterRecent(reqs, now, rl.window)

	if len(reqs) >= rl.max {
		return false
	}

	reqs = append(reqs, now)
	rl.requests[ip] = reqs
	return true
}

func filterRecent(times []time.Time, now time.Time, window time.Duration) []time.Time {
	var recent []time.Time
	cutoff := now.Add(-window)
	for _, t := range times {
		if t.After(cutoff) {
			recent = append(recent, t)
		}
	}
	return recent
}

// ==== Утилиты ====

// clientIP извлекает IP клиента (простая версия, без XFF для безопасности)
func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// randomToken генерирует криптографически стойкий токен
func randomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// validateOrigin строго проверяет Origin/Referer против вайтлиста
func validateOrigin(allowed []string, originStr string) bool {
	u, err := url.Parse(originStr)
	if err != nil {
		return false
	}

	for _, allowedStr := range allowed {
		au, err := url.Parse(allowedStr)
		if err != nil {
			continue
		}
		if u.Scheme == au.Scheme && u.Host == au.Host {
			return true
		}
	}
	return false
}

// isStateChanging определяет методы, меняющие состояние
func isStateChanging(method string) bool {
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		return true
	}
	return false
}

// ==== Middleware ====

type middleware func(http.Handler) http.Handler

// secureHeaders устанавливает безопасные HTTP заголовки
func secureHeaders(cfg Config) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
			w.Header().Set("Content-Security-Policy", "default-src 'self'; frame-ancestors 'none'")

			if r.TLS != nil {
				w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
			}
			next.ServeHTTP(w, r)
		})
	}
}

// limitBody ограничивает размер тела запроса
func limitBody(max int64) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Body != nil {
				r.Body = http.MaxBytesReader(w, r.Body, max)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// hostWhitelist проверяет Host против вайтлиста
func hostWhitelist(allowed []string) middleware {
	allow := make(map[string]struct{}, len(allowed))
	for _, h := range allowed {
		allow[strings.ToLower(h)] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := allow[strings.ToLower(r.Host)]; !ok {
				http.Error(w, "invalid host", http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// csrfOriginGuard проверяет Origin/Referer для state-changing запросов
func csrfOriginGuard(allowedOrigins []string) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isStateChanging(r.Method) {
				next.ServeHTTP(w, r)
				return
			}

			origin := r.Header.Get("Origin")
			referer := r.Header.Get("Referer")

			// Проверяем Origin
			if origin != "" && validateOrigin(allowedOrigins, origin) {
				next.ServeHTTP(w, r)
				return
			}

			// Fallback на Referer (менее строго)
			if referer != "" && validateOrigin(allowedOrigins, referer) {
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, "CSRF check failed", http.StatusForbidden)
		})
	}
}

// corsStrict реализует строгий CORS
func corsStrict(allowedOrigins []string) middleware {
	origins := make(map[string]struct{}, len(allowedOrigins))
	for _, o := range allowedOrigins {
		origins[o] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" {
				if _, ok := origins[origin]; ok {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Vary", "Origin")
					w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,X-CSRF-Token")
					w.Header().Set("Access-Control-Allow-Credentials", "true")

					if r.Method == http.MethodOptions {
						w.WriteHeader(http.StatusNoContent)
						return
					}
				} else if r.Method == http.MethodOptions {
					http.Error(w, "CORS not allowed", http.StatusForbidden)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

// recoverer ловит паники
func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic recovered: %v", rec)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// requestLogger логирует запросы без чувствительных данных
func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := clientIP(r)

		next.ServeHTTP(w, r)

		status := "200"
		if tw, ok := w.(*responseWriter); ok {
			status = fmt.Sprintf("%d", tw.status)
		}

		log.Printf("%s %s %s %s ip=%s ua=%q status=%s dur=%v",
			r.Method, r.URL.Path, r.Proto, status, ip,
			r.UserAgent(), time.Since(start))
	})
}

// rateLimit ограничивает запросы по IP
func rateLimit(rl *rateLimiter) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r)
			if !rl.allow(ip) {
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// responseWriter для логирования статус-кода
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// ==== Хэндлеры ====

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("ok"))
}

// uploadHandler безопасно обрабатывает загрузку файлов
func uploadHandler(maxMB int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверка CSRF токена в form или header
		csrfToken := r.Header.Get("X-CSRF-Token")
		if csrfToken == "" {
			csrfToken = r.FormValue("csrf_token")
		}
		if csrfToken == "" {
			http.Error(w, "CSRF token required", http.StatusForbidden)
			return
		}

		// Ограниченный парсинг multipart
		if err := r.ParseMultipartForm(maxMB << 20); err != nil {
			http.Error(w, "invalid form", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "file missing", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Валидация имени файла
		name := filepath.Base(header.Filename)
		if name == "." || strings.Contains(name, "..") || strings.ContainsAny(name, "/\\") {
			http.Error(w, "invalid filename", http.StatusBadRequest)
			return
		}

		// Валидация расширения
		ext := strings.ToLower(filepath.Ext(name))
		allowed := make(map[string]bool)
		for _, e := range strings.Split(AllowedUploadExtensions, ",") {
			allowed[strings.TrimSpace(e)] = true
		}
		if !allowed[ext] {
			http.Error(w, "unsupported file type", http.StatusUnsupportedMediaType)
			return
		}

		// Здесь: проверка MIME-type, сигнатур, антивирус и т.д.
		// Для примера просто подтверждаем
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"status":"ok","filename":"%s","size":%d}`,
			name, header.Size)))
	}
}

// loginHandler устанавливает защищенную сессию
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// В реальности: валидация credentials
	token, err := randomToken(SessionTokenLength / 2)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		MaxAge:   SessionMaxAge,
		Secure:   r.TLS != nil,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

	// Возвращаем CSRF токен для последующих запросов
	csrfToken, _ := randomToken(16)
	w.Header().Set("X-CSRF-Token", csrfToken)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","csrf_token":"` + csrfToken + `"}`))
}

// ==== Сборка стека ====

func chain(h http.Handler, m ...middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

// tlsConfig безопасные настройки TLS + HTTP/2
func tlsConfig() *tls.Config {
	return &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		NextProtos:               []string{"h2", "http/1.1"},
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},
	}
}

// validateTLSFiles проверяет наличие TLS сертификатов
func validateTLSFiles(certFile, keyFile string) error {
	if certFile == "" || keyFile == "" {
		return fmt.Errorf("TLS enabled but cert/key files not specified")
	}

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		return fmt.Errorf("TLS cert file not found: %s", certFile)
	}
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		return fmt.Errorf("TLS key file not found: %s", keyFile)
	}
	return nil
}

// ==== Main ====

func main() {
	cfg := LoadConfig()

	// Валидация TLS
	if cfg.UseTLS {
		if err := validateTLSFiles(cfg.CertFile, cfg.KeyFile); err != nil {
			log.Fatal(err)
		}
	}

	// Rate limiter
	rl := newRateLimiter(cfg.RateLimitMax, cfg.RateLimitWindow)

	// Роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)
	mux.Handle("/api/upload", uploadHandler(MaxUploadFileMB))
	mux.HandleFunc("/api/login", loginHandler)

	// Стек middleware (порядок важен!)
	handler := chain(
		mux,
		requestLogger,      // логирование
		recoverer,          // восстановление от паник
		rateLimit(rl),      // ограничение скорости
		secureHeaders(cfg), // безопасные заголовки
		hostWhitelist(cfg.AllowedHosts),
		corsStrict(cfg.AllowedOrigins),
		csrfOriginGuard(cfg.AllowedOrigins),
		limitBody(cfg.MaxBodyBytes),
	)

	// Сервер с таймаутами
	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}

	if cfg.UseTLS {
		srv.TLSConfig = tlsConfig()
	}

	// Graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		log.Println("shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("server shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("starting server on %s (tls=%v)", cfg.Addr, cfg.UseTLS)
	var err error
	if cfg.UseTLS {
		err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	<-idleConnsClosed
	log.Println("server stopped")
}
