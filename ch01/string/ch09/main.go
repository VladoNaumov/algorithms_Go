package main

import (
	"fmt"
	"strings"
)

func isPalindromePermutation(s string) bool {
	// 1. Убираем пробелы и приводим всё к нижнему регистру
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// 2. Создаём карту частот
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	// 3. Считаем количество символов с нечётной частотой
	oddCount := 0
	for _, count := range freq {
		if count%2 != 0 {
			oddCount++
		}
	}

	// 4. Если нечётных символов больше одного — нельзя сделать палиндром
	return oddCount <= 1
}

func main() {
	tests := []string{
		"carrace",                     // можно -> "racecar"
		"daily",                       // нельзя
		"aabb",                        // можно -> "abba"
		"A man a plan a canal Panama", // можно, если убрать пробелы
		"",                            // пустая строка — палиндром
	}

	for _, word := range tests {
		result := isPalindromePermutation(word)
		fmt.Printf("%q → %v\n", word, result)
	}
}
