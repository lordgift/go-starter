package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
$ go run wordcount.go < input.txt
*/
func main() {
	var wordCount int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// wordCount += len(strings.Split(line, " "))
		wordCount += len(strings.Fields(line))
	}

	fmt.Print(wordCount)
}
