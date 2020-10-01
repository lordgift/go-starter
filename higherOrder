package main

import "fmt"

func main() {

	fmt.Print("test")

	higherOrder(func(isOk bool) {
		if isOk {
			fmt.Print("99999")
		}
	})
}

func higherOrder(completion func(bool)) {
	completion(true)
}
