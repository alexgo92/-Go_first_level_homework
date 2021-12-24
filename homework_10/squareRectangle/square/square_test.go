package square

import (
	"fmt"
	"math"
	"testing"
)

func isCloseEnough (a float64, b float64) bool {
	const accuracy = 0.0001
	return math.Abs(a - b) < accuracy
}

func TestRectangle_SquareAB(t *testing.T) {
	TestCase := []struct {
		name string
		input Rectangle
		expected float64
	} {
		{
			name: "case one",
			input: Rectangle{
				sideA: 2.35,
				sideB: 3.45,
			},
			expected: 8.1075,
		},
		{
			name: "case two",
			input: Rectangle{
				sideA: 7.01,
				sideB: 3.00,
			},
			expected: 21.03,
		},
	}
	for _, cse := range TestCase {
		t.Run(cse.name, func(t *testing.T) {
			testSquare := cse.input.SquareAB()
			if !isCloseEnough(cse.expected, testSquare) {
				t.Errorf("expected %.4f got %.4f", cse.expected, testSquare)
			}
		})
	}
}

func BenchmarkSquareAB(b *testing.B) {
	var sidesRectangle = Rectangle{}
	sidesRectangle.sideA = 5.43
	sidesRectangle.sideB = 4.34
	for i := 0; i < b.N; i++ {
		sidesRectangle.SquareAB()
	}
}

func ExampleRectangle_SquareAB() {
	// объявляем переменную типа структура
	var sidesRectangle = Rectangle{}
	// объявляем переменную типа интерфейс, методы структуры удовлетворяют интерфейсу
	var rectInterface SquareFigure
	// инициализируем интерфей как указатель на структуру
	rectInterface = &sidesRectangle
	// считываем длину стороны а
	fmt.Println("Введите длину стороны a прямоугольника")
	fmt.Scanln(&sidesRectangle.sideA)
	//считываем длину стороны b
	fmt.Println("Введите длину стороны b прямоугольника")
	fmt.Scanln(&sidesRectangle.sideB)

	//используем функцию SquareAB() для рассчета площади прямоугольника
	//затем выводим результат в консоль
	fmt.Printf("Площадь прямоугольника = %.2f\n", rectInterface.SquareAB())
}
