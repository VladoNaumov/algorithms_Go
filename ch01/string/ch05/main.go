package main

import (
	"errors"
	"fmt"
)

const (
	INT32_MAX = 1<<31 - 1
	INT32_MIN = -1 << 31
)

// Atoi32 парсит строку в int32 с поведением:
// - пропускает ведущие пробелы
// - учитывает один знак +/-, если есть
// - читает подряд идущие цифры и останавливается при первом не-цифровом символе
// - при переполнении возвращает INT32_MAX / INT32_MIN
// - если не найдено ни одной цифры — возвращает 0 и ошибку
func Atoi32(s string) (int32, error) {
	i := 0
	n := len(s)

	// 1) пропускаем ведущие пробелы
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0, errors.New("no digits")
	}

	// 2) знак
	sign := int64(1)
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// 3) читаем цифры
	var acc int64 = 0
	digitsFound := false
	for i < n {
		ch := s[i]
		if ch < '0' || ch > '9' {
			break
		}
		digitsFound = true
		d := int64(ch - '0')

		// проверка переполнения: acc * 10 + d > INT32_MAX (в учёте знака)
		// делаем проверку в положительной форме, учитывая sign
		if sign == 1 {
			if acc > (int64(INT32_MAX)-d)/10 {
				return int32(INT32_MAX), nil
			}
		} else { // sign == -1
			// хотим проверить, что -(acc*10 + d) >= INT32_MIN
			// или acc*10 + d <= -INT32_MIN  (поскольку INT32_MIN отрицательное)
			// -INT32_MIN == 1<<31  (т.е. 2147483648)
			if acc > (int64(-INT32_MIN)-d)/10 {
				return int32(INT32_MIN), nil
			}
		}

		acc = acc*10 + d
		i++
	}

	if !digitsFound {
		return 0, errors.New("no digits")
	}

	res := acc * sign
	return int32(res), nil
}

func main() {
	cases := []string{
		"   -42",
		"4193 with words",
		"9223372036854775808", // слишком большое
		"words and 987",
		"    +000123",
		" -2147483649", // чуть меньше INT32_MIN -> переполнение вниз
		"2147483648",   // чуть больше INT32_MAX -> переполнение вверх
		"",             // пустая строка
		"   +",         // только знак, нет цифр
	}

	for _, c := range cases {
		val, err := Atoi32(c)
		if err != nil {
			fmt.Printf("%q -> error: %v (returned %d)\n", c, err, val)
		} else {
			fmt.Printf("%q -> %d\n", c, val)
		}
	}
}
