package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Use wg.Go and correctly capture loop variables by creating per-iteration copies.
	greetings := []string{"hello", "greetings", "good day"}
	for _, salutation := range greetings {
		s := salutation
		wg.Go(func() {
			fmt.Println(s)
		})
	}

	// Special case: also capture the index correctly by creating per-iteration copies.
	for idx, salutation := range greetings {
		i := idx
		s := salutation
		wg.Go(func() {
			fmt.Println("s1:", i, s)
		})
	}

	wg.Wait()
}
