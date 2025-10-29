package main

//изучение работы данного алгоритма  (studying the operation of this algorithm)

import (
	"fmt"
)

// 1. Максимальная прибыль (одна сделка), prices{7, 1, 5, 3, 6, 4}, prices {7, 6, 4, 3, 1}
func maxProfitOneTransaction(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			if profit := price - minPrice; profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// 2. Анализ баланса (Кадане) — максимальная сумма подмассива
func maxSubarraySum(changes []int) int {
	maxSoFar := changes[0]
	currentMax := changes[0]
	for i := 1; i < len(changes); i++ {
		if currentMax < 0 {
			currentMax = changes[i]
		} else {
			currentMax += changes[i]
		}
		if currentMax > maxSoFar {
			maxSoFar = currentMax
		}
	}
	return maxSoFar
}

func main() {
	// 1
	fmt.Println("\n1)  Максимальная прибыль (одна сделка):")
	fmt.Println(maxProfitOneTransaction([]int{7, 1, 5, 3, 6, 4})) // ожидает 5
	fmt.Println(maxProfitOneTransaction([]int{7, 6, 4, 3, 1}))    // ожидает 0
	fmt.Println(maxProfitOneTransaction([]int{3, 1, 20, 3, 6, 1, 13}))

	// 2
	fmt.Println("\n2) Анализ баланса (Кадане):")
	fmt.Println(maxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // ожидает 6 (4,-1,2,1)
	fmt.Println(maxSubarraySum([]int{1, -2, -3, -4}))                 // ожидает 1

}
