package main

import "fmt"

func main() {
	//3. С клавиатуры вводится трехзначное число.
	//Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	var number, hundreds, dozens, units int
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&number)
	hundreds = number / 100
	dozens = number / 10 % 10
	units = number % 10
	fmt.Println("Количество сотен =", hundreds)
	fmt.Println("Количество десятков =", dozens)
	fmt.Println("Количество единиц =", units)
}
