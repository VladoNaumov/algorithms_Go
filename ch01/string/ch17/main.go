package main

import (
	"fmt"
)

// removeDuplicates удаляет повторяющиеся символы из строки,
func removeDuplicates(s string) string {
	seen := make(map[rune]bool)
	result := make([]rune, 0)

	for _, ch := range s {
		if !seen[ch] {
			seen[ch] = true
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	examples := []string{
		"banana",
		"abracadabra",
		"mississippi",
		"hello world",
		"1122334455",
	}

	for _, s := range examples {
		fmt.Printf("%s → %s\n", s, removeDuplicates(s))
	}
}
