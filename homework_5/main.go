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

func numberFibonacci(num int64) int64 {
	if num == 1 {
		return 1
	} else if num == 0 {
		return 0
	}
	return numberFibonacci(num-2) + numberFibonacci(num-1)
}

func main() {
	defer fmt.Println("Goodbye")

	valuesFib := map[int64]int64{}

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

		if _, ok := valuesFib[num]; ok {
			fmt.Printf("Fibonacci number = %v\n", valuesFib[num])
			continue
		} else {
			switch {
			// для положительных индексов
			case num >= 0:
				valuesFib[num] = numberFibonacci(num)
				fmt.Printf("Fibonacci number = %v\n", valuesFib[num])
				continue
			// для отрицательных индексов
			case num < 0:
				if num%2 == 0 {
					valuesFib[num] = numberFibonacci(-num) * (-1)
					fmt.Printf("Fibonacci number = %v\n",
						valuesFib[num])
				} else {
					valuesFib[num] = numberFibonacci(-num)
					fmt.Printf("Fibonacci number = %v\n", valuesFib[num])
				}
				continue
			}
		}
	}
}
