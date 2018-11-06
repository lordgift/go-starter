package main

import (
	"fmt"
	"time"
)

/*
loop & print now time every second.
*/
func main() {
	for { //infinite loop
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
