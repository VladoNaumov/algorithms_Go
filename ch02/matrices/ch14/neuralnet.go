package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ------------------- ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ -------------------

// ReLU и производная
func relu(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}
func dRelu(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}

// Sigmoid и производная
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
func dSigmoid(x float64) float64 {
	return x * (1.0 - x)
}

// ------------------- НЕЙРОСЕТЬ -------------------

type NeuralNet struct {
	W1, W2 [][]float64 // веса: W1[input][hidden], W2[hidden][1]
	B1, B2 []float64   // смещения
}

func NewNeuralNet(input, hidden, output int) *NeuralNet {
	rand.Seed(time.Now().UnixNano())

	// Инициализация весов
	W1 := make([][]float64, input)
	for i := range W1 {
		W1[i] = make([]float64, hidden)
		for j := range W1[i] {
			W1[i][j] = rand.NormFloat64() * 0.1
		}
	}

	W2 := make([][]float64, hidden)
	for i := range W2 {
		W2[i] = make([]float64, output)
		for j := range W2[i] {
			W2[i][j] = rand.NormFloat64() * 0.1
		}
	}

	B1 := make([]float64, hidden)
	B2 := make([]float64, output)

	return &NeuralNet{W1: W1, W2: W2, B1: B1, B2: B2}
}

// Прямой проход
func (nn *NeuralNet) Forward(x []float64) (float64, []float64, []float64, []float64, []float64) {
	// Скрытый слой: z1 = W1·x + b1
	z1 := make([]float64, len(nn.B1))
	for i := range z1 {
		z1[i] = nn.B1[i]
		for j := range x {
			z1[i] += x[j] * nn.W1[j][i]
		}
	}
	a1 := make([]float64, len(z1))
	for i := range a1 {
		a1[i] = relu(z1[i])
	}

	// Выходной слой: z2 = W2·a1 + b2
	z2 := nn.B2[0]
	for i := range a1 {
		z2 += a1[i] * nn.W2[i][0]
	}
	a2 := []float64{sigmoid(z2)}

	return a2[0], z1, a1, nil, a2 // z2 не нужен
}

// Обучение (один шаг) — ИСПРАВЛЕННЫЙ
func (nn *NeuralNet) Train(x []float64, y float64, lr float64) {
	// Прямой проход
	pred, z1, a1, _, a2 := nn.Forward(x)

	// 1. Градиент ошибки по предсказанию
	dLoss := 2 * (pred - y)

	// 2. Градиент по z2 (выход)
	dZ2 := dLoss * dSigmoid(a2[0])

	// 3. Обновляем W2 и B2
	for i := range nn.W2 {
		nn.W2[i][0] -= lr * dZ2 * a1[i]
	}
	nn.B2[0] -= lr * dZ2

	// 4. Градиент по a1
	dA1 := make([]float64, len(a1))
	for i := range dA1 {
		dA1[i] = dZ2 * nn.W2[i][0]
	}

	// 5. Градиент по z1
	dZ1 := make([]float64, len(z1))
	for i := range dZ1 {
		dZ1[i] = dA1[i] * dRelu(z1[i])
	}

	// 6. Обновляем W1 и B1
	for i := range nn.W1 {
		for j := range nn.W1[i] {
			nn.W1[i][j] -= lr * dZ1[j] * x[i]
		}
	}
	for i := range nn.B1 {
		nn.B1[i] -= lr * dZ1[i]
	}
}

// ------------------- ОБУЧЕНИЕ НА XOR -------------------

func main() {
	nn := NewNeuralNet(2, 4, 1)

	// Данные XOR
	inputs := [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	targets := []float64{0, 1, 1, 0}

	fmt.Println("Обучение нейросети на XOR...")

	for epoch := 0; epoch < 10000; epoch++ {
		for i := range inputs {
			nn.Train(inputs[i], targets[i], 0.1)
		}

		// Выводим loss каждые 2000 эпох
		if epoch%2000 == 0 {
			loss := 0.0
			for i := range inputs {
				pred, _, _, _, _ := nn.Forward(inputs[i])
				loss += (pred - targets[i]) * (pred - targets[i])
			}
			fmt.Printf("Эпоха %d, Loss: %.6f\n", epoch, loss/4)
		}
	}

	// Тестирование
	fmt.Println("\nРезультаты:")
	for i := range inputs {
		pred, _, _, _, _ := nn.Forward(inputs[i])
		fmt.Printf("%v → %.3f (ожидается %.0f)\n", inputs[i], pred, targets[i])
	}
}
