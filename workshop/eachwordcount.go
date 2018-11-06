package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
$ go run eachwordcount.go < input.txt
*/
func main() {
	var wordCount int
	mapCount := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		words := strings.Fields(line)
		wordCount += len(words)
		for _, w := range words {
			mapCount[w]++
		}
	}

	fmt.Print(mapCount)
}
