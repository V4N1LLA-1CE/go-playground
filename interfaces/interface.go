package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

func main() {
	c := Car("Audi R8")
	t := Truck("Range Rover")
	m := Motorcycle("Yahama R9")

	sendToLift(c)
	sendToLift(t)
	sendToLift(m)
}

type Lifter interface {
	PickLift() int
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcycle) PickLift() int {
	return SmallLift
}

func (c Car) PickLift() int {
	return StandardLift
}

func (t Truck) PickLift() int {
	return LargeLift
}

func sendToLift(l Lifter) {
	switch l.PickLift() {
	case SmallLift:
		fmt.Printf("Send %v to small lift\n", l)
	case StandardLift:
		fmt.Printf("Send %v to standard lift\n", l)
	case LargeLift:
		fmt.Printf("Send %v to large lift\n", l)
	}
}
