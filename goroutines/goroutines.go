package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return sum
		}
		if err != nil {
			fmt.Print("Error:", err)
		}
		// remove newline
		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Error:", err)
		}

		sum += num
	}
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt"}
	sum := 0

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		rd := bufio.NewReader(file)

		calculate := func() {
			fileSum := sumFile(*rd)
			sum += fileSum
		}

		go calculate()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(sum)

}
