package main

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatCurrency форматирует число amount в валютную строку.
func FormatCurrency(amount float64, symbol string, symbolBefore bool, decimals int, thousandSep, decimalSep rune, spaceBetween bool, parenthesesForNegative bool) string {
	// 1) Обрабатываем знак
	neg := amount < 0
	absVal := amount
	if neg {
		absVal = -absVal
	}

	// 2) Форматируем с округлением в строку с десятичной точкой (локаль-нейтрально)
	// strconv.FormatFloat делает округление до заданной точности
	numStr := strconv.FormatFloat(absVal, 'f', decimals, 64) // пример: "1234567.89" или "1234.50" или "1000"

	// 3) Разделяем целую и дробную часть (decimal point — '.')
	intPart := numStr
	fracPart := ""
	if decimals > 0 {
		parts := strings.SplitN(numStr, ".", 2)
		intPart = parts[0]
		if len(parts) == 2 {
			fracPart = parts[1]
		} else {
			// если по какой-то причине дробной части нет (маловероятно), заполняем нулями
			fracPart = strings.Repeat("0", decimals)
		}
	}

	// 4) Вставляем разделитель тысяч в intPart (с конца)
	var b strings.Builder
	runes := []rune(intPart)
	n := len(runes)
	for i, r := range runes {
		// позиция с начала: i, с конца индекс = n-1-i
		// вставляем символ, потом, если осталось три цифры группы справа — ставим thousandSep
		b.WriteRune(r)
		// if there are characters remaining and position from right is multiple of 3 -> insert sep
		remaining := n - 1 - i
		if remaining > 0 && remaining%3 == 0 {
			// вставляем разделитель тысяч
			b.WriteRune(thousandSep)
		}
	}
	intWithSep := b.String()

	// 5) Собираем число с дробной частью (и нужным десятичным сепаратором)
	var numberBuilder strings.Builder
	numberBuilder.WriteString(intWithSep)
	if decimals > 0 {
		numberBuilder.WriteRune(decimalSep)
		// гарантируем длину дробной части = decimals (FormatFloat уже дал нужную длину)
		if len(fracPart) < decimals {
			fracPart = fracPart + strings.Repeat("0", decimals-len(fracPart))
		}
		numberBuilder.WriteString(fracPart)
	}
	numberStr := numberBuilder.String()

	// 6) Собираем окончательный результат с символом валюты и знаком
	var result string
	space := ""
	if spaceBetween {
		space = " "
	}

	if symbolBefore {
		// Символ перед числом: "$1,234.56" или "($1,234.56)" при parenthesesForNegative
		if neg && parenthesesForNegative {
			result = fmt.Sprintf("(%s%s%s)", symbol, space, numberStr)
		} else {
			prefix := ""
			if neg && !parenthesesForNegative {
				prefix = "-" // минус перед всей конструкцией
			}
			result = fmt.Sprintf("%s%s%s", prefix, symbol, space+numberStr)
		}
	} else {
		// Символ после числа: "1 234,56 €"
		if neg && parenthesesForNegative {
			result = fmt.Sprintf("(%s%s%s)", numberStr, space, symbol)
		} else {
			prefix := ""
			if neg && !parenthesesForNegative {
				prefix = "-" // минус перед числом
			}
			result = fmt.Sprintf("%s%s%s%s", prefix, numberStr, space, symbol)
			result = strings.TrimSpace(result) // убрать лишний пробел, если space == ""
		}
	}

	return result
}

func main() {
	examples := []struct {
		amount              float64
		symbol              string
		symbolBefore        bool
		decimals            int
		thousandSep, decSep rune
		spaceBetween        bool
		parenthesesForNeg   bool
		description         string
	}{
		{1234567.891, "$", true, 2, ',', '.', false, false, "US format, $ before"},
		{-1234.5, "€", false, 2, ' ', ',', true, false, "EU format, symbol after with space"},
		{-1234.5, "€", false, 2, ' ', ',', true, true, "EU format, negative in parentheses"},
		{1000.0, "¥", true, 0, ',', '.', false, false, "No decimals, yen"},
		{1234.5678, "USD", true, 3, ',', '.', true, false, "3 decimals, symbol before with space"},
		{0, "€", false, 2, ' ', ',', true, false, "Zero value"},
	}

	for _, ex := range examples {
		out := FormatCurrency(ex.amount, ex.symbol, ex.symbolBefore, ex.decimals, ex.thousandSep, ex.decSep, ex.spaceBetween, ex.parenthesesForNeg)
		fmt.Printf("%-45s -> %s\n", ex.description, out)
	}

	// Пара конкретных примеров, как в задаче:
	fmt.Println()
	fmt.Println("Concrete examples:")
	fmt.Println(FormatCurrency(1234567.891, "$", true, 2, ',', '.', false, false)) // "$1,234,567.89"
	fmt.Println(FormatCurrency(-1234.5, "€", false, 2, ' ', ',', true, false))     // "-1 234,50 €"
	fmt.Println(FormatCurrency(-1234.5, "€", false, 2, ' ', ',', true, true))      // "(1 234,50 €)"
}
