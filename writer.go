package main

import (
	"fmt"
	"log"
	"os"
)

/*
create file 'output' by using writer
*/
func main() {
	f, err := os.OpenFile("output", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Hello, World")
	if err != nil {
		log.Fatal(err)
	}
}
