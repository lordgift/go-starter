package main

import (
	"fmt"
	"time"
)

type MyInt int

func main() {

	myInt := MyInt(5)

	now := time.Now()
	fmt.Println(now)

}
