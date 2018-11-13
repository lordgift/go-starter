package stats

import "testing"

func Test_Min(t *testing.T) {

	data := []float64{1, 2, 3, 4, 5}

	r := Min(data)
	if r != 1 {
		t.Error("Expected 1 but got", r)
	}
}

func Test_Max(t *testing.T) {

	data := []float64{1, 2, 3, 4, 5}

	r := Max(data)
	if r != 5 {
		t.Error("Expected 5 but got", r)
	}
}

func Test_Average(t *testing.T) {

	data := []float64{1, 2, 3, 4, 5}

	r := Average(data)
	if r != 3 {
		t.Error("Expected 3 but got", r)
	}
}
