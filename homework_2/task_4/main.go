package main

import (
	"fmt"
)

func main() {
	//Написать приложение, которое ищет все простые числа от 0 до N
	//включительно. Число N должно быть задано из стандартного потока
	//ввода.
	var number int
	fmt.Println("Введите число")
	fmt.Scanln(&number)
	flag := false
	for i := 2; i <= number; i++ {
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				flag = true
			}
		}
		if !flag {
			fmt.Println("Простое число из заданного диапазона:", i)
		}
		flag = false
	}
	//сложность вашего алгоритма - сколько операций нужно выполнить,
	//чтобы найти все простые числа от 0 до N?
	/*
		На этот вопрос затрудняюсь ответить,
		буду благодарен за пояснение
	*/
}
