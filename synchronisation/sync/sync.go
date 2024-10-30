package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	counter int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetters(word string) int {
	letters := 0
	for _, char := range word {
		if unicode.IsLetter(char) {
			letters += 1
		}
	}
	return letters
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	totalLetters := Count{}
	var wg sync.WaitGroup

	fmt.Println("Start typing!")
	fmt.Println("[LINUX] - End typing and count letters with `ctrl + d`")
	fmt.Println("[MAC] - End typing and count letters with `ctrl^ + d`")
	fmt.Println("[WINDOWS] - End typing and count letters with `ctrl + z`")

	for {

		if s.Scan() {
			line := s.Text()
			words := getWords(line)

			for _, word := range words {
				wordCopy := word
				wg.Add(1)

				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()
					sum := countLetters(wordCopy)
					totalLetters.counter += sum
				}()

			}
		} else {
			break
		}
	}

	wg.Wait()

	totalLetters.Lock()
	result := totalLetters.counter
	totalLetters.Unlock()

	fmt.Printf("\nYou have typed in %v letters.\n", result)
}
