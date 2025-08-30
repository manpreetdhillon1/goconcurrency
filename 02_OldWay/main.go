package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // always call Done when finished
	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work
	// time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// launch 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // increment counter
		go worker(i, &wg) // start goroutine
	}

	wg.Wait() // block until counter goes back to 0
	fmt.Println("All workers finished")
}
