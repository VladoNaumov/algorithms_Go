package main

import (
	"fmt"
	"strings"
)

// ExcelColumnToNumber преобразует буквенное обозначение столбца (например "AA") в число.
func ExcelColumnToNumber(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("пустая строка недопустима")
	}

	s = strings.ToUpper(s) // приведение к верхнему регистру
	result := 0

	for _, ch := range s {
		if ch < 'A' || ch > 'Z' {
			return 0, fmt.Errorf("недопустимый символ: %q", ch)
		}
		val := int(ch-'A') + 1
		result = result*26 + val
	}

	return result, nil
}

func main() {
	tests := []string{
		"A", "Z", "AA", "AZ", "BA", "ZZ", "AAA", "XFD", "abc", "",
	}

	for _, t := range tests {
		n, err := ExcelColumnToNumber(t)
		if err != nil {
			fmt.Printf("%-5q -> error: %v\n", t, err)
		} else {
			fmt.Printf("%-5q -> %d\n", t, n)
		}
	}
}
