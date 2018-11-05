package main

import (
	"fmt"
)

type customer struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	//shorthand declaration
	name := "John"
	fmt.Println("Hello, ", name)

	fmt.Println("----- Branch Statement -----")

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

	fmt.Println("----- Slice Statement -----")

	//Slice (array)
	var nums []int
	fmt.Println(len(nums))

	nums = append(nums, 1)
	nums = append(nums, 2)

	fmt.Println(nums)
	fmt.Println(len(nums))

	//panic: runtime error: index out of range
	// fmt.Println(nums[0], nums[2])

	//Slice (array) also have shorthand declaration & initial elements
	nums2 := []int{1, 2, 3, 4, 5}

	fmt.Println(nums2[:5])
	fmt.Println(nums2[1:2])
	fmt.Println(nums2[3:])

	//allocate 10 elements for Slice (array)
	nums3 := make([]int, 10)
	fmt.Println(nums3)

	fmt.Println("----- Map Statement -----")
	scores := map[string]int{
		"John": 100,
	}
	fmt.Println(scores)

	fmt.Println("----- struct -----")
	var cust customer
	fmt.Printf("%#v\n", cust)
}
