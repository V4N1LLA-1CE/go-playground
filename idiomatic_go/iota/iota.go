package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {

	case Add:
		return lhs + rhs

	case Subtract:
		return lhs - rhs

	case Multiply:
		return lhs * rhs

	case Divide:
		return lhs / rhs

	}

	panic("Unhandled operation")
}

func main() {
	add := Operation(Add)
	subtract := Operation(Subtract)
	multiply := Operation(Multiply)
	divide := Operation(Divide)

	fmt.Println(add.calculate(2, 2))
	fmt.Println(multiply.calculate(3, 2))
	fmt.Println(divide.calculate(16, 2))
	fmt.Println(subtract.calculate(14, 4))
}
