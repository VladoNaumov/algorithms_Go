package main

//изучение работы данного алгоритма  (studying the operation of this algorithm)

import (
	"fmt"
)

// 6. Единственный сотрудник (все кроме одного встречаются дважды) — XOR
func singleEmployeeID(ids []int) int {
	x := 0
	for _, id := range ids {
		x ^= id
	}
	return x
}

func main() {
	// 6
	fmt.Println("\n6) singleEmployeeID:")
	fmt.Println(singleEmployeeID([]int{7, 3, 5, 3, 5, 7, 4})) // ожидает 4
	fmt.Println(singleEmployeeID([]int{2, 2, 1}))             // ожидает 1
}
