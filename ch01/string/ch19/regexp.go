package main

import (
	"fmt"
	"regexp"
)

// isValidEmailRegex — проверка e-mail через регулярное выражение
func isValidEmailRegex(email string) bool {
	// Примерное правило: имя может содержать буквы, цифры, точки, дефисы, подчёркивания;
	// домен — буквы, цифры, дефисы; TLD — минимум 2 буквы
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)
	return re.MatchString(email)
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
		"test@exam_ple.com",
		"user@domain.c",
		"user@domain.co.uk",
	}

	for _, e := range emails {
		fmt.Printf("%-25s → %v\n", e, isValidEmailRegex(e))
	}
}
