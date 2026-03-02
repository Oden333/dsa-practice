package structs

import (
	"fmt"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			input:    []int{2, 1, 2, 4, 3},
			expected: []int{4, 2, 4, -1, -1},
		},
		{
			name:     "базовый пример",
			input:    []int{1, 3, 4, 2},
			expected: []int{3, 4, -1, -1},
		},
		{
			name:     "убывающий массив - все -1",
			input:    []int{4, 3, 2, 1},
			expected: []int{-1, -1, -1, -1},
		},
		{
			name:     "возрастающий массив - все находят ответ",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 3, 4, -1},
		},
		{
			name:     "дубликаты",
			input:    []int{2, 1, 2, 4},
			expected: []int{4, 2, 4, -1},
		},
		{
			name:     "один элемент",
			input:    []int{7},
			expected: []int{-1},
		},
		{
			name:     "пустой массив",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "все одинаковые",
			input:    []int{5, 5, 5},
			expected: []int{-1, -1, -1},
		},
		{
			name:     "чередование",
			input:    []int{5, 1, 4, 2, 3},
			expected: []int{-1, 4, -1, 3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Копируем входные данные, чтобы тесты были изолированы
			// (на случай если функция мутирует входной слайс)
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			got := nextGreaterElement(inputCopy)

			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.expected) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.expected)
			}
		})
	}
}
func TestNextLowerElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			input:    []int{2, 1, 2, 4, 3},
			expected: []int{1, -1, -1, 3, -1},
		},

		{
			input:    []int{2, 1, 5, 6, 2, 3},
			expected: []int{1, -1, 2, 2, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			got := nextLowerElement(inputCopy)

			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.expected) {
				t.Errorf("nextLowerElement() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkNextGreaterElement(b *testing.B) {
	// Генерируем массив разного размера
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			// Худший случай: убывающий массив (стек растёт до n)
			nums := make([]int, size)
			for i := range nums {
				nums[i] = size - i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = nextGreaterElement(nums)
			}
		})
	}
}
