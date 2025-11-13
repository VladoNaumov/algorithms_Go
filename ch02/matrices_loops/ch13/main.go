package main

import (
	"fmt"
	"math"
)

// 36
func addVectors(a, b []float64) []float64 {
	if len(a) != len(b) {
		return nil
	}
	c := make([]float64, len(a))
	for i := range a {
		c[i] = a[i] + b[i]
	}
	return c
}

// 37
func subtractVectors(a, b []float64) []float64 {
	if len(a) != len(b) {
		return nil
	}
	c := make([]float64, len(a))
	for i := range a {
		c[i] = a[i] - b[i]
	}
	return c
}

// 38
func scalarMultiplyVec(v []float64, k float64) []float64 {
	c := make([]float64, len(v))
	for i, val := range v {
		c[i] = val * k
	}
	return c
}

// 39
func dotProduct(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0
	}
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

// 40
func l2Norm(v []float64) float64 {
	sum := 0.0
	for _, val := range v {
		sum += val * val
	}
	return math.Sqrt(sum)
}

// 41
func l1Norm(v []float64) float64 {
	sum := 0.0
	for _, val := range v {
		sum += math.Abs(val)
	}
	return sum
}

// 42
func lInfNorm(v []float64) float64 {
	max := 0.0
	for _, val := range v {
		abs := math.Abs(val)
		if abs > max {
			max = abs
		}
	}
	return max
}

// 43
func normalize(v []float64) []float64 {
	norm := l2Norm(v)
	if norm == 0 {
		return make([]float64, len(v)) // нулевой вектор
	}
	unit := make([]float64, len(v))
	for i, val := range v {
		unit[i] = val / norm
	}
	return unit
}

// 44
func isOrthogonal(a, b []float64) bool {
	return math.Abs(dotProduct(a, b)) < 1e-9
}

// 45
func angleBetween(a, b []float64) float64 {
	dot := dotProduct(a, b)
	norma := l2Norm(a)
	normb := l2Norm(b)
	if norma == 0 || normb == 0 {
		return 0
	}
	cosTheta := dot / (norma * normb)
	if cosTheta < -1 {
		cosTheta = -1
	}
	if cosTheta > 1 {
		cosTheta = 1
	}
	return math.Acos(cosTheta)
}

// Вспомогательная функция: печать вектора
func printVector(v []float64, name string) {
	fmt.Printf("%s: [", name)
	for i, val := range v {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", val)
	}
	fmt.Println("]")
}

func main() {
	a := []float64{1, 2, 3}
	b := []float64{4, -1, 2}

	// 36. Сложение
	c := addVectors(a, b)
	printVector(c, "36 a + b")

	// 37. Вычитание
	d := subtractVectors(a, b)
	printVector(d, "37 a - b")

	// 38. Умножение на скаляр
	e := scalarMultiplyVec(a, 2.5)
	printVector(e, "38 a * 2.5")

	// 39. Скалярное произведение
	dot := dotProduct(a, b)
	fmt.Printf("39 a · b = %.2f\n", dot)

	// 40. L2-норма
	l2 := l2Norm(a)
	fmt.Printf("40 ||a||₂ = %.2f\n", l2)

	// 41. L1-норма
	l1 := l1Norm(a)
	fmt.Printf("41 ||a||₁ = %.2f\n", l1)

	// 42. L∞-норма
	linf := lInfNorm(a)
	fmt.Printf("42 ||a||∞ = %.2f\n", linf)

	// 43. Нормализация
	normA := normalize(a)
	printVector(normA, "43 normalized a")

	// 44. Ортогональность
	ortho := isOrthogonal(a, b)
	fmt.Printf("44 orthogonal: %v\n", ortho)

	// 45. Угол
	angle := angleBetween(a, b)
	fmt.Printf("45 angle: %.2f rad (%.1f°)\n", angle, angle*180/math.Pi)
}
