package main

import (
	"fmt"
)

// MostFrequentChar возвращает символ с максимальной частотой и саму частоту
func MostFrequentChar(s string) (rune, int, bool) {
	if len(s) == 0 {
		return 0, 0, false // пустая строка
	}

	counts := make(map[rune]int)

	// 1. считаем частоты
	for _, ch := range s {
		counts[ch]++
	}

	// 2. ищем максимум
	var maxChar rune
	maxCount := 0
	for ch, c := range counts {
		if c > maxCount {
			maxChar = ch
			maxCount = c
		}
	}

	return maxChar, maxCount, true
}

func main() {
	tests := []string{
		"abbccc",
		"a a b",
		"Hello, World!",
		"аааБбвввв", // пример с кириллицей
		"",
	}

	for _, s := range tests {
		ch, count, ok := MostFrequentChar(s)
		if !ok {
			fmt.Printf("%q -> пустая строка\n", s)
			continue
		}
		fmt.Printf("%q -> символ %q встречается %d раз(а)\n", s, ch, count)
	}
}
