package main

import (
	"fmt"
)

// setZeroes модифицирует матрицу на месте: если matrix[i][j] == 0,
// то вся i-я строка и j-й столбец становятся нулями
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])

	// Флаги: используем первую строку и первый столбец
	firstRowHasZero := false
	firstColHasZero := false

	// Проверяем, есть ли 0 в первой строке
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Проверяем, есть ли 0 в первом столбце
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// ШАГ 1: Используем первую строку и столбец как маркеры
	// Если matrix[i][j] == 0 → помечаем matrix[i][0] и matrix[0][j]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// ШАГ 2: Обнуляем строки по маркерам в первом столбце
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// ШАГ 3: Обнуляем столбцы по маркерам в первой строке
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// ШАГ 4: Обнуляем первую строку, если нужно
	if firstRowHasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// ШАГ 5: Обнуляем первый столбец, если нужно
	if firstColHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// Вспомогательная функция: печать матрицы
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// Пример 1
	matrix1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println("До:")
	printMatrix(matrix1)
	setZeroes(matrix1)
	fmt.Println("После:")
	printMatrix(matrix1)
	// Ожидаем:
	// 1 0 1
	// 0 0 0
	// 1 0 1

	// Пример 2
	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	fmt.Println("До:")
	printMatrix(matrix2)
	setZeroes(matrix2)
	fmt.Println("После:")
	printMatrix(matrix2)
	// Ожидаем:
	// 0 0 0 0
	// 0 4 5 0
	// 0 3 1 0
}
