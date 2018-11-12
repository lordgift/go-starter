package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	defer f.Close()
}

func main() {

	go readFile("input.txt")
	readFile("input2.txt")
}
