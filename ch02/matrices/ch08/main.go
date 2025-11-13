package main

import "fmt"

func flipVertical(matrix [][]int) {
	for _, row := range matrix {
		left, right := 0, len(row)-1
		for left < right {
			row[left], row[right] = row[right], row[left]
			left++
			right--
		}
	}
}

func main() {
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	flipVertical(m1)
	fmt.Println("8:", m1)
}
