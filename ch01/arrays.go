package main

import (
	"fmt"
	"math"
	"sort"
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

// 3. Сортировка по категориям (Dutch National Flag) — ин-месте
func sortPriorities(orders []int) {
	low, mid, high := 0, 0, len(orders)-1
	for mid <= high {
		switch orders[mid] {
		case 0:
			orders[low], orders[mid] = orders[mid], orders[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			orders[mid], orders[high] = orders[high], orders[mid]
			high--
		}
	}
}

// 4. Сдвиг данных (ротация вправо на K)
func rotateRight(logs []int, k int) {
	n := len(logs)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}
	reverse := func(a []int, l, r int) {
		for l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	reverse(logs, 0, n-1)
	reverse(logs, 0, k-1)
	reverse(logs, k, n-1)
}

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

// 6. Единственный сотрудник (все кроме одного встречаются дважды) — XOR
func singleEmployeeID(ids []int) int {
	x := 0
	for _, id := range ids {
		x ^= id
	}
	return x
}

// 7. Продукт конкурентов (product of array except self) — без деления
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	left := make([]int, n)
	right := make([]int, n)
	out := make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		out[i] = left[i] * right[i]
	}
	return out
}

// 8. Максимальная последовательная убыль цен
func maxConsecutiveDeclines(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	current := 0
	best := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			current++
			if current > best {
				best = current
			}
		} else {
			current = 0
		}
	}
	return best
}

// 9. Сдвиг неактивных пользователей (переместить 0 в конец, сохранять порядок)
func moveZerosToEnd(users []int) {
	write := 0
	for _, v := range users {
		if v != 0 {
			users[write] = v
			write++
		}
	}
	for i := write; i < len(users); i++ {
		users[i] = 0
	}
}

// 10. Топ-3 самых частых рейтингов (1..5)
func topThreeRatings(ratings []int) []int {
	count := make([]int, 6) // индекс 1..5
	for _, r := range ratings {
		if r >= 1 && r <= 5 {
			count[r]++
		}
	}
	type pair struct {
		rating    int
		frequency int
	}
	pairs := make([]pair, 0, 5)
	for r := 1; r <= 5; r++ {
		pairs = append(pairs, pair{rating: r, frequency: count[r]})
	}
	// сортируем по убыванию частоты, при равенстве — по убыванию рейтинга
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].frequency == pairs[j].frequency {
			return pairs[i].rating > pairs[j].rating
		}
		return pairs[i].frequency > pairs[j].frequency
	})
	result := make([]int, 0, 3)
	for i := 0; i < 3 && i < len(pairs); i++ {
		if pairs[i].frequency > 0 {
			result = append(result, pairs[i].rating)
		}
	}
	return result
}

// 11. Баланс транзакций — день с минимальным накопленным балансом
// возвращаем индекс дня (0-based) и минимальный баланс
func dayWithMinCumulativeBalance(transactions []int) (int, int) {
	minSum := math.MaxInt64
	cur := 0
	minIndex := 0
	for i, t := range transactions {
		cur += t
		if cur < minSum {
			minSum = cur
			minIndex = i
		}
	}
	return minIndex, minSum
}

// 12. Проверка ISBN-10 (строка из 10 цифр)
func isValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 10; i++ {
		d := int(isbn[i] - '0')
		if d < 0 || d > 9 {
			return false
		}
		sum += (i + 1) * d
	}
	return sum%11 == 0
}

// 13. Поиск пиковой нагрузки (найти index пика)
func findPeak(energy []int) (int, int) {
	// предположение: массив длины >=3 и пик существует
	for i := 1; i < len(energy)-1; i++ {
		if energy[i] > energy[i-1] && energy[i] > energy[i+1] {
			return i, energy[i]
		}
	}
	return -1, 0 // на случай, если пик не найден
}

// 14. Самая длинная серия побед/поражений (1/0)
func longestSameRun(results []int) int {
	if len(results) == 0 {
		return 0
	}
	best := 1
	cur := 1
	for i := 1; i < len(results); i++ {
		if results[i] == results[i-1] {
			cur++
		} else {
			if cur > best {
				best = cur
			}
			cur = 1
		}
	}
	if cur > best {
		best = cur
	}
	return best
}

// 15. Проверка уникальности ID (есть ли дубликаты)
func hasDuplicates(ids []int) bool {
	seen := make(map[int]bool)
	for _, id := range ids {
		if seen[id] {
			return true
		}
		seen[id] = true
	}
	return false
}

// 16. Слияние двух отсортированных массивов в первом (in-place).
// first — срез длины m+n, где первые m элементов — реальные значения,
// tail (m..m+n-1) зарезервирован под копирование. second — срез длины n.
func mergeIntoFirst(first []int, m int, second []int, n int) {
	// pos — индекс для записи с конца (m+n-1)
	pos := m + n - 1
	i := m - 1 // указатель на последний реальный элемент first
	j := n - 1 // указатель на последний элемент second

	for j >= 0 {
		// если в first еще есть элементы и они больше, чем current second:
		if i >= 0 && first[i] > second[j] {
			first[pos] = first[i]
			i--
		} else {
			first[pos] = second[j]
			j--
		}
		pos--
	}
	// оставшиеся элементы first уже на местах (если были)
}

// 17. Мажоритарный элемент (Boyer-Moore Voting) — элемент, встречающийся > n/2 раз.
// Предполагается, что такой элемент гарантированно существует.
func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count = 1
		} else if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// 18. Количество подмассивов с суммой равной k (с учётом отрицательных)
// используем префикс-суммы и карту частот
func countSubarraysWithSum(nums []int, k int) int {
	prefixFreq := make(map[int]int)
	prefixFreq[0] = 1 // пустая префикс-сумма
	prefixSum := 0
	count := 0
	for _, v := range nums {
		prefixSum += v
		need := prefixSum - k
		if freq, ok := prefixFreq[need]; ok {
			count += freq
		}
		prefixFreq[prefixSum]++
	}
	return count
}

// 19. Найти повторяющееся число в массиве длины n+1, содержащем числа 1..n,
// ровно одно число дублируется (Floyd's Tortoise and Hare)
func findDuplicateFloyd(nums []int) int {
	// Phase 1: найти встречу в цикле
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	// Phase 2: найти вход в цикл
	ptr1 := nums[0]
	ptr2 := slow
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}

// ---------------- main: тесты из задания ----------------
func main() {
	// 1
	fmt.Println("\n1)  Максимальная прибыль (одна сделка):")
	fmt.Println(maxProfitOneTransaction([]int{7, 1, 5, 3, 6, 4})) // ожидает 5
	fmt.Println(maxProfitOneTransaction([]int{7, 6, 4, 3, 1}))    // ожидает 0

	// 2
	fmt.Println("\n2) Анализ баланса (Кадане):")
	fmt.Println(maxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // ожидает 6 (4,-1,2,1)
	fmt.Println(maxSubarraySum([]int{1, -2, -3, -4}))                 // ожидает 1

	// 3
	fmt.Println("\n3) sortPriorities (Dutch Flag):")
	a1 := []int{2, 0, 1, 0, 2, 1}
	sortPriorities(a1)
	fmt.Println(a1) // ожидает [0,0,1,1,2,2]
	a2 := []int{1, 1, 2, 0, 0, 2}
	sortPriorities(a2)
	fmt.Println(a2)

	// 4
	fmt.Println("\n4) rotateRight:")
	l1 := []int{1, 2, 3, 4, 5}
	rotateRight(l1, 2)
	fmt.Println(l1) // ожидает [4,5,1,2,3]
	l2 := []int{3, 2, 1}
	rotateRight(l2, 1)
	fmt.Println(l2) // ожидает [1,3,2]? (проверим) actually rotating right by 1 -> [1,3,2]

	// 5
	fmt.Println("\n5) findMissingPayment:")
	fmt.Println(findMissingPayment([]int{3, 1, 4, 2, 6, 5})) // N=7 -> missing 7? Но здесь len=6 -> N=7 -> sumExpect 28 sumActual 21 -> 7
	fmt.Println(findMissingPayment([]int{1, 2, 4}))          // len=3 -> N=4 -> missing 3

	// 6
	fmt.Println("\n6) singleEmployeeID:")
	fmt.Println(singleEmployeeID([]int{7, 3, 5, 3, 5, 7, 4})) // ожидает 4
	fmt.Println(singleEmployeeID([]int{2, 2, 1}))             // ожидает 1

	// 7
	fmt.Println("\n7) productExceptSelf:")
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) // ожидает [24,12,8,6]
	fmt.Println(productExceptSelf([]int{2, 3, 5}))    // ожидает [15,10,6]

	// 8
	fmt.Println("\n8) maxConsecutiveDeclines:")
	fmt.Println(maxConsecutiveDeclines([]int{5, 4, 3, 4, 3, 2, 1})) // ожидает 3
	fmt.Println(maxConsecutiveDeclines([]int{1, 2, 3, 2, 1}))       // ожидает 2

	// 9
	fmt.Println("\n9) moveZerosToEnd:")
	u1 := []int{1, 0, 2, 0, 3, 4}
	moveZerosToEnd(u1)
	fmt.Println(u1) // ожидает [1,2,3,4,0,0]
	u2 := []int{0, 0, 1, 2}
	moveZerosToEnd(u2)
	fmt.Println(u2) // ожидает [1,2,0,0]

	// 10
	fmt.Println("\n10) topThreeRatings:")
	fmt.Println(topThreeRatings([]int{1, 2, 2, 3, 3, 3, 4, 4, 5})) // ожидает [3,4,2]
	fmt.Println(topThreeRatings([]int{5, 5, 5, 4, 4, 3}))          // ожидает [5,4,3]

	// 11
	fmt.Println("\n11) dayWithMinCumulativeBalance:")
	idx, bal := dayWithMinCumulativeBalance([]int{7, -3, -10, 4, 2, 8})
	fmt.Println(idx, bal) // индекс минимального накопленного баланса
	idx2, bal2 := dayWithMinCumulativeBalance([]int{-2, -3, -4})
	fmt.Println(idx2, bal2)

	// 12
	fmt.Println("\n12) isValidISBN10:")
	fmt.Println(isValidISBN10("0321146530")) // true
	fmt.Println(isValidISBN10("0131103628")) // true

	// 13
	fmt.Println("\n13) findPeak:")
	pidx, pval := findPeak([]int{1, 2, 3, 1, 4, 2})
	fmt.Println(pidx, pval) // один из пиков, например index 2 value 3
	pidx2, pval2 := findPeak([]int{1, 2, 1})
	fmt.Println(pidx2, pval2)

	// 14
	fmt.Println("\n14) longestSameRun:")
	fmt.Println(longestSameRun([]int{1, 1, 0, 0, 0, 1, 1})) // ожидает 3
	fmt.Println(longestSameRun([]int{0, 1, 0, 1}))          // ожидает 1

	// 15
	fmt.Println("\n15) hasDuplicates:")
	fmt.Println(hasDuplicates([]int{1, 2, 3, 4})) // false
	fmt.Println(hasDuplicates([]int{1, 2, 2, 3})) // true

	//16
	fmt.Println("\n17) Merge two sorted arrays into first (in-place):")
	first := []int{1, 3, 5, 0, 0, 0} // m=3, места для three
	second := []int{2, 4, 6}         // n=3
	fmt.Println("before:", first)
	mergeIntoFirst(first, 3, second, 3)
	fmt.Println("after: ", first) // [1 2 3 4 5 6]

	//17
	fmt.Println("\n18) Majority element (Boyer-Moore):")
	fmt.Println("[2,2,1,1,1,2,2] ->", majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2

	//18
	fmt.Println("\n19) Count subarrays with sum == k:")
	fmt.Println("[1,1,1], k=2 ->", countSubarraysWithSum([]int{1, 1, 1}, 2))             // 2
	fmt.Println("[1,2,3,-1,2], k=3 ->", countSubarraysWithSum([]int{1, 2, 3, -1, 2}, 3)) // несколько примеров

	//19
	fmt.Println("\n20) Find duplicate (Floyd):")
	fmt.Println("[3,1,3,4,2] ->", findDuplicateFloyd([]int{3, 1, 3, 4, 2})) // 3
}

// go run main.go
