package main

import (
	"bufio"
	"fmt"
	"strings"
)

/*
create file 'output' by using writer
*/
func main() {
	str := "Go is an open source programming language \nthat makes it easy to build simple, reliable, \nand efficient software."

	scanner := bufio.NewScanner(strings.NewReader(str))
	for lineNO := 1; scanner.Scan(); lineNO++ {
		line := scanner.Text()
		fmt.Printf("%d:%s\n", lineNO, line)
	}
}
