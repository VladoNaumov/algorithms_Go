package main

import (
	"fmt"
	"strings"
)

// getDomain извлекает домен из URL
func getDomain(url string) string {
	// 1. Убираем протокол
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")

	// 2. Убираем "www."
	url = strings.TrimPrefix(url, "www.")

	// 3. Если есть путь — обрезаем всё после первого "/"
	if idx := strings.Index(url, "/"); idx != -1 {
		url = url[:idx]
	}

	// 4. Возвращаем только домен
	return url
}

func main() {
	tests := []string{
		"https://www.google.com/maps?q=test",
		"http://example.org/about",
		"https://my-site.net",
		"www.github.com/user/repo",
		"ftp://ftp.example.com/files", // даже нестандартный случай
	}

	for _, t := range tests {
		fmt.Printf("🔗 %s → 🌐 %s\n", t, getDomain(t))
	}
}
