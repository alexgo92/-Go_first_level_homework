package main

import "fmt"

func main() {
	//1. Напишите программу для вычисления площади прямоугольника.
	//Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
	var a, b float64
	fmt.Println("Введите длину стороны a прямоугольника")
	fmt.Scanln(&a)
	fmt.Println("Введите длину стороны b прямоугольника")
	fmt.Scanln(&b)
	fmt.Printf("Площадь прямоугольника = %.2f\n", a*b)
}
