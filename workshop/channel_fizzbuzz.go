package main

import (
	"fmt"
)

func genFizzBuzz(begin, end int) chan string {
	out := make(chan string)
	go func() {
		for i := begin; i <= end; i++ {
			switch {
			case i%15 == 0:
				out <- "FizzBuzz"
			case i%3 == 0:
				out <- "Fizz"
			case i%5 == 0:
				out <- "Buzz"
			default:
				out <- fmt.Sprint(i)
			}
		}
		//close channel
		close(out)
	}()
	return out
}

func genEven(begin, end int) chan int {
	out := make(chan int)
	go func() {
		for i := begin; i <= end; i++ {
			if i%2 == 0 {
				out <- i
			}
		}
		//close channel
		close(out)
	}()
	return out
}

func genOdd(begin, end int) chan int {
	out := make(chan int)
	go func() {
		for i := begin; i <= end; i++ {
			if i%2 == 1 {
				out <- i
			}
		}
		//close channel
		close(out)
	}()
	return out
}

func main() {
	for v := range genFizzBuzz(1, 100) {
		fmt.Println(v)
	}
	fmt.Println("--------------------")
	for v := range genEven(1, 100) {
		fmt.Println(v)
	}
	fmt.Println("--------------------")
	for v := range genOdd(1, 100) {
		fmt.Println(v)
	}
	// done := make(chan struct{})
	// go func() {
	// 	for v := range getInt(1, 100) {
	// 		fmt.Println(v)
	// 	}
	// 	done <- struct{}{}
	// }()
}
