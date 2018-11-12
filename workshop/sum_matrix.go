package main

import (
	"fmt"
)

func SumMatrix(a [][]float64, b [][]float64) [][]float64 {
	//initiate matrix c
	c := make([][]float64, len(a))
	for i := range c {
		c[i] = make([]float64, len(a[i]))
	}

	//sum matrix
	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] + b[i][j]
			fmt.Printf("c[%d][%d] => %.0f \n", i, j, c[i][j])
		}
	}
	fmt.Println("------------")
	return c
}

func main() {
	a := [][]float64{
		{1, 2},
		{3, 4},
	}

	b := [][]float64{
		{5, 6},
		{7, 8},
	}

	SumMatrix(a, b)

}
