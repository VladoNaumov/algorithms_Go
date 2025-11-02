package main

import (
	"fmt"
	"math"
)

// 1. Two Sum Closest
func twoSumClosest(nums []int, target int) int {
	if len(nums) < 2 {
		return 0
	}
	left, right := 0, len(nums)-1
	minDiff := math.MaxInt32 //  Max int32 value = 2147483647
	closestSum := 0

	for left < right {
		sum := nums[left] + nums[right]
		diff := sum - target
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			minDiff = diff
			closestSum = sum
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return closestSum
}

// 2. Sorted Squares
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			result[pos] = leftSq
			left++
		} else {
			result[pos] = rightSq
			right--
		}
		pos--
	}
	return result
}

// 3. Remove Duplicates In-Place
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	write := 1
	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[write-1] {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

func main() {

	fmt.Println("\n1) Two Sum Closest:")
	result := twoSumClosest([]int{-5, 1, 3, 6, 8}, 10)
	fmt.Println(result) // 9

	fmt.Println("\n2)  Sorted Squares:")
	result1 := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println(result1) // [0 1 9 16 100]

	fmt.Println("\n3)  Remove Duplicates In-Place:")
	nums := []int{2, 2, 3, 3, 5, 6, 6}
	newLen := removeDuplicates(nums)
	fmt.Println(nums[:newLen]) // [2 3 5 6]
	fmt.Println(newLen)        // 4

}
