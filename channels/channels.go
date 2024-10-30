package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	data int
}

func runJob(ch chan<- int, j Job) {
	ch <- square(j)
}

func square(j Job) int {
	return j.data * j.data
}

func makeJobs() []Job {
	res := []Job{}
	for i := 0; i < 30; i++ {
		res = append(res, Job{i})
	}
	return res
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	jobs := makeJobs()
	results := make(chan int)

	for _, job := range jobs {
		go runJob(results, job)
	}

	resultCount := 0
	resultsValue := []int{}
	for {
		result := <-results
		resultsValue = append(resultsValue, result)
		resultCount++
		if resultCount == len(jobs) {
			break
		}
	}

	for _, res := range resultsValue {
		fmt.Println("Job complete:", res)
	}
}
