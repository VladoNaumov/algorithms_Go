package main

import (
	"fmt"
	"strings"
	"unicode"
)

// isValidEmail проверяет базовую корректность email.
func isValidEmail(email string) bool {
	// 1 Не должно быть пробелов
	if strings.Contains(email, " ") {
		return false
	}

	// 2 Разделяем по '@'
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false // должно быть ровно одно '@'
	}

	local := parts[0]
	domain := parts[1]

	// 3️ Проверяем, что обе части непустые
	if local == "" || domain == "" {
		return false
	}

	// 4️ В домене должен быть хотя бы один '.'
	if !strings.Contains(domain, ".") {
		return false
	}

	// 5️ '.' не должен быть первым или последним символом домена
	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return false
	}

	// 6️ Проверка допустимых символов (упрощённая)
	for _, r := range email {
		if unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func main() {
	emails := []string{
		"user@example.com",
		"user.name@domain.org",
		"user@sub.domain.com",
		"user@@example.com",
		"@example.com",
		"user@.com",
		"user@domain",
		"user domain@example.com",
		" test@example.com",
		"user@example.com ",
	}

	for _, e := range emails {
		fmt.Printf("%-25s → %v\n", e, isValidEmail(e))
	}
}
