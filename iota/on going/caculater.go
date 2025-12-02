package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	DIvide
)

type Operation int

func (op Operation) calculate(x, y int) int {
	switch op {
	case Add:
		return x + y
	case Subtract:
		return x - y
	case Multiply:
		return x * y
	case DIvide:
		return x / y
	}
	panic("aaaaaaaaaaaaaaaaaaaaaaaaaa")
}
func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))
	Subtract := Operation(Subtract)
	fmt.Println(Subtract.calculate(8, 3))
	Multiply := Operation(Multiply)
	fmt.Println(Multiply.calculate(3, 5))
	DIvide := Operation(DIvide)
	fmt.Println(DIvide.calculate(56, 8))
}
