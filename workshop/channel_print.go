package main

import (
	"fmt"
)

func getInt() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			out <- i
		}
	}()
	return out
}

func main() {
	channel := getInt()
	for i := 1; i <= 100; i++ {
		fmt.Println(<-channel)
	}
}
