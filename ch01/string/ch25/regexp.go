package main

import (
	"fmt"
	"regexp"
)

// isValidFileName проверяет, соответствует ли имя шаблону file_###.txt
func isValidFileName(name string) bool {
	re := regexp.MustCompile(`^file_\d{3}\.txt$`)
	return re.MatchString(name)
}

func main() {
	tests := []string{
		"file_123.txt",
		"file_12.txt",
		"file_abc.txt",
		"file_999.txt",
		"myfile_123.txt",
		"file_123.tx",
	}

	for _, t := range tests {
		fmt.Printf("%-15s → %v\n", t, isValidFileName(t))
	}
}
