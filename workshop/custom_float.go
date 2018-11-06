package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f *MyFloat) power(power int) {
	*f = MyFloat(math.Pow(float64(*f), float64(power)))
}

func main() {
	a := MyFloat(2.0)
	a.power(2)
	fmt.Println(a) // => 4.0
}
