package main

import (
	"fmt"
)

func getInt(begin int, end int) chan int {
	out := make(chan int)
	go func() {
		for i := begin; i <= end; i++ {
			out <- i
		}
		//close channel
		close(out)
	}()
	return out
}

func main() {
	begin, end := 1, 100
	channel := getInt(begin, end)
	for v := range channel {
		fmt.Println(v)
	}
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(<-channel)
	// }
}
