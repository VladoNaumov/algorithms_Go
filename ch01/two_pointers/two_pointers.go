package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"unicode"
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
	var result [][]int
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

//8. Удаление старых версий (Remove Element In-Place)

func removeElement(nums []int, val int) int {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != val {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

// 9. Сжатие данных (String Compression)
func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	write := 0
	read := 0

	for read < len(chars) {
		char := chars[read]
		count := 0
		for read < len(chars) && chars[read] == char {
			read++
			count++
		}
		chars[write] = char
		write++
		if count > 1 {
			for _, d := range strconv.Itoa(count) {
				chars[write] = byte(d)
				write++
			}
		}
	}
	return write
}

// 10. Поиск медианы двух отсортированных баз (Median of Two Sorted Arrays)
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := totalLeft - i

		Aleft := math.Inf(-1)
		if i > 0 {
			Aleft = float64(nums1[i-1])
		}
		Aright := math.Inf(1)
		if i < m {
			Aright = float64(nums1[i])
		}
		Bleft := math.Inf(-1)
		if j > 0 {
			Bleft = float64(nums2[j-1])
		}
		Bright := math.Inf(1)
		if j < n {
			Bright = float64(nums2[j])
		}

		if Aleft <= Bright && Bleft <= Aright {
			if (m+n)%2 == 1 {
				return math.Max(Aleft, Bleft)
			}
			return (math.Max(Aleft, Bleft) + math.Min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return 0.0
}

// 11. Парковка (Two-Side Matching)
func carParking(spots, cars []int) int {
	sort.Ints(spots)
	sort.Ints(cars)
	count := 0
	j := 0
	for i := 0; i < len(spots) && j < len(cars); i++ {
		if cars[j] <= spots[i] {
			count++
			j++
		}
	}
	return count
}

// 12. Оптимизация долга/актива
func matchDebtsCreditors(debtors, creditors []int) [][2]int {
	sort.Ints(debtors)
	sort.Ints(creditors)
	var pairs [][2]int
	i, j := 0, 0
	for i < len(debtors) && j < len(creditors) {
		if debtors[i] <= creditors[j] {
			pairs = append(pairs, [2]int{debtors[i], creditors[j]})
			i++
			j++
		} else {
			i++
		}
	}
	return pairs
}

// 13. Корректировка палиндрома
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}
		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// 14. Нахождение пары с близкой суммой
func closestSumPair(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	minDiff := math.MaxInt32
	var result []int

	for left < right {
		sum := nums[left] + nums[right]
		diff := int(math.Abs(float64(sum - target)))
		if diff < minDiff {
			minDiff = diff
			result = []int{nums[left], nums[right]}
		}
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return result
		}
	}
	return result
}

// 15. Разделение положительных/отрицательных
func separateNegatives(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] < 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
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

	fmt.Println("\n8) Remove Element In-Place:")
	nums2 := []int{3, 2, 2, 3}
	val2 := 3
	newLen2 := removeElement(nums2, val2)
	fmt.Println(nums2[:newLen2]) // [2 2]
	fmt.Println(newLen)          // 2

	fmt.Println("\n9) String Compression:")
	chars := []byte{'A', 'A', 'A', 'A', 'B'}
	newLen3 := compress(chars)
	fmt.Println(string(chars[:newLen3])) // A4B
	fmt.Println(newLen)                  // 3

	fmt.Println("\n10) Median of Two Sorted Arrays:")
	median := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(median) // 2.0

	fmt.Println("\n11) Two-Side Matching:")
	spots := []int{2, 3, 4}
	cars := []int{1, 3, 4}
	result2 := carParking(spots, cars)
	fmt.Println(result2) // 3

	fmt.Println("\n12) Оптимизация долга/актива:")
	debtors := []int{1, 3, 5}
	creditors := []int{2, 6}
	pairs := matchDebtsCreditors(debtors, creditors)
	fmt.Println(pairs) // [[1 2] [3 6]]

	fmt.Println("\n13) Корректировка палиндрома:")
	s = "A man, a plan, a canal, Panama"
	fmt.Println(isPalindrome(s)) // true

	fmt.Println("\n14) Корректировка палиндрома:")
	nums3 := []int{1, 2, 3, 4, 5}
	target := 8
	pair := closestSumPair(nums3, target)
	fmt.Println(pair) // [3 5]

	fmt.Println("\n15) Корректировка палиндрома:")
	nums4 := []int{1, -2, 3, -4, 5}
	separateNegatives(nums4)
	fmt.Println(nums4) // [-2 -4 1 3 5] или [-2 -4 3 1 5]
}
