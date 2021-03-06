package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/**
same as eachwordcount.go but read file (no need input redirect via CLI)
*/
func main() {
	var wordCount int
	mapCount := make(map[string]int)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
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
