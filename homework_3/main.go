// Калькулятор
package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

// функция рассчета факториала c использованием рекурсии (прочитал дополнительно)
func factorialIs(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialIs(n-1)
}

// функция проверки числа на целочисленность
func isInt(str string) bool {
	for _, s := range str {
		if !unicode.IsDigit(s) {
			return false
		}
	}
	return true
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
		} else {
			fmt.Printf("%.2f\n", firsNumber/secondNumber)
		}
	case "^":
		fmt.Println("Введите число:")
		var number float64
		fmt.Scanln(&number)

		fmt.Println("Введите степень числа:")
		var degreeOf float64
		fmt.Scanln(&degreeOf)

		fmt.Printf("%.2f\n", math.Pow(number, degreeOf))
	case "!n":
		fmt.Println("Введите целое число больше 0")
		var str string
		fmt.Scan(&str)

		//  проверка на целочисленность
		if isInt(str) {
			number, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Введите целое число больше 0")
			} else if number <= 0 {
				fmt.Println("Число должно быть больше 0")
			} else {
				fmt.Println("Факториал =", factorialIs(number))
			}
		} else if str[0] == 45 {
			fmt.Println("Число должно быть положительным")
		} else {
			fmt.Println("Должно быть целое число")
		}
	default:
		fmt.Println("Вы ввели неверный оператор!")
	}
}
