package main

import (
	"fmt"
	"math"
	"sort"
)

// Task 1: Слияние интервалов (Объединение встреч)
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}

// Task 2: Вставка интервала (Новая задача)
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}
	return result
}

// Task 3: Непересекающиеся интервалы (Оптимизация)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
	}
	return len(intervals) - count
}

// Task 4: Пересечение интервалов (Общее время)
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])
		if lo <= hi {
			result = append(result, []int{lo, hi})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}

// Task 5: Интервалы планировщика (Ресурсы)
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	maxRooms, rooms, endIdx := 0, 0, 0
	for _, start := range starts {
		if start < ends[endIdx] {
			rooms++
		} else {
			endIdx++
		}
		maxRooms = max(maxRooms, rooms)
	}
	return maxRooms
}

// Task 6: Проверка на пересечение
func isOverlap(interval1, interval2 []int) bool {
	return interval1[0] <= interval2[1] && interval2[0] <= interval1[1]
}

// Task 7: Пересечение двух списков интервалов
// Similar to Task 4, reusing intervalIntersection

// Task 8: Минимальное количество стрел (Кампании)
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	arrows := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			arrows++
			end = points[i][1]
		}
	}
	return arrows
}

// Task 9: Добавление интервалов (Начисление)
// Assuming update intervals with value, but simplified to add value to range in array
func addToIntervals(arr []int, updates [][]int) []int { // updates: [start, end, val]
	diff := make([]int, len(arr)+1)
	for _, u := range updates {
		diff[u[0]] += u[2]
		if u[1]+1 < len(arr) {
			diff[u[1]+1] -= u[2]
		}
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += diff[i]
		arr[i] += sum
	}
	return arr
}

// Task 10: Перезагрузка интервалов (Свободное время)
func findFreeTime(schedule [][]int, workStart, workEnd int) [][]int {
	busy := merge(schedule)
	free := [][]int{}
	prevEnd := workStart
	for _, interval := range busy {
		if interval[0] > prevEnd {
			free = append(free, []int{prevEnd, interval[0]})
		}
		prevEnd = max(prevEnd, interval[1])
	}
	if prevEnd < workEnd {
		free = append(free, []int{prevEnd, workEnd})
	}
	return free
}

// Task 11: Конфликты в расписании
// Find common free time, but task says return intervals when both free.
// Assuming inputs are busy times, find common free.
func commonFree(schedule1, schedule2 [][]int) [][]int {
	union := merge(append(schedule1, schedule2...))
	return findFreeTime(union, 0, math.MaxInt32) // Assuming full day 0 to max
}

// Task 12: Планирование встреч
func canSchedule(intervals [][]int, d int) bool {
	merged := merge(intervals)
	for i := 1; i < len(merged); i++ {
		if merged[i][0]-merged[i-1][1] >= d {
			return true
		}
	}
	return false
}

// Task 13: Загруженность сервера
func maxConcurrentUsers(logs [][]int) int {
	events := []struct{ time, delta int }{}
	for _, log := range logs {
		events = append(events, struct{ time, delta int }{log[0], 1})
		events = append(events, struct{ time, delta int }{log[1], -1})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].time < events[j].time || (events[i].time == events[j].time && events[i].delta < events[j].delta)
	})
	maxCount, count := 0, 0
	for _, e := range events {
		count += e.delta
		maxCount = max(maxCount, count)
	}
	return maxCount
}

// Task 14: Календарь занятости (Summary)
// Similar to merge intervals

// Task 15: Конфликты в бронировании
func hasConflict(bookings [][]int) bool {
	sort.Slice(bookings, func(i, j int) bool {
		return bookings[i][0] < bookings[j][0]
	})
	for i := 1; i < len(bookings); i++ {
		if bookings[i][0] < bookings[i-1][1] {
			return true
		}
	}
	return false
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
	fmt.Println("Task 1:", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]

	// Task 2
	fmt.Println("Task 2:", insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})) // [[1 5] [6 9]]

	// Task 3
	fmt.Println("Task 3:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // 1

	// Task 4
	fmt.Println("Task 4:", intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}})) // [[1 2] [5 5] [8 10] [15 23] [24 24] [25 25]]

	// Task 5
	fmt.Println("Task 5:", minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})) // 2

	// Task 6
	fmt.Println("Task 6:", isOverlap([]int{1, 3}, []int{2, 4})) // true

	// Task 7: Using intervalIntersection
	fmt.Println("Task 7:", intervalIntersection([][]int{{1, 3}, {5, 7}}, [][]int{{2, 6}})) // [[2 3] [5 6]]

	// Task 8
	fmt.Println("Task 8:", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})) // 2

	// Task 9
	arr := []int{0, 0, 0, 0, 0}
	updates := [][]int{{1, 3, 2}}
	fmt.Println("Task 9:", addToIntervals(arr, updates)) // [0 2 2 2 0]

	// Task 10
	fmt.Println("Task 10:", findFreeTime([][]int{{1, 2}, {5, 6}}, 0, 10)) // [[0 1] [2 5] [6 10]]

	// Task 11
	fmt.Println("Task 11:", commonFree([][]int{{1, 3}, {6, 7}}, [][]int{{2, 4}})) // [[0 1] [4 6] [7 2147483647]] (trimmed)

	// Task 12
	fmt.Println("Task 12:", canSchedule([][]int{{0, 30}, {60, 90}}, 20)) // true (30-60 >=20)

	// Task 13
	fmt.Println("Task 13:", maxConcurrentUsers([][]int{{1, 5}, {2, 7}, {4, 5}, {4, 8}})) // 4

	// Task 14: Using merge
	fmt.Println("Task 14:", merge([][]int{{1, 3}, {2, 4}})) // [[1 4]]

	// Task 15
	fmt.Println("Task 15:", hasConflict([][]int{{1, 2}, {3, 4}, {2, 3}})) // true
}
