package main

import "testing"

// Тестируем функцию генерации данных
func TestGenerateRandomElements(t *testing.T) {
	size := 404
	result, _ := generateRandomElements(size)

	if len(result) != size {
		t.Errorf("wrong slice size: want %d, have %d", size, len(result))
	}
}

// Тестируем функцию нахождения максимума
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"maxPositive", []int{1, 3, 2, 10, 5}, 10},
		{"maxNegative", []int{-10, -20, -3, -100}, -3},
		{"allEqual", []int{7, 7, 7}, 7},
		{"oneElement", []int{42}, 42},
		{"emptySlice", []int{}, 0},
	}
	for _, nn := range tests {
		t.Run(nn.name, func(t *testing.T) {
			got, _ := maximum(nn.data)
			if got != nn.want {
				t.Errorf("maximum() = %d, want %d", got, nn.want)
			}
		})
	}
}
