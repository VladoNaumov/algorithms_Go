package main

import "fmt"

// 11.
// rotateMatrix поворачивает квадратную матрицу на angle градусов по часовой
func rotateMatrix(matrix [][]int, angle int) {
	n := len(matrix)
	angle = (angle%360 + 360) % 360 // нормализуем
	if angle == 0 {
		return
	}

	// 180° = дважды 90°
	if angle == 180 {
		rotateMatrix(matrix, 90)
		rotateMatrix(matrix, 90)
		return
	}

	// 270° = трижды 90°
	if angle == 270 {
		rotateMatrix(matrix, 90)
		rotateMatrix(matrix, 90)
		rotateMatrix(matrix, 90)
		return
	}

	// 90°: транспонирование + реверс строк
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

// 12.
func rotateLayers(matrix [][]int, k int) {
	m, n := len(matrix), len(matrix[0])
	k %= max(m, n) // на всякий случай

	layers := min(m, n) / 2
	for layer := 0; layer < layers; layer++ {
		top, left := layer, layer
		bottom, right := m-1-layer, n-1-layer

		// количество элементов в кольце
		count := 2*(bottom-top+1) + 2*(right-left) - 4
		kRing := k % count

		for i := 0; i < kRing; i++ {
			// сохраняем верхний левый
			tmp := matrix[top][left]

			// left → top
			for r := top; r < bottom; r++ {
				matrix[r][left] = matrix[r+1][left]
			}
			// bottom → left
			for c := left; c < right; c++ {
				matrix[bottom][c] = matrix[bottom][c+1]
			}
			// right → bottom
			for r := bottom; r > top; r-- {
				matrix[r][right] = matrix[r-1][right]
			}
			// top → right
			for c := right; c > left+1; c-- {
				matrix[top][c] = matrix[top][c-1]
			}
			matrix[top][left+1] = tmp
		}
	}
}

// 13.
func isSymmetric(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}

// 14.
func rle2D(matrix [][]int) []struct{ Val, Count int } {
	if len(matrix) == 0 {
		return nil
	}
	var result []struct{ Val, Count int }
	prev := matrix[0][0]
	count := 1

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if matrix[i][j] == prev {
				count++
			} else {
				result = append(result, struct{ Val, Count int }{prev, count})
				prev, count = matrix[i][j], 1
			}
		}
	}
	result = append(result, struct{ Val, Count int }{prev, count})
	return result
}

// 15.
func swapDiagonals(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		matrix[i][i], matrix[i][n-1-i] = matrix[i][n-1-i], matrix[i][i]
	}
}

// 16.
func rowColSums(matrix [][]int) ([]int, []int) {
	m, n := len(matrix), len(matrix[0])
	rowSums := make([]int, m)
	colSums := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowSums[i] += matrix[i][j]
			colSums[j] += matrix[i][j]
		}
	}
	return rowSums, colSums
}

// 17.
func transposeInPlace(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 18.
func findMax(matrix [][]int) (maxVal, row, col int) {
	maxVal = matrix[0][0]
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > maxVal {
				maxVal, row, col = matrix[i][j], i, j
			}
		}
	}
	return
}

// 19.
func countEvenOdd(matrix [][]int) (even, odd int) {
	for i := range matrix {
		for _, v := range matrix[i] {
			if v%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	return
}

// 20.
func shiftRowsRight(matrix [][]int, k int) {
	m, n := len(matrix), len(matrix[0])
	k %= n
	for i := 0; i < m; i++ {
		row := append(matrix[i][n-k:], matrix[i][:n-k]...)
		copy(matrix[i], row)
	}
}

// 21.
func shiftColsDown(matrix [][]int, k int) {
	m, n := len(matrix), len(matrix[0])
	k %= m
	for j := 0; j < n; j++ {
		col := make([]int, m)
		for i := 0; i < m; i++ {
			col[i] = matrix[i][j]
		}
		shifted := append(col[m-k:], col[:m-k]...)
		for i := 0; i < m; i++ {
			matrix[i][j] = shifted[i]
		}
	}
}

// 22.
func zeroNegatives(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < 0 {
				matrix[i][j] = 0
			}
		}
	}
}

// 23
func multiplyBy(matrix [][]int, k int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] *= k
		}
	}
}

// 24.
func diagonalSums(matrix [][]int) (main, anti int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		main += matrix[i][i]
		anti += matrix[i][n-1-i]
	}
	if n%2 == 1 {
		anti -= matrix[n/2][n/2] // центр учтён дважды
	}
	return
}

// 25.
func reverseTraverse(matrix [][]int) []int {
	result := []int{}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			result = append(result, matrix[i][j])
		}
	}
	return result
}

func main() {
	// 11. Поворот 90°
	m11 := [][]int{{1, 2}, {3, 4}}
	rotateMatrix(m11, 90)
	fmt.Println("11 (90°):", m11) // [[3 1] [4 2]]

	// 12. Ротация колец
	m12 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotateLayers(m12, 1)
	fmt.Println("12 (rotate layers):")
	printMatrix(m12)

	// 13. Проверка симметрии
	m13 := [][]int{{1, 2, 1}, {2, 3, 2}, {1, 2, 1}}
	fmt.Println("13 symmetric:", isSymmetric(m13)) // true

	// 14. RLE 2D
	m14 := [][]int{{1, 1, 2}, {2, 2, 2}, {3, 3, 3}}
	rle := rle2D(m14)
	fmt.Println("14 RLE:")
	for _, v := range rle {
		fmt.Printf("(%d, %d) ", v.Val, v.Count)
	}
	fmt.Println()

	// 15. Замена диагоналей
	m15 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	swapDiagonals(m15)
	fmt.Println("15 swap diagonals:", m15)

	// 16. Сумма строк и столбцов
	rowSums, colSums := rowColSums([][]int{{1, 2}, {3, 4}})
	fmt.Println("16 rows:", rowSums, "cols:", colSums)

	// 17. Транспонирование in-place
	m17 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	transposeInPlace(m17)
	fmt.Println("17 transpose in-place:", m17)

	// 18. Максимальный элемент
	maxVal, r, c := findMax([][]int{{-1, 2, 3}, {4, 5, -9}})
	fmt.Printf("18 max: %d at [%d,%d]\n", maxVal, r, c)

	// 19. Чётные/нечётные
	even, odd := countEvenOdd([][]int{{1, 2}, {3, 4}, {5, 6}})
	fmt.Println("19 even:", even, "odd:", odd)

	// 20. Сдвиг строк вправо
	m20 := [][]int{{1, 2, 3}, {4, 5, 6}}
	shiftRowsRight(m20, 1)
	fmt.Println("20 shift rows right:", m20)

	// 21. Сдвиг столбцов вниз
	m21 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	shiftColsDown(m21, 1)
	fmt.Println("21 shift cols down:")
	printMatrix(m21)

	// 22. Замена отрицательных на 0
	m22 := [][]int{{-1, 2}, {-3, -4}}
	zeroNegatives(m22)
	fmt.Println("22 zero negatives:", m22)

	// 23. Умножение на константу
	m23 := [][]int{{1, 2}, {3, 4}}
	multiplyBy(m23, 3)
	fmt.Println("23 multiply by 3:", m23)

	// 24. Сумма диагоналей
	mainSum, antiSum := diagonalSums([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println("24 main:", mainSum, "anti:", antiSum)

	// 25. Обратный обход
	fmt.Println("25 reverse:", reverseTraverse([][]int{{1, 2}, {3, 4}}))
}

// Вспомогательная функция печати
func printMatrix(m [][]int) {
	for _, row := range m {
		fmt.Println(row)
	}
}
