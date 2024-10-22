package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// pointer to structure doesn't need dereferencing
	counter.hits += 1
}

func replace(old *string, new string, counter *Counter) {
	// deference before assigning value
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}
	greeting := "Hello World"

	fmt.Printf("Counter before(%p): %d\n", &counter, counter.hits)
	fmt.Printf("Greeting before(%p): %s\n", &greeting, greeting)

	replace(&greeting, "Hi!", &counter)

	fmt.Printf("Counter after(%p): %d\n", &counter, counter.hits)
	fmt.Printf("Greeting after(%p): %s\n", &greeting, greeting)

}
