package main

import (
	"fmt"
	"math"
	"sort"
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

// 4. Захват воды (Trapping Rain Water)
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	maxL, maxR := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxL {
				maxL = height[left]
			} else {
				water += maxL - height[left]
			}
			left++
		} else {
			if height[right] >= maxR {
				maxR = height[right]
			} else {
				water += maxR - height[right]
			}
			right--
		}
	}
	return water
}

// 5. Сортировка чёт/нечёт (Stable Partition by Parity)
func sortByParity(nums []int) []int {
	result := make([]int, len(nums))
	i := 0
	for _, n := range nums {
		if n%2 == 0 {
			result[i] = n
			i++
		}
	}
	for _, n := range nums {
		if n%2 != 0 {
			result[i] = n
			i++
		}
	}
	return result
}

// 6  Сумма трёх компонентов (Three Sum)
func threeSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L, R := i+1, n-1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == target {
				result = append(result, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum < target {
				L++
			} else {
				R--
			}
		}
	}
	return result
}

// 7. Обратный порядок регионов (Reverse Vowels)
func reverseVowels(s string) string {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	bytes := []byte(s)
	L, R := 0, len(bytes)-1

	for L < R {
		for L < R && !vowels[bytes[L]] {
			L++
		}
		for L < R && !vowels[bytes[R]] {
			R--
		}
		bytes[L], bytes[R] = bytes[R], bytes[L]
		L++
		R--
	}
	return string(bytes)
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

	fmt.Println("\n4)  Trapping Rain Water:")
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6

	fmt.Println("\n5)  Stable Partition by Parity:")
	fmt.Println(sortByParity([]int{2, 11, 4, 7, 30, 1})) // [2 4 30 11 7 1]

	fmt.Println("\n6)  Three Sum:")
	nums1 := []int{1, 2, -1, -2, 0}
	fmt.Println(threeSum(nums1, 0)) // [[-2 0 2] [-1 0 1]]

	fmt.Println("\n7) Reverse Vowels:")
	s := "Helsinki"
	fmt.Println(reverseVowels(s)) // Hilsenke

}
