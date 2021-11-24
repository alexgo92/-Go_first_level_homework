package main

import (
	"fmt"
	"math"
)

func main() {
	//2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
	//Площадь круга должна вводиться пользователем с клавиатуры.
	var s float64
	fmt.Println("Введите площадь круга")
	fmt.Scanln(&s)
	//диаметр
	var d float64
	d = 2 * math.Sqrt(s/math.Pi)
	fmt.Printf("Диаметр окружности = %.2f\n", d)
	//длина
	var l float64
	l = math.Sqrt(s * 4 * math.Pi)
	fmt.Printf("Длина окружности = %.2f\n", l)
}
