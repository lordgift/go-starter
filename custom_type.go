package main

import (
	"fmt"
	"time"
)

type MyInt int

/* func (<receiver>) <method_name>(<parameter>) <return_type> */
func (i MyInt) add(n MyInt) MyInt {
	return i + n
}

func main() {

	myInt := MyInt(5)
	myInt.add(20)

	now := time.Now()
	fmt.Println(now)

}
