package main

import (
	"fmt"
	"strings"
)

// normalizeSpaces удаляет лишние пробелы в строке:
// - убирает ведущие и конечные пробелы,
// - заменяет несколько пробелов подряд на один,
// - не трогает символы, кроме пробелов.
func normalizeSpaces(s string) string {
	parts := strings.Fields(s)
	return strings.Join(parts, " ")
}

func main() {
	tests := []string{
		"  Привет   мир   !  ",
		"   Один    два    три   ",
		"Без   лишних  пробелов",
		"",
		"   ",
	}

	for _, t := range tests {
		fmt.Printf("Исходная: [%s]\n→ Нормализовано: [%s]\n\n", t, normalizeSpaces(t))
	}
}
