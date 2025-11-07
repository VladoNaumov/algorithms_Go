package main

import (
	"fmt"
	"strings"
	"unicode"
)

// longestWord возвращает самое длинное слово и его длину.
func longestWord(sentence string) (string, int) {
	// 1️ Убираем пунктуацию
	clean := make([]rune, 0)
	for _, r := range sentence {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			clean = append(clean, r)
		}
	}

	// 2️ Разбиваем по пробелам
	words := strings.Fields(string(clean))

	longest := ""
	maxLen := 0

	// 3️ Перебираем слова
	for _, word := range words {
		length := len([]rune(word)) // учитываем Unicode
		if length > maxLen {
			maxLen = length
			longest = word
		}
	}

	return longest, maxLen
}

func main() {
	examples := []string{
		"Сегодня прекрасный солнечный день!",
		"Fly high, dream big!",
		"Go, go, gophers!",
		"Короткий тест.",
	}

	for _, text := range examples {
		word, length := longestWord(text)
		fmt.Printf("«%s» → самое длинное слово: %q (%d букв)\n", text, word, length)
	}
}
