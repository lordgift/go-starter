package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1

	//5th channel cant add because channel buffer = 4
	// ch <- 1
	fmt.Println(<-ch)
}
