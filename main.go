package main

import (
	"fmt"
)

func main() {
	//shorthand declaration
	name := "John"
	fmt.Println("Hello, ", name)

	//inline declaration
	if inlineDeclare := "john"; inlineDeclare != "I'm Groot" {
		fmt.Println(inlineDeclare + " not equal to 'I'm Groot'")
	}

	//switch-case is auto break (use fallthrough to continue if need)
	switch "A" {
	case "A":
		fmt.Println("This is A case")
		fallthrough
	case "B":
		fmt.Println("This is B case")
	case "C":
		fmt.Println("This is C case")
	default:
		fmt.Println("default case")
	}
}
