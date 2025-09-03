package main

import (
	"fmt"
	"sync"
)

func add(wg *sync.WaitGroup, x ...[]int) {
	defer wg.Done()
	fmt.Println("Adding numbers...")
	var sum int
	for _, slice := range x {
		for _, v := range slice {
			sum += v
		}
	}
	fmt.Println("Sum:", sum)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // always call Done when finished
	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work
	// time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Create three different slices to demonstrate variadic parameter
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{10, 20, 30}
	slice3 := []int{100, 200, 300, 400}

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	fmt.Println("Slice 3:", slice3)
	fmt.Printf("Expected sum: %d\n", 1+2+3+4+5+10+20+30+100+200+300+400)

	wg.Add(1)
	go add(&wg, slice1, slice2, slice3)
	// launch 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // increment counter
		go worker(i, &wg) // start goroutine
	}

	wg.Wait() // block until counter goes back to 0
	fmt.Println("All workers finished")
}
