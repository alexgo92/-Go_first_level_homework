package square

import "fmt"

type Rectangle struct {
	sideA float64
	sideB float64
}

func (r Rectangle) SquareAB() float64 {
	return r.sideA*r.sideB
}

type SquareFigure interface {
	SquareAB() float64
}

func Square() {
	var sidesRectangle = Rectangle{}
	var rectInterface SquareFigure
	rectInterface = &sidesRectangle
	fmt.Println("Введите длину стороны a прямоугольника")
	fmt.Scanln(&sidesRectangle.sideA)
	fmt.Println("Введите длину стороны b прямоугольника")
	fmt.Scanln(&sidesRectangle.sideB)
	fmt.Printf("Площадь прямоугольника = %.2f\n", rectInterface.SquareAB())
}