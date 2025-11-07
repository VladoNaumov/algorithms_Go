package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func CompressRLE(s string) string {
	if len(s) == 0 {
		return ""
	}

	var b strings.Builder
	count := 1
	prev := rune(s[0])

	for _, ch := range s[1:] {
		if ch == prev {
			count++
		} else {
			b.WriteRune(prev)
			b.WriteString(strconv.Itoa(count))
			prev = ch
			count = 1
		}
	}
	// последний блок
	b.WriteRune(prev)
	b.WriteString(strconv.Itoa(count))

	return b.String()
}

// DecompressRLE восстанавливает строку из RLE.
func DecompressRLE(s string) (string, error) {
	var b strings.Builder
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; {

		if !unicode.IsLetter(runes[i]) {
			return "", fmt.Errorf("некорректный формат: ожидается символ, найдено %q", runes[i])
		}
		char := runes[i]
		i++

		if i >= n || !unicode.IsDigit(runes[i]) {
			return "", fmt.Errorf("некорректный формат: после %q отсутствует число", char)
		}

		start := i
		for i < n && unicode.IsDigit(runes[i]) {
			i++
		}
		countStr := string(runes[start:i])
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return "", fmt.Errorf("ошибка чтения числа: %v", err)
		}

		for j := 0; j < count; j++ {
			b.WriteRune(char)
		}
	}

	return b.String(), nil
}

func main() {
	examples := []string{
		"aaabbcddd",
		"aaaaa",
		"abbbbbccaa",
		"",
	}

	for _, s := range examples {
		compressed := CompressRLE(s)
		decompressed, err := DecompressRLE(compressed)

		fmt.Printf("Исходная: %-10q → Сжато: %-10q → Восстановлено: %-10q", s, compressed, decompressed)
		if err != nil {
			fmt.Printf(" (ошибка: %v)", err)
		}
		fmt.Println()
	}

	// Пример некорректного формата
	fmt.Println("\nПроверка ошибок:")
	bad := "a3bX"
	_, err := DecompressRLE(bad)
	fmt.Printf("Ввод: %q → Ошибка: %v\n", bad, err)
}
