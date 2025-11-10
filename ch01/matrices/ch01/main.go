package main

import (
	"fmt"
)

// spiralOrder обходит матрицу по спирали и возвращает элементы в порядке обхода
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// Определяем границы обхода
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	result := []int{}

	// Продолжаем, пока границы не пересекутся
	for top <= bottom && left <= right {
		// 1. Идём слева направо по верхней строке
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++ // Сдвигаем верхнюю границу вниз

		// 2. Идём сверху вниз по правому столбцу
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right-- // Сдвигаем правую границу влево

		// Проверяем, не пересеклись ли границы (чтобы не дублировать)
		if top <= bottom {
			// 3. Идём справа налево по нижней строке
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom-- // Сдвигаем нижнюю границу вверх
		}

		if left <= right {
			// 4. Идём снизу вверх по левому столбцу
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++ // Сдвигаем левую границу вправо
		}
	}

	return result
}

// Вспомогательная функция для красивого вывода матрицы
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	// Пример 1: 3x3 матрица
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Матрица:")
	printMatrix(matrix1)
	fmt.Println("Спиральный обход:", spiralOrder(matrix1))
	// Вывод: [1 2 3 6 9 8 7 4 5]

	fmt.Println()

	// Пример 2: 4x4 матрица
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("Матрица:")
	printMatrix(matrix2)
	fmt.Println("Спиральный обход:", spiralOrder(matrix2))
	// Вывод: [1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10]
}
