package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	// simulating traffic over the network
	wait()

	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()

	hits.count += 1
	fmt.Println("server iteration", iteration)
}

type Hits struct {
	count int
	sync.Mutex
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup
	hitCounter := Hits{}

	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Printf("waiting for goroutines...\n\n")
	wg.Wait()
	fmt.Println("done!")

	hitCounter.Lock()
	fmt.Printf("Hit Counter: {count: %v}\n", hitCounter.count)
	hitCounter.Unlock()

}
