package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isValidFileNameSimple(name string) bool {
	if !strings.HasPrefix(name, "file_") || !strings.HasSuffix(name, ".txt") {
		return false
	}

	// Вырезаем часть между "file_" и ".txt"
	core := strings.TrimSuffix(strings.TrimPrefix(name, "file_"), ".txt")

	// Проверяем, что ровно 3 символа и все — цифры
	if len(core) != 3 {
		return false
	}
	for _, ch := range core {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isValidFileNameSimple("file_123.txt")) // true
	fmt.Println(isValidFileNameSimple("file_12.txt"))  // false
	fmt.Println(isValidFileNameSimple("file_abc.txt")) // false
}
