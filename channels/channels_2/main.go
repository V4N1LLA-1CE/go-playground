package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%#v\n", sleepSort([]int{1, 5, 6, 2, 3}))
}

/*
for every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- send "n" over a channel
*/
func sleepSort(values []int) []int {
	ch := make(chan int)
	out := []int{}

	for _, n := range values {
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	for range values {
		out = append(out, <-ch)
	}

	return out
}
