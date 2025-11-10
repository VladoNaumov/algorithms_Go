package main

import "fmt"

/*
-------------------------------------------------
 4. Проход по диагоналям (i+j = const)
    -------------------------------------------------
*/
func diagonalTraverse(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := []int{}

	// диагональ = i+j
	for k := 0; k < m+n-1; k++ {
		// стартовая точка диагонали
		i, j := 0, k
		if k >= n {
			i, j = k-n+1, n-1
		}
		// собираем элементы в порядке возрастания i (спуск вниз-влево)
		tmp := []int{}
		for i < m && j >= 0 {
			tmp = append(tmp, matrix[i][j])
			i++
			j--
		}
		// чётные диагонали — переворачиваем (идём снизу-вверх)
		if k%2 == 0 {
			for x := len(tmp) - 1; x >= 0; x-- {
				res = append(res, tmp[x])
			}
		} else {
			res = append(res, tmp...)
		}
	}
	return res
}

func main() {
	fmt.Println("4:", diagonalTraverse([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
