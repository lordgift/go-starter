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
	done := make(chan struct{})
	go func() {
		for v := range getInt(1, 100) {
			fmt.Println(v)
		}
		// for i := 1; i <= 100; i++ {
		// 	fmt.Println(<-channel)
		// }
		done <- struct{}{}
	}()

	go func() {
		for v := range getInt(200, 500) {
			fmt.Println(v)
		}
		done <- struct{}{}
	}()

	<-done
	<-done
}
