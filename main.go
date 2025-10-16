package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// ==== Конфиг (упростим: читаем из ENV при старте) ====

type Config struct {
	Addr              string        // :8080 или :443
	AllowedHosts      []string      // example.com, api.example.com:443
	AllowedOrigins    []string      // https://example.com, https://app.example.com
	MaxHeaderBytes    int           // напр. 1 << 20
	MaxBodyBytes      int64         // напр. 10 << 20 (10MB)
	ReadHeaderTimeout time.Duration // 5s
	ReadTimeout       time.Duration // 10s
	WriteTimeout      time.Duration // 20s
	IdleTimeout       time.Duration // 60s
	UseTLS            bool          // если true — ждём TLS (сертификаты задаются отдельно)
	CertFile          string
	KeyFile           string
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func loadConfig() Config {
	// Для простоты — часть значений захардкожена,
	// в реальном проекте — парсить ENV/файл.
	return Config{
		Addr:              getenv("ADDR", ":8080"),
		AllowedHosts:      splitCSV(getenv("ALLOWED_HOSTS", "localhost:8080,127.0.0.1:8080")),
		AllowedOrigins:    splitCSV(getenv("ALLOWED_ORIGINS", "http://localhost:8080")),
		MaxHeaderBytes:    1 << 20,
		MaxBodyBytes:      10 << 20,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
		UseTLS:            strings.ToLower(getenv("USE_TLS", "false")) == "true",
		CertFile:          getenv("TLS_CERT_FILE", ""),
		KeyFile:           getenv("TLS_KEY_FILE", ""),
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

// ==== Утилиты ====

func randomToken(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

// ==== Middleware ====

type middleware func(http.Handler) http.Handler

// 1) Безопасные заголовки + строгий HSTS при TLS
func secureHeaders(cfg Config) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("Referrer-Policy", "no-referrer")
			// Минимальная CSP: разрешаем только свой домен
			w.Header().Set("Content-Security-Policy", "default-src 'self'; frame-ancestors 'none'")
			if r.TLS != nil {
				// 2 года, включая поддомены
				w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
			}
			next.ServeHTTP(w, r)
		})
	}
}

// 2) Ограничение тела запроса
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

// 3) Валидация Host (защита от Host header attacks)
func hostWhitelist(allowed []string) middleware {
	allow := make(map[string]struct{}, len(allowed))
	for _, h := range allowed {
		allow[strings.ToLower(h)] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host := strings.ToLower(r.Host)
			if _, ok := allow[host]; !ok {
				http.Error(w, "invalid host", http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// 4) Простой CSRF-щит: для методов, меняющих состояние, проверяем Origin/Referer против вайтлиста
func csrfOriginGuard(allowedOrigins []string) middleware {
	allow := make(map[string]struct{}, len(allowedOrigins))
	for _, o := range allowedOrigins {
		allow[o] = struct{}{}
	}
	isStateChanging := func(m string) bool {
		switch m {
		case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
			return true
		default:
			return false
		}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isStateChanging(r.Method) {
				next.ServeHTTP(w, r)
				return
			}
			origin := r.Header.Get("Origin")
			ref := r.Header.Get("Referer")
			ok := false
			if origin != "" {
				_, ok = allow[origin]
			}
			if !ok && ref != "" {
				// Проверяем префикс (полный URL строго не требуется)
				for o := range allow {
					if strings.HasPrefix(ref, o) {
						ok = true
						break
					}
				}
			}
			if !ok {
				http.Error(w, "csrf check failed", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// 5) Простой CORS (строгий)
func corsStrict(allowedOrigins []string) middleware {
	allow := make(map[string]struct{}, len(allowedOrigins))
	for _, o := range allowedOrigins {
		allow[o] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" {
				if _, ok := allow[origin]; ok {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Vary", "Origin")
					w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					if r.Method == http.MethodOptions {
						w.WriteHeader(http.StatusNoContent)
						return
					}
				} else {
					// Неизвестный Origin — блокируем preflight
					if r.Method == http.MethodOptions {
						http.Error(w, "cors not allowed", http.StatusForbidden)
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

// 6) Восстановление после паники
func recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v", rec)
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// 7) Простой логгер (без тела и без секретов)
func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := clientIP(r)
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s ip=%s ua=%q dur=%s",
			r.Method, r.URL.Path, r.Proto, ip, r.UserAgent(), time.Since(start))
	})
}

func clientIP(r *http.Request) string {
	// Не доверяем XFF по умолчанию. Если есть доверенный прокси — добавьте явную обработку.
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// ==== Роуты / хэндлеры ====

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok"))
}

func uploadHandler(maxFileMB int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Пример безопасного разбора multipart: часть в память, остальное — во временный файл.
		if err := r.ParseMultipartForm(maxFileMB << 20); err != nil {
			http.Error(w, "bad form", http.StatusBadRequest)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "file missing", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Простейшая проверка имени (без ../) + ограничение расширений
		name := filepath.Base(header.Filename)
		if name == "." || strings.Contains(name, "..") {
			http.Error(w, "bad filename", http.StatusBadRequest)
			return
		}
		lname := strings.ToLower(name)
		if !(strings.HasSuffix(lname, ".png") || strings.HasSuffix(lname, ".jpg") || strings.HasSuffix(lname, ".jpeg")) {
			http.Error(w, "unsupported file type", http.StatusUnsupportedMediaType)
			return
		}

		// Здесь вы бы читали ограниченно (io.CopyN) и проверяли сигнатуры, размер и т.д.
		// Для примера просто подтверждаем.
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"uploaded","name":"` + name + `"}`))
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Демонстрация установки защищённой cookie (НЕ храните в cookie чувствительные данные!)
	token := randomToken(16)
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		Secure:   r.TLS != nil,         // только по HTTPS
		HttpOnly: true,                 // недоступна JS
		SameSite: http.SameSiteLaxMode, // разумно по умолчанию
	})
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

// ==== Сборка middleware-цепочки ====

func chain(h http.Handler, m ...middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

// ==== TLS c безопасными настройками (минимум шума) ====

func tlsConfig() *tls.Config {
	return &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		// Серверные настройки по умолчанию в Go уже безопасные.
	}
}

// ==== main ====

func main() {
	cfg := loadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthHandler)
	mux.Handle("/api/upload", uploadHandler(10)) // 10MB на файл
	mux.HandleFunc("/api/login", loginHandler)

	// Глобальные middleware:
	handler := chain(
		mux,
		requestLogger,
		recoverer,
		secureHeaders(cfg),
		hostWhitelist(cfg.AllowedHosts),
		corsStrict(cfg.AllowedOrigins),
		csrfOriginGuard(cfg.AllowedOrigins),
		limitBody(cfg.MaxBodyBytes),
	)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           handler,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}

	// Грациозное завершение
	idleConnsClosed := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch
		log.Println("shutdown: stopping server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("listening on %s (tls=%v)", cfg.Addr, cfg.UseTLS)
	var err error
	if cfg.UseTLS {
		srv.TLSConfig = tlsConfig()
		err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
	} else {
		err = srv.ListenAndServe()
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
	<-idleConnsClosed
}
