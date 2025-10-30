package main

//изучение работы данного алгоритма  (studying the operation of this algorithm)

import (
	"fmt"
)

// 5. Пропущенный платёж (числа 1..N, одно пропущено)
func findMissingPayment(payments []int) int {
	n := len(payments) + 1 // поскольку одно пропущено
	sumExpected := n * (n + 1) / 2
	sumActual := 0
	for _, v := range payments {
		sumActual += v
	}
	return sumExpected - sumActual
}

func main() {
	// 5
	fmt.Println("\n5) findMissingPayment:")
	fmt.Println(findMissingPayment([]int{3, 1, 4, 2, 6, 5}))
	fmt.Println(findMissingPayment([]int{1, 2, 3, 5, 6, 7}))
}
