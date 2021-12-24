package insertion

import (
	"errors"
	"fmt"
	"io"
)

func InsertSort(slice []int) []int{
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j] < slice[j-1]; j-- {
			slice[j-1], slice[j] = slice[j], slice[j-1]
		}
	}
	return slice
}

func Sort() {
	var num int
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Exit")
		return
	}
	var numbers []int
	numbers = append(numbers, num)

	for {
		_, err := fmt.Scanln(&num)
		if errors.Is(err, io.EOF) {
			fmt.Println("Sorted")
			break
		}
		numbers = append(numbers, num)
	}

	sortedSliceNumbers := InsertSort(numbers)
	fmt.Println(sortedSliceNumbers)
}
