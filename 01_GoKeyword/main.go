package main

import (
	"fmt"
	"sync"
)

func alpha(id int) {
	fmt.Println("Alpha", id)
}

func beta(id int, msg string) {
	fmt.Println("Beta", id, msg)
}

func gamma(name string) {
	fmt.Println("Gamma", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Go(func() { alpha(1) })
	wg.Go(func() { beta(2, "hello") })
	wg.Go(func() { gamma("gopher") })

	wg.Wait()
}
