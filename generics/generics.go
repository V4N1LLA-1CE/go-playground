package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numbers interface {
	constraints.Float | constraints.Integer
}

func clamp[T Numbers](value, min, max T) T {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func main() {
	fmt.Printf("%v\n", clamp(5.23, 1, 20))
	fmt.Printf("%v\n", clamp(5.23, 5.3, 20))
}
