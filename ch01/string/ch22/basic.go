package main

import (
	"fmt"
	"unicode"
)

// isStrongPassword проверяет, соответствует ли пароль 4 критериям:
// 1. Есть хотя бы одна строчная буква
// 2. Есть хотя бы одна заглавная буква
// 3. Есть хотя бы одна цифра
// 4. Есть хотя бы один спецсимвол
func isStrongPassword(password string) bool {
	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

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
		if isStrongPassword(t) {
			fmt.Printf("ok '%s' — надёжный пароль\n", t)
		} else {
			fmt.Printf("no '%s' — слабый пароль\n", t)
		}
	}
}
