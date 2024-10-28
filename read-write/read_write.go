package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodbye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0

	fmt.Println("Commands:\nhello - says hello\nbye - says bye\nq - quit")

	for scanner.Scan() {

		// loop infinitely until user enters 'q' or 'Q'
		if strings.ToLower(scanner.Text()) == "q" {
			break
		}

		text := strings.TrimSpace(scanner.Text())

		switch text {
		case CmdHello:
			numCommands++
			fmt.Println("Response: Hello!")
			break
		case CmdGoodbye:
			numCommands++
			fmt.Println("Response: Bye...")
		}

		if text != "" {
			numLines += 1
		}
	}

	fmt.Printf("You entered %v lines\n", numLines)
	fmt.Printf("You entered %v commands\n", numCommands)
}
