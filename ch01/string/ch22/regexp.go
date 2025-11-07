package main

import (
	"fmt"
	"regexp"
)

func isStrongPasswordRegex(password string) bool {
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*\(\)\-_=\+\[\]\{\};:'",<\.>/?\\|]`).MatchString(password)

	return hasLower && hasUpper && hasDigit && hasSpecial
}

func main() {
	tests := []string{
		"Qwerty@123",
		"password123",
		"HELLO@2025",
		"Test123",
		"GoLang!1",
	}

	for _, t := range tests {
		if isStrongPasswordRegex(t) {
			fmt.Printf("ok '%s' — надёжный пароль\n", t)
		} else {
			fmt.Printf("no '%s' — слабый пароль\n", t)
		}
	}
}
