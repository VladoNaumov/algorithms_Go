package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// IsValidIPv4 проверяет, является ли s валидным IPv4-адресом.
func IsValidIPv4(s string, allowLeadingZeros bool) bool {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return false
	}

	for _, p := range parts {
		if p == "" {
			return false
		}

		if len(p) > 3 {
			return false
		}

		if !allowLeadingZeros && len(p) > 1 && p[0] == '0' {
			return false
		}

		for _, r := range p {
			if !unicode.IsDigit(r) {
				return false
			}
		}

		num, err := strconv.Atoi(p)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}

	}

	return true
}

func main() {
	tests := []string{
		"192.168.1.1",
		"0.0.0.0",
		"255.255.255.255",
		"256.100.0.1",
		"192.168.1",
		"01.2.3.4",
		"1.1.1.01",
		"1..1.1",
		" 1.1.1.1 ",
		"1.1.1.1\n",
		"123.045.067.089",
		"1.1.1.1.1",
		"1.1.one.1",
	}

	fmt.Println("=== Запрет ведущих нулей (обычное поведение) ===")
	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, IsValidIPv4(t, false))
	}

	fmt.Println("\n=== Разрешены ведущие нули ===")
	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, IsValidIPv4(t, true))
	}
}
