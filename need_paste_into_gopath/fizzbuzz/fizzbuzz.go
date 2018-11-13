package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(n int) string {

	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.FormatInt(int64(n), 10)
	}
}

// func main() {
// 	for no := 1; no <= 15; no++ {
// 		fmt.Printf("%s,", FizzBuzz(no))
// 	}
// }
