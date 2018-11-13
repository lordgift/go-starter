package lcdexample_test

import (
	"fmt"
	"lcdexample"
	"testing"
)

func Test_0_To_9(t *testing.T) {
	tcs := []struct {
		n      int
		expect string
	}{
		{0, ` _ 
| |
|_|`},
	}

	//   ...   ._.   ._.   ...   ._.   ._.   ._.   ._.   ._.
	//   ..|   ._|   ._|   |_|   |_.   |_.   ..|   |_|   |_|
	//   ..|   |_.   ._|   ..|   ._|   |_|   ..|   |_|   ..|

	for _, tc := range tcs {
		r := lcdexample.Lcd(tc.n)
		if r != tc.expect {
			t.Errorf("Expected %s but got %s", tc.expect, r)
		}
	}
}

func ExampleLCD() {
	r := lcdexample.Lcd(0)
	fmt.Println(r)
}
