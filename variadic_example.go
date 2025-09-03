package main

import (
	"fmt"
	"strings"
	"sync"
)

// Function that accepts variadic parameter ...[]int
func addSlices(wg *sync.WaitGroup, name string, slices ...[]int) {
	defer wg.Done()

	fmt.Printf("\n=== %s ===\n", name)
	fmt.Printf("Number of slices received: %d\n", len(slices))

	// Show what we received
	for i, slice := range slices {
		fmt.Printf("Slice %d: %v\n", i+1, slice)
	}

	// Calculate sum
	var total int
	for sliceIndex, slice := range slices {
		var sliceSum int
		for _, value := range slice {
			sliceSum += value
			total += value
		}
		fmt.Printf("Slice %d sum: %d\n", sliceIndex+1, sliceSum)
	}

	fmt.Printf("Total sum: %d\n", total)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Demonstrating variadic parameter ...[]int\n")

	// Example 1: No slices
	fmt.Println("Example 1: Calling with no slices")
	wg.Add(1)
	go addSlices(&wg, "No Slices")

	// Example 2: One slice
	fmt.Println("\nExample 2: Calling with one slice")
	wg.Add(1)
	go addSlices(&wg, "One Slice", []int{1, 2, 3, 4, 5})

	// Example 3: Multiple slices
	fmt.Println("\nExample 3: Calling with multiple slices")
	slice1 := []int{10, 20, 30}
	slice2 := []int{100, 200}
	slice3 := []int{1000, 2000, 3000, 4000}

	wg.Add(1)
	go addSlices(&wg, "Multiple Slices", slice1, slice2, slice3)

	// Example 4: Mixed approach - some direct, some variables
	fmt.Println("\nExample 4: Mixed approach")
	wg.Add(1)
	go addSlices(&wg, "Mixed",
		[]int{5, 15, 25},
		slice1,
		[]int{99, 101})

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Key Points about ...[]int:")
	fmt.Println("• The ... means 'zero or more'")
	fmt.Println("• []int means 'slice of integers'")
	fmt.Println("• So ...[]int means 'zero or more slices of integers'")
	fmt.Println("• Inside the function, it becomes [][]int (slice of slices)")
	fmt.Println("• You can pass any number of []int arguments")
	fmt.Println(strings.Repeat("=", 50))
}
