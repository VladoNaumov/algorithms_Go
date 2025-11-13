package main

import (
	"fmt"
	"math"
	"sort"
)

// Task 1: Длиннейшая подстрока без повторяющихся символов
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	maxLen := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if idx, ok := charMap[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		charMap[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// Task 2: Минимальное окно-подстрока (Поиск тегов)
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	window := make(map[byte]int)
	left, minLen, start := 0, len(s)+1, 0
	count := 0
	for right := 0; right < len(s); right++ {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				count++
			}
		}
		for count == len(need) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			if _, ok := need[s[left]]; ok {
				window[s[left]]--
				if window[s[left]] < need[s[left]] {
					count--
				}
			}
			left++
		}
	}
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

// Task 3: Подмассив с заданной суммой (Целевой бюджет)
func subarraySum(nums []int, k int) int {
	prefix := make(map[int]int)
	prefix[0] = 1
	sum, count := 0, 0
	for _, num := range nums {
		sum += num
		if val, ok := prefix[sum-k]; ok {
			count += val
		}
		prefix[sum]++
	}
	return count
}

// Task 4: Начало анаграмм (Анализ текстов)
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	pCount := [26]int{}
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
	}
	window := [26]int{}
	result := []int{}
	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++
		if i >= len(p) {
			window[s[i-len(p)]-'a']--
		}
		if window == pCount {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

// Task 5: Максимальный средний балл (Прогноз)
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}

// Task 6: Максимум в скользящем окне
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	deque := []int{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

// Task 7: Максимальное число последовательных периодов
func longestOnes(nums []int, k int) int {
	left, zeros, maxLen := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// Task 8: Самая короткая подстрока с K уникальными символами
func shortestSubstringWithKUnique(s string, k int) int {
	if k > len(s) {
		return -1
	}
	charCount := make(map[byte]int)
	left, minLen := 0, math.MaxInt32
	unique := 0
	for right := 0; right < len(s); right++ {
		charCount[s[right]]++
		if charCount[s[right]] == 1 {
			unique++
		}
		for unique == k {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				unique--
			}
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return -1
	}
	return minLen
}

// Task 9: Подмассив с произведением меньше K
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod, left, count := 1, 0, 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count
}

// Task 10: Минимальная длина подмассива (сумма ≥ S)
func minSubArrayLen(target int, nums []int) int {
	left, sum, minLen := 0, 0, math.MaxInt32
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

// Task 11: Средняя температура за K дней
func dailyTemperatures(temps []int, k int) []float64 {
	if len(temps) < k {
		return []float64{}
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += temps[i]
	}
	result := []float64{float64(sum) / float64(k)}
	for i := k; i < len(temps); i++ {
		sum += temps[i] - temps[i-k]
		result = append(result, float64(sum)/float64(k))
	}
	return result
}

// Task 12: Максимальное число клиентов в смене
func maxConcurrentClients(intervals [][]int) int {
	events := []struct {
		time, delta int
	}{}
	for _, interval := range intervals {
		events = append(events, struct{ time, delta int }{interval[0], 1})
		events = append(events, struct{ time, delta int }{interval[1], -1})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].delta < events[j].delta
		}
		return events[i].time < events[j].time
	})
	maxCount, count := 0, 0
	for _, e := range events {
		count += e.delta
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

// Task 13: Максимальная прибыль с двумя транзакциями
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy1, sell1 := math.MinInt32, 0
	buy2, sell2 := math.MinInt32, 0
	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1+p)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, buy2+p)
	}
	return sell2
}

// Task 14: Минимальное количество товаров для набора
func minWindowWithIngredients(s string, ingredients string) string {
	if len(ingredients) > len(s) {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(ingredients); i++ {
		need[ingredients[i]]++
	}
	window := make(map[byte]int)
	left, minLen, start := 0, len(s)+1, 0
	count := 0
	for right := 0; right < len(s); right++ {
		if _, ok := need[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				count++
			}
		}
		for count == len(need) {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			if _, ok := need[s[left]]; ok {
				window[s[left]]--
				if window[s[left]] < need[s[left]] {
					count--
				}
			}
			left++
		}
	}
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

// Task 15: Длиннейшая последовательность покупок
func longestSubarrayWithAtMostK(items []int, k int) int {
	count := make(map[int]int)
	left, maxLen := 0, 0
	for right := 0; right < len(items); right++ {
		count[items[right]]++
		for count[items[right]] > k {
			count[items[left]]--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Task 1
	fmt.Println("Task 1:", lengthOfLongestSubstring("abcabcbb")) // 3

	// Task 2
	fmt.Println("Task 2:", minWindow("ADOBECODEBANC", "ABC")) // "BANC"

	// Task 3
	fmt.Println("Task 3:", subarraySum([]int{1, 1, 1}, 2)) // 2

	// Task 4
	fmt.Println("Task 4:", findAnagrams("cbaebabacd", "abc")) // [0 6]

	// Task 5
	fmt.Println("Task 5:", findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // 12.75

	// Task 6
	fmt.Println("Task 6:", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3 3 5 5 6 7]

	// Task 7
	fmt.Println("Task 7:", longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)) // 6

	// Task 8
	fmt.Println("Task 8:", shortestSubstringWithKUnique("abcabc", 3)) // 3

	// Task 9
	fmt.Println("Task 9:", numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)) // 8

	// Task 10
	fmt.Println("Task 10:", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})) // 2

	// Task 11
	fmt.Println("Task 11:", dailyTemperatures([]int{1, 2, 3, 4, 5}, 3)) // [2 3 4]

	// Task 12
	fmt.Println("Task 12:", maxConcurrentClients([][]int{{1, 3}, {2, 4}, {3, 5}})) // 3

	// Task 13
	fmt.Println("Task 13:", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})) // 6

	// Task 14
	fmt.Println("Task 14:", minWindowWithIngredients("ADOBECODEBANC", "ABC")) // "BANC"

	// Task 15
	fmt.Println("Task 15:", longestSubarrayWithAtMostK([]int{1, 2, 1, 2, 3, 1, 2}, 2)) // 7
}
