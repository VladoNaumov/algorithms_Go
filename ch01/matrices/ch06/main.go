package main

import "fmt"

/*
-------------------------------------------------
 6. Проверка судоку
    -------------------------------------------------
*/
func isValidSudoku(board [][]byte) bool {
	row := [9]map[byte]bool{}
	col := [9]map[byte]bool{}
	box := [9]map[byte]bool{}
	for i := range row {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			b := board[i][j]
			bx := (i/3)*3 + j/3
			if row[i][b] || col[j][b] || box[bx][b] {
				return false
			}
			row[i][b] = true
			col[j][b] = true
			box[bx][b] = true
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("6:", isValidSudoku(board))
}
