package main

import (
	"fmt"
	"sort"
)

// === Функция 1: Прямой способ ===
func LongestCommonPrefixDirect(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	for i := 0; i < len(first); i++ {
		ch := first[i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != ch {
				return first[:i]
			}
		}
	}
	return first
}

// === Функция 2: Через сортировку ===
func LongestCommonPrefixSort(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := make([]string, len(strs))
	copy(s, strs)
	sort.Strings(s)

	a := s[0]
	b := s[len(s)-1]
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[:i]
}

func main() {
	tests := [][]string{
		{"/home/user/docs/report.doc", "/home/user/docs/notes.txt", "/home/user/docs/"},
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"single"},
		{},
		{"", "anything"},
		{"prefixSame", "prefixSame", "prefixSame"},
	}

	fmt.Println("=== Проверка прямым способом ===")
	for _, t := range tests {
		fmt.Printf("Вход: %v\n→ Общий префикс: %q\n\n", t, LongestCommonPrefixDirect(t))
	}

	fmt.Println("=== Проверка через сортировку ===")
	for _, t := range tests {
		fmt.Printf("Вход: %v\n→ Общий префикс: %q\n\n", t, LongestCommonPrefixSort(t))
	}
}

/*
| Пример                                                                            | Объяснение результата                                                 |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `["/home/user/docs/report.doc", "/home/user/docs/notes.txt", "/home/user/docs/"]` | Все строки начинаются с `/home/user/docs/` — это и есть общий префикс |
| `["flower", "flow", "flight"]`                                                    | Общие первые буквы `fl`                                               |
| `["dog", "racecar", "car"]`                                                       | Нет общего начала — пустая строка                                     |
| `["single"]`                                                                      | Только одна строка — весь текст и есть общий префикс                  |
| `[]`                                                                              | Пустой список — пустой результат                                      |
| `["", "anything"]`                                                                | Первая строка пустая, значит общий префикс тоже пуст                  |
| `["prefixSame", "prefixSame", "prefixSame"]`                                      | Все строки одинаковые — общий префикс равен им полностью              |
*/
