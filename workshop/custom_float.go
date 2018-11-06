package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (i MyFloat) power(power int) float64 {
	result := math.Pow(float64(i), float64(power))
	return result
}

func main() {
	a := MyFloat(2.0)

	fmt.Println(a.power(2)) // => 4.0
}
