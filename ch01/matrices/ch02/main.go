package main

import (
	"fmt"
)

// rotate поворачивает квадратную матрицу на 90° по часовой стрелке
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 || n != len(matrix[0]) {
		return
	}

	// ШАГ 1: Транспонируем матрицу (по главной диагонали)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Меняем matrix[i][j] и matrix[j][i]
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// ШАГ 2: Отражаем каждую строку по горизонтали (реверс строк)
	for i := 0; i < n; i++ {
		// Меняем элементы с двух концов строки
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
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
	// Пример 1: 3x3
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Исходная матрица:")
	printMatrix(matrix1)

	rotate(matrix1)

	fmt.Println("После поворота на 90° (по часовой):")
	printMatrix(matrix1)
	// Ожидаем:
	// 7 4 1
	// 8 5 2
	// 9 6 3

	fmt.Println("---")

	// Пример 2: 4x4
	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	fmt.Println("Исходная матрица:")
	printMatrix(matrix2)

	rotate(matrix2)

	fmt.Println("После поворота на 90° (по часовой):")
	printMatrix(matrix2)
	// Ожидаем:
	// 15 13  2  5
	// 14  3  4  1
	// 12  6  8  9
	// 16  7 10 11
}
