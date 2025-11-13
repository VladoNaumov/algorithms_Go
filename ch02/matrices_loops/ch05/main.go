package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
	}
	dirs := []int{-1, 0, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, cnt := 0, 0
			for di := range dirs {
				for dj := range dirs {
					ni, nj := i+di, j+dj
					if ni >= 0 && ni < m && nj >= 0 && nj < n {
						sum += img[ni][nj]
						cnt++
					}
				}
			}
			out[i][j] = sum / cnt
		}
	}
	return out
}

func main() {
	img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println("5:", imageSmoother(img))
}
