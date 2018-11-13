package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tcs := []struct {
		n      int
		expect string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
	}

	for _, tc := range tcs {
		r := FizzBuzz(tc.n)
		if r != tc.expect {
			t.Errorf("Expected %d but got %s", tc.n, tc.expect)
		}
	}
}

func TestFizzBuzz_3_Must_Fizz(t *testing.T) {
	// r := FizzBuzz(3)
	// if r != "Fizz" {
	// 	t.Error("Expected Fizz but got", r)
	// }

	testcases := []int{3, 6, 9, 12}
	for _, tc := range testcases {
		r := FizzBuzz(tc)
		if r != "Fizz" {
			t.Error("Expected Fizz but got", r)
		}
	}

}
