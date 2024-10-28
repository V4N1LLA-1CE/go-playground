package main

import "fmt"

func init() {
	fmt.Println("Init function: Runs before main")
}

func main() {
	fmt.Println("Main function: Run main")
}
