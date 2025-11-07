package main

import (
	"fmt"
)

// IsBalanced проверяет, правильно ли закрыты скобки () [] {} в строке.
// Игнорирует все символы, не являющиеся скобками.
func IsBalanced(s string) bool {
	// стек для открывающих скобок
	stack := make([]rune, 0, len(s))

	// соответствие закрывающей -> открывающей
	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			// push
			stack = append(stack, r)
		case ')', ']', '}':
			// если стек пуст — сразу неправильно
			if len(stack) == 0 {
				return false
			}
			// смотрим верх стека
			top := stack[len(stack)-1]
			// если не совпадает по типу — неправильно
			if top != match[r] {
				return false
			}
			// pop
			stack = stack[:len(stack)-1]
		default:
			// другие символы игнорируем
		}
	}

	// корректно, если в стеке ничего не осталось
	return len(stack) == 0
}

func main() {
	tests := []string{
		"({[]})",
		"(]",
		"(()",
		"",                 // пустая строка
		"no brackets here", // без скобок
		"[({})]{}()",
		"{[)]}",                           // неправильное вложение
		"func(a, b) { return [1, 2, 3] }", // кодоподобная строка
	}

	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, IsBalanced(t))
	}
}
