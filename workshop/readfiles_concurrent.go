package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string, done chan bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	defer f.Close()
	done <- true
}

func main() {
	done := make(chan bool)
	go readFile("input.txt", done)
	go readFile("input2.txt", done)

	//try to get value from channel (for check it done)
	<-done
	<-done
}
