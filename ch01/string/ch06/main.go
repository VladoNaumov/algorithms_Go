package main

import (
	"fmt"
)

// IsIsomorphic проверяет, изоморфны ли строки a и b.
// Алгоритм:
// - если длины разные — false
// - используем две map: mapAB (a->b) и mapBA (b->a)
// - идём по символам: проверяем существующие соответствия и создаём новые
func IsIsomorphic(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	mapAB := make(map[rune]rune)
	mapBA := make(map[rune]rune)

	ra := []rune(a)
	rb := []rune(b)

	for i := 0; i < len(ra); i++ {
		x := ra[i]
		y := rb[i]

		if v, ok := mapAB[x]; ok {
			if v != y {
				return false
			}
		} else {
			mapAB[x] = y
		}

		if v, ok := mapBA[y]; ok {
			if v != x {
				return false
			}
		} else {
			mapBA[y] = x
		}
	}
	return true
}

func main() {
	tests := [][2]string{
		{"foo", "app"},
		{"bar", "foo"},
		{"paper", "title"},
		{"egg", "add"},
		{"", ""},
		{"a", "b"},
		{"ab", "aa"},
		{"ключ-123", "кодо-456"}, // пример с Unicode и символами
	}

	for _, t := range tests {
		a, b := t[0], t[1]
		fmt.Printf("%q vs %q -> %v\n", a, b, IsIsomorphic(a, b))
	}
}
