// 1. Напишите приложение,
// рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// 2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

var FibMap = map[int64]int64{
	0: 0,
	1: 1,
}

func numberFibonacci(num int64) int64 {
	if num == 0 {
		return 0
	}
	if _, ok := FibMap[num]; ok {
		FibMap[num+1] = FibMap[num] + FibMap[num-1]
		return FibMap[num]
	}
	return numberFibonacci(num-2) + numberFibonacci(num-1)
}

func main() {
	defer fmt.Println("Goodbye")

	for {
		fmt.Println("(if you want exit -> please, enter 'Ctrl + D')\n" +
			"Please, enter the index number in the Fibonacci system:")

		var number string
		if _, err := fmt.Scanln(&number); errors.Is(err, io.EOF) {
			fmt.Println("Exit")
			break
		}

		var num int64
		num, err := strconv.ParseInt(number, 10, 64)

		if err != nil {
			fmt.Println(err)
			continue
		}

		switch {
		// для положительных индексов
		case num >= 0:
			fmt.Printf("Fibonacci number = %v\n", numberFibonacci(num))
			continue
		// для отрицательных индексов
		case num < 0:
			if num%2 == 0 {
				fmt.Printf("Fibonacci number = %v\n", (numberFibonacci(-num) * (-1)))
			} else {
				fmt.Printf("Fibonacci number = %v\n", numberFibonacci(-num))
			}
			continue
		}
	}
}
