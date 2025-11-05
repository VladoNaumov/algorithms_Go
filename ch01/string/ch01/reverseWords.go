package main

import (
	"fmt"
	"strings"
)

// ReverseWordsSimple 1. Разворот слов в адресе
func ReverseWordsSimple(s string) string {
	// strings.Fields разбивает строку по любым пробельным символам и автоматически убирает лишние пробелы.
	words := strings.Fields(s)

	// переворачиваем slice слов на месте
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// соединяем слова одним пробелом
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ReverseWordsSimple("ул. Ленина 10 кв. 5"))
	fmt.Println(ReverseWordsSimple("  ул.   Ленина 10   кв. 5  "))
	fmt.Println(ReverseWordsSimple("Москва, Suomi"))
	fmt.Println(ReverseWordsSimple("    .      .       .      "))
}
