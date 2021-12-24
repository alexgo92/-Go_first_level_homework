package bubbleSort

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}

func TestNumbers_BubbleSort(t *testing.T) {
	testCase := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "simple case one",
			input:    []int{4, 1, 3, 8},
			expected: []int{1, 3, 4, 8},
		},
		{
			name:     "simple case two",
			input:    []int{4, 6, 2, 4, 0},
			expected: []int{0, 2, 4, 4, 6},
		},
	}

	for _, cse := range testCase {
		t.Run(cse.name, func(t *testing.T) {
			testNums := BubbleSort(cse.input)
			if !Equal(testNums, cse.expected) {
				t.Errorf("expected %v got %v", cse.expected, testNums)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{3,6,8,1,3,23,65,23,1})
	}
}

func ExampleNumbers_BubbleSort() {
	// слайс вводим с клавиатуры, чтобы закончить ввод слайса жмем "Ctrl + D"
	// тут уже указан готовый слайс для наглядности
	sliceOfNumbers := []int{1,5,0,4,2}

	sliceOfNumbersSorted := BubbleSort(sliceOfNumbers)

	fmt.Println("Отсартированный слайс чисел", sliceOfNumbersSorted)
}
