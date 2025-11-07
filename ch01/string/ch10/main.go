package main

import (
	"fmt"
	"strings"
)

// SimplifyPath упрощает Unix-путь:
func SimplifyPath(path string) string {
	if path == "" {
		return "."
	}

	isAbs := strings.HasPrefix(path, "/")
	parts := strings.Split(path, "/")

	stack := make([]string, 0, len(parts))

	for _, p := range parts {
		if p == "" || p == "." {
			// пропускаем пустые компоненты и "."
			continue
		}
		if p == ".." {
			if len(stack) > 0 {
				// если в стеке есть компонент — попаем его
				// (это корректно как для абсолютного, так и для относительного пути)
				stack = stack[:len(stack)-1]
			} else {
				// стек пуст
				if isAbs {
					// для абсолютного пути нельзя подниматься выше корня — пропускаем
					continue
				} else {
					// для относительного пути сохраняем ".." в стеке
					stack = append(stack, "..")
				}
			}
		} else {
			// обычный компонент — добавляем
			stack = append(stack, p)
		}
	}

	// собираем результат
	if isAbs {
		if len(stack) == 0 {
			return "/"
		}
		return "/" + strings.Join(stack, "/")
	} else {
		if len(stack) == 0 {
			return "."
		}
		return strings.Join(stack, "/")
	}
}

func main() {
	tests := []string{
		"/a/./b/../c/",
		"/../",               // попытка подняться выше корня
		"/home//foo/",        // множественные слэши
		"a/b/../../c",        // относительный: поднимается на уровень
		"../../a/b",          // относительный с ведущими ..
		"",                   // пустая строка
		"./././",             // текущая директория
		"//a///b//./c/..//d", // комплексный пример
		"/a/b/c/../../..",    // абсолютный возвращается в корень
		"../../..",           // только ведущие .. в относительном
	}

	for _, p := range tests {
		fmt.Printf("in: %-25q -> out: %q\n", p, SimplifyPath(p))
	}
}
