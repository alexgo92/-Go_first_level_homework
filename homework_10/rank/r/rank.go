package rank

import "fmt"

func NumberRank(number int) map[string]int {
	rankMap := make(map[string]int, 3)
	hundreds := number / 100
	rankMap["hundreds"] = hundreds
	dozens := number / 10 % 10
	rankMap["dozens"] = dozens
	units := number % 10
	rankMap["units"] = units
	return rankMap
}

func Rank() {
	//3. С клавиатуры вводится трехзначное число.
	//Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	var number int
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&number)
	rankMap := NumberRank(number)
	fmt.Println("Количество сотен =", rankMap["hundreds"])
	fmt.Println("Количество десятков =", rankMap["dozens"])
	fmt.Println("Количество единиц =", rankMap["units"])
}