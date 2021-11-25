// Калькулятор
package main

import (
	"fmt"
	"math"
	"strconv"
)

// функция рассчета факториала c использованием рекурсии (прочитал дополнительно)
func factorialIs(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorialIs(n-1)
}

func main() {
	fmt.Println("Введите математическую операцию (+, -, * , /, ^, !n)")
	var operation string
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Println("Введите первое число:")
		var firsNumber float64
		fmt.Scanln(&firsNumber)

		fmt.Println("Введите второе число:")
		var secondNumber float64
		fmt.Scanln(&secondNumber)

		fmt.Printf("%.2f\n", firsNumber+secondNumber)
	case "-":
		fmt.Println("Введите первое число:")
		var firsNumber float64
		fmt.Scanln(&firsNumber)

		fmt.Println("Введите второе число:")
		var secondNumber float64
		fmt.Scanln(&secondNumber)

		fmt.Printf("%.2f\n", firsNumber-secondNumber)
	case "*":
		fmt.Println("Введите первое число:")
		var firsNumber float64
		fmt.Scanln(&firsNumber)

		fmt.Println("Введите второе число:")
		var secondNumber float64
		fmt.Scanln(&secondNumber)

		fmt.Printf("%.2f\n", firsNumber*secondNumber)
	case "/":
		fmt.Println("Введите первое число:")
		var firsNumber float64
		fmt.Scanln(&firsNumber)

		fmt.Println("Введите второе число:")
		var secondNumber float64
		fmt.Scanln(&secondNumber)

		if secondNumber == 0 {
			fmt.Println("На ноль делить нельзя")
			break
		}
		fmt.Printf("%.2f\n", firsNumber/secondNumber)
	case "^":
		fmt.Println("Введите число:")
		var number float64
		fmt.Scanln(&number)

		fmt.Println("Введите степень числа:")
		var degreeOf float64
		fmt.Scanln(&degreeOf)

		fmt.Printf("%.2f\n", math.Pow(number, degreeOf))
	case "!n":
		fmt.Println("Введите целое число начиная с 0")
		var str string
		fmt.Scan(&str)

		if num, err := strconv.ParseUint(str, 10, 64); err == nil {
			fmt.Println(factorialIs(num))
		} else {
			fmt.Println("Введите целое положитель число")
		}
	default:
		fmt.Println("Вы ввели неверный оператор!")
	}
}
