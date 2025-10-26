package main

import (
	"fmt"
	"time"
)

func doOuter(a, b int) int {
	return a + b
}

func doMain(theIterac int) int {
	start := time.Now() // Запоминаем время начала

	theArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	theRes := 0

	for theIterac > 0 {
		theIterac--
		sum := 0
		for i, v := range theArr {
			if i == 0 {
				sum = v
			} else {
				sum = doOuter(sum, v)
			}
		}
		theRes = sum
	}

	elapsed := time.Since(start) // Вычисляем время выполнения
	fmt.Printf("Код выполнялся %s\n", elapsed)

	return theRes
}

func main() {
	fmt.Println("Результат:", doMain(1_000_000_000))
}
