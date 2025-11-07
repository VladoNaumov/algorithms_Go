package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Normalize : приводим к нижнему регистру, убираем пробелы и все символы, которые не являются буквами или цифрами.
func Normalize(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
		// игнорируем пробелы, знаки пунктуации и т.д.
	}
	return b.String()
}

// IsAnagramCount Метод 1: сравнение по частотам символов (O(n) время)
func IsAnagramCount(a, b string) bool {
	na := Normalize(a)
	nb := Normalize(b)

	if len(na) != len(nb) {
		return false
	}

	// Используем map[rune]int, т.к. работаем с Unicode
	count := make(map[rune]int)
	for _, r := range na {
		count[r]++
	}
	for _, r := range nb {
		if count[r] == 0 {
			return false
		}
		count[r]--
	}
	// все должны быть нулевыми
	return true
}

// IsAnagramSort Метод 2: сортировка символов и сравнение (просто, но O(n log n))
func IsAnagramSort(a, b string) bool {
	na := Normalize(a)
	nb := Normalize(b)

	if len(na) != len(nb) {
		return false
	}

	ra := []rune(na)
	rb := []rune(nb)
	sort.Slice(ra, func(i, j int) bool { return ra[i] < ra[j] })
	sort.Slice(rb, func(i, j int) bool { return rb[i] < rb[j] })

	// сравниваем отсортированные слайсы рун
	if len(ra) != len(rb) {
		return false
	}
	for i := range ra {
		if ra[i] != rb[i] {
			return false
		}
	}
	return true
}

func main() {
	tests := [][2]string{
		{"listen", "silent"},
		{"Astronomer", "Moon starer"},
		{"Hello", "Olelh"},
		{"", ""},
		{"a", "a"},
		{"a  b", "b a"},
		{"ул. Ленина 10", "10 Ленина ул."},
		{"not anagram", "definitely not"},
	}

	fmt.Println("=== Проверка по подсчёту частот ===")
	for _, t := range tests {
		fmt.Printf("%q vs %q -> %v\n", t[0], t[1], IsAnagramCount(t[0], t[1]))
	}

	fmt.Println("\n=== Проверка через сортировку ===")
	for _, t := range tests {
		fmt.Printf("%q vs %q -> %v\n", t[0], t[1], IsAnagramSort(t[0], t[1]))
	}
}
