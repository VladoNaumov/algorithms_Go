package main

import (
	"fmt"
	"strings"
)

// countWords считает количество слов в строке
func countWords(text string) int {
	// 1. Удаляем знаки препинания
	punctuations := []string{".", ",", "!", "?", ";", ":", "\"", "'", "(", ")", "-", "—"}
	for _, p := range punctuations {
		text = strings.ReplaceAll(text, p, "")
	}

	// 2. Разделяем по пробелам
	words := strings.Fields(text)

	// 3. Возвращаем количество непустых слов
	return len(words)
}

func main() {
	tests := []string{
		"Привет, мир! Это тест.",
		"  Один, два,   три! ",
		"Hello, world!",
		"",
		"Только-тире—и—точки...",
	}

	for _, t := range tests {
		fmt.Printf("📘 \"%s\" → %d слов(а)\n", t, countWords(t))
	}
}
