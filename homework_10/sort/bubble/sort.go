package bubbleSort

import (
	"errors"
	"fmt"
	"io"
)

func BubbleSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		for j := 1; j < len(numbers); j++ {
			if numbers[j-1] > numbers [j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
		}
	}
	return numbers
}

func Sort() {
	var num int
	if _, err := fmt.Scanln(&num); errors.Is(err, io.EOF) {
		fmt.Println("Exit")
		return
	}
	var numbers = []int{num}
	for {
		if _, err := fmt.Scanln(&num); errors.Is(err, io.EOF) {
			fmt.Println("Sort:")
			break
		}
		numbers = append(numbers, num)
	}
	sortedNumbers := BubbleSort(numbers)
	fmt.Println(sortedNumbers)
}
