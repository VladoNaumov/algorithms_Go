package main

import "fmt"

/*
-------------------------------------------------
 7. Самая большая область 1 (гистограмма + стек)
    -------------------------------------------------
*/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	maxArea := 0

	for i := 0; i < m; i++ {
		// обновляем высоты текущей строки
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

// гистограмма → наибольший прямоугольник (монотонный стек)
func largestRectangleArea(h []int) int {
	stack := []int{}
	area := 0
	h = append(h, 0) // маркер конца
	for i := 0; i < len(h); i++ {
		for len(stack) > 0 && h[stack[len(stack)-1]] > h[i] {
			height := h[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area = max(area, height*width)
		}
		stack = append(stack, i)
	}
	return area
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 7
	mat := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println("7:", maximalRectangle(mat))
}
