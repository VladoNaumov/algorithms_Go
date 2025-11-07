package main

import (
	"fmt"
)

// RotN выполняет сдвиг букв на n позиций (по модулю 26)
func RotN(s string, n int) string {
	result := []rune{}

	for _, ch := range s {
		// латинские большие буквы
		if ch >= 'A' && ch <= 'Z' {
			rot := ((ch-'A')+rune(n))%26 + 'A'
			result = append(result, rot)
		} else if ch >= 'a' && ch <= 'z' { // латинские маленькие буквы
			rot := ((ch-'a')+rune(n))%26 + 'a'
			result = append(result, rot)
		} else {
			// оставляем символ без изменений (пробелы, цифры, знаки)
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	text := "Hello, World!"
	rot13 := RotN(text, 13)
	back := RotN(rot13, 13)

	fmt.Println("Оригинал: ", text)
	fmt.Println("ROT13:    ", rot13)
	fmt.Println("Двойной ROT13 (обратно):", back)
}
