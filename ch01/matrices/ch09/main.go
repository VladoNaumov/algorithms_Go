package main

import "fmt"

/*
-------------------------------------------------
 9. Транспонирование (m×n → n×m)
    -------------------------------------------------
*/
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}

func main() {
	fmt.Println("9:", transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
