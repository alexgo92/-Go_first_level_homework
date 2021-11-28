package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	integers := []int64{}

	for {
		var integer int64
		//конструкцию if взял из урока но не смог разобраьтся почему тут знак ";"
		//и как он работает
		//эта конструкция тоже не понятна: "errors.Is(err, io.EOF)"
		//буду благодарен за пояснение
		if _, err := fmt.Scanln(&integer); errors.Is(err, io.EOF) {
			fmt.Println("Exit")
			break
		}
		integers = append(integers, integer)
	}

	// insert-sort
	for i := 1; i < len(integers); i++ {
		for j := i; j > 0 && (integers[j] < integers[j-1]); j-- {
			integers[j-1], integers[j] = integers[j], integers[j-1]
		}
	}
	fmt.Println(integers)
}
