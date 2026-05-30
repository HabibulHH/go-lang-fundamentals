package main

import (
	"fmt"
	"sync"
	"time"
)

// Sequential Processing
//         ↓
// One Goroutine Per Job
//         ↓
// Problem: Too Many Goroutines
//         ↓
// Need Controlled Concurrency (Worker Pool)
//         ↓
// Job Queue
//         ↓
// Workers
//         ↓
// Worker Pool

func worker(id int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d Processed %s\n", id, url)
		results <- url
	}
}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()

	urls := []string{
		"url1", "url2", "url3", "url4", "url5", "url6",
	}

	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	const workerCount = 2

	for workerId := 1; workerId <= workerCount; workerId++ {
		wg.Go(func() { worker(workerId, jobs, results) })
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	wg.Wait()

	close(results)

	for result := range results {
		fmt.Printf("Recieved result: %s\n", result)
	}

	fmt.Printf("It took %s ms", time.Since(startTime))
}
