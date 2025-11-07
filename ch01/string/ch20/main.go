package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// interval хранит полуинтервал [start, end) в байтовых индексах исходной строки.
type interval struct {
	start int
	end   int
}

// highlightKeywords подсвечивает ключевые слова в тексте, оборачивая их в <b>...</b>.
func highlightKeywords(text string, keywords []string, caseInsensitive, wholeWord bool) string {
	if text == "" || len(keywords) == 0 {
		return text
	}

	// Подготовка исходного поискового текста (возможно, в нижнем регистре)
	searchText := text
	if caseInsensitive {
		searchText = strings.ToLower(text)
	}

	var intervals []interval

	for _, kw := range keywords {
		if kw == "" {
			continue
		}
		searchKW := kw
		if caseInsensitive {
			searchKW = strings.ToLower(kw)
		}

		// Поиск всех вхождений (включая пересекающиеся)
		start := 0
		for {
			idx := strings.Index(searchText[start:], searchKW)
			if idx == -1 {
				break
			}
			absStart := start + idx
			absEnd := absStart + len(searchKW) // байтовые индексы; корректны для slice(text[absStart:absEnd])

			// Если включена опция wholeWord, проверим границы:
			if wholeWord {
				if !isWordBoundary(text, absStart, absEnd) {
					// не целое слово — пропускаем это вхождение
					start = absStart + 1 // продвигаемся на 1 байт и ищем дальше (позволяет найти пересечения)
					continue
				}
			}

			intervals = append(intervals, interval{start: absStart, end: absEnd})

			// Продолжаем поиск начиная со следующего байта (чтобы позволить пересекающиеся вхождения)
			start = absStart + 1
			if start >= len(searchText) {
				break
			}
		}
	}

	if len(intervals) == 0 {
		return text
	}

	// Сортируем по start
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return intervals[i].start < intervals[j].start
	})

	// Сливаем пересечения/смежности: если next.start <= cur.end -> расширяем cur.end = max(cur.end, next.end)
	merged := make([]interval, 0, len(intervals))
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		ni := intervals[i]
		if ni.start <= cur.end {
			if ni.end > cur.end {
				cur.end = ni.end
			}
		} else {
			merged = append(merged, cur)
			cur = ni
		}
	}
	merged = append(merged, cur)

	// Строим итоговую строку, вставляя теги по байтовым индексам
	var b strings.Builder
	prev := 0
	for _, it := range merged {
		// добавляем промежуток между предыдущим и началом тега
		if prev < it.start {
			b.WriteString(text[prev:it.start])
		}
		// открывающий тег, сам текст, закрывающий тег
		b.WriteString("<b>")
		b.WriteString(text[it.start:it.end])
		b.WriteString("</b>")
		prev = it.end
	}
	// остаток
	if prev < len(text) {
		b.WriteString(text[prev:])
	}
	return b.String()
}

// isWordBoundary проверяет, что вхождение [start,end) является целым словом:
func isWordBoundary(s string, start, end int) bool {
	// Проверка левой границы
	if start > 0 {

		left := []rune(s[:start])
		if len(left) > 0 {
			r := left[len(left)-1]
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				return false
			}
		}
	}
	// Правая граница
	if end < len(s) {
		right := []rune(s[end:])
		if len(right) > 0 {
			r := right[0]
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				return false
			}
		}
	}
	return true
}

func main() {
	text := "Моя кошка любит молоко. Кошка — это доброе животное. ababa"
	//keywords := []string{"кошка", "молоко", "аб", "aba", "ababa"}

	fmt.Println("Исходный текст:")
	fmt.Println(text)
	fmt.Println()

	fmt.Println("1) Простая подсветка (регистр учитывается):")
	fmt.Println(highlightKeywords(text, []string{"кошка", "молоко"}, false, false))
	fmt.Println()

	fmt.Println("2) Нечувствительная к регистру:")
	fmt.Println(highlightKeywords(text, []string{"КОШКА", "молоко"}, true, false))
	fmt.Println()

	fmt.Println("3) Обработка пересечений и пересекающихся ключевых слов (пример латиницы):")
	latin := "ababa"
	fmt.Println("text:", latin)
	fmt.Println("keywords: [\"aba\"] ->", highlightKeywords(latin, []string{"aba"}, false, false))
	fmt.Println("keywords: [\"aba\",\"ababa\"] ->", highlightKeywords(latin, []string{"aba", "ababa"}, false, false))
	fmt.Println("keywords: [\"ab\",\"aba\",\"ababa\"] ->", highlightKeywords(latin, []string{"ab", "aba", "ababa"}, false, false))
	fmt.Println()

	fmt.Println("4) Только целые слова (wholeWord=true):")
	fmt.Println(highlightKeywords("hello world, helloworld hello_world", []string{"hello"}, false, true))
}
