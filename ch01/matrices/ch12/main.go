package main

import "fmt"

// 26
func addMatrices(A, B [][]int) [][]int {
	m, n := len(A), len(A[0])
	C := make([][]int, m)
	for i := range C {
		C[i] = make([]int, n)
		for j := range C[i] {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

// 27
func subtractMatrices(A, B [][]int) [][]int {
	m, n := len(A), len(A[0])
	C := make([][]int, m)
	for i := range C {
		C[i] = make([]int, n)
		for j := range C[i] {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

// 28
func scalarMultiply(A [][]int, k int) [][]int {
	m, n := len(A), len(A[0])
	C := make([][]int, m)
	for i := range C {
		C[i] = make([]int, n)
		copy(C[i], A[i])
		for j := range C[i] {
			C[i][j] *= k
		}
	}
	return C
}

// 29
func multiplyMatrices(A, B [][]int) [][]int {
	m, p := len(A), len(A[0])
	n := len(B[0])
	if p != len(B) {
		return nil
	}
	C := make([][]int, m)
	for i := range C {
		C[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < p; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// 30
func matrixPower(A [][]int, n int) [][]int {
	if n == 0 {
		size := len(A)
		I := make([][]int, size)
		for i := range I {
			I[i] = make([]int, size)
			I[i][i] = 1
		}
		return I
	}
	if n == 1 {
		return deepCopy(A)
	}
	half := matrixPower(A, n/2)
	res := multiplyMatrices(half, half)
	if n%2 == 1 {
		res = multiplyMatrices(res, A)
	}
	return res
}

// 31
func determinant(A [][]int) int {
	n := len(A)
	if n == 2 {
		return A[0][0]*A[1][1] - A[0][1]*A[1][0]
	}
	if n == 3 {
		return A[0][0]*(A[1][1]*A[2][2]-A[1][2]*A[2][1]) -
			A[0][1]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) +
			A[0][2]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])
	}
	return 0
}

// 32
func trace(A [][]int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i][i]
	}
	return sum
}

// 33
func isDiagonal(A [][]int) bool {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && A[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

// 34
func isIdentity(A [][]int) bool {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i == j && A[i][j] != 1) || (i != j && A[i][j] != 0) {
				return false
			}
		}
	}
	return true
}

// 35
func sumMatrix(A [][]int) int {
	sum := 0
	for i := range A {
		for j := range A[i] {
			sum += A[i][j]
		}
	}
	return sum
}

func deepCopy(A [][]int) [][]int {
	B := make([][]int, len(A))
	for i := range B {
		B[i] = make([]int, len(A[i]))
		copy(B[i], A[i])
	}
	return B
}

func main() {
	// 26. Сложение
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	fmt.Println("26 A + B =", addMatrices(A, B))

	// 27. Вычитание
	fmt.Println("27 A - B =", subtractMatrices(A, B))

	// 28. Умножение на скаляр
	C := scalarMultiply(A, 3)
	fmt.Println("28 A * 3 =", C)

	// 29. Умножение матриц
	X := [][]int{{1, 2}, {3, 4}}
	Y := [][]int{{5, 6}, {7, 8}}
	fmt.Println("29 X × Y =", multiplyMatrices(X, Y))

	// 30. Возведение в степень
	fmt.Println("30 X^3 =", matrixPower(X, 3))

	// 31. Определитель
	fmt.Println("31 det(X) =", determinant(X))

	// 32. След
	fmt.Println("32 trace(X) =", trace(X))

	// 33. Диагональная?
	D := [][]int{{1, 0}, {0, 4}}
	fmt.Println("33 isDiagonal(D):", isDiagonal(D))

	// 34. Единичная?
	I := [][]int{{1, 0}, {0, 1}}
	fmt.Println("34 isIdentity(I):", isIdentity(I))

	// 35. Сумма элементов
	fmt.Println("35 sum(A) =", sumMatrix(A))
}
