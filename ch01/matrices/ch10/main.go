package main

import (
	"fmt"
)

/*
-------------------------------------------------
 10. Сумма подматрицы (префиксные суммы 2D)
    -------------------------------------------------
*/
type NumMatrix struct{ prefix [][]int }

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	p := make([][]int, m+1)
	for i := range p {
		p[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			p[i][j] = matrix[i-1][j-1] + p[i-1][j] + p[i][j-1] - p[i-1][j-1]
		}
	}
	return NumMatrix{p}
}
func (nm *NumMatrix) SumRegion(r1, c1, r2, c2 int) int {
	a, b, c, d := r1+1, c1+1, r2+1, c2+1
	return nm.prefix[d][c] - nm.prefix[d][b-1] - nm.prefix[a-1][c] + nm.prefix[a-1][b-1]
}

/*
-------------------------------------------------

	Тесты
	-------------------------------------------------
*/
func main() {

	nm := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	fmt.Println("10:", nm.SumRegion(2, 1, 4, 3))

}
