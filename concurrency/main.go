package main

import (
	"fmt"
	"time"
)

func Counter() {
	for i := 1; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	//go ==> new goroutine
	go Counter()
	go Counter()

	time.Sleep(11 * time.Second)
}
