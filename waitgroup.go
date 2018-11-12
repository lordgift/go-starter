package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	wg.Done()
	wg.Done()
	wg.Wait()
	fmt.Println("Exit")
}
