package main

import "fmt"

// variadic parameter
func sum(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	// combine into array
	combined := append(a, b...)

	// break array into (a, b, c...)
	fmt.Println(sum(combined...))
}
