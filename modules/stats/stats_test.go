package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	tests := []struct {
		in  []float64
		out float64
	}{
		{[]float64{3.11, 4, 0, 10, 9.8}, 0},
		{[]float64{0, 4, 5, 10, 9.8}, 0},
		{[]float64{10, 4, 5, 10, 9.8, 0}, 0},
	}
	for _, tc := range tests {
		assert.Equal(t, Min(tc.in), tc.out)
	}

}

func TestMax(t *testing.T) {
	tests := []struct {
		in  []float64
		out float64
	}{
		{[]float64{3.11, 4, 0, 10, 9.8}, 10},
		{[]float64{10, 4, 5, 0, 9.8}, 10},
		{[]float64{0, 4, 5, 10, 9.8, 10}, 10},
	}
	for _, tc := range tests {
		assert.Equal(t, Max(tc.in), tc.out)
	}

}

func TestAverage(t *testing.T) {
	tests := []struct {
		in  []float64
		out float64
	}{
		{[]float64{9.8}, 9.8},
		{[]float64{1, 10, 9.8, 10}, 7.7},
		{[]float64{3.11, 4, 0, 10, 9.8, 24.10}, 8.501666666666667},
	}
	for _, tc := range tests {
		assert.Equal(t, Average(tc.in), tc.out)
	}

}

func TestReport(t *testing.T) {
	tests := []struct {
		in  []float64
		out Report
	}{
		{[]float64{9.8}, Report{
			Min:     9.8,
			Max:     9.8,
			Average: 9.8,
		}},
		{[]float64{1, 10, 9.8, 10}, Report{
			Min:     1,
			Max:     10,
			Average: 7.7,
		}},
		{[]float64{3.11, 4, 0, 10, 9.8, 24.10}, Report{
			Min:     0,
			Max:     24.10,
			Average: 8.501666666666667,
		}},
	}
	for _, tc := range tests {
		assert.Equal(t, MakeReport(tc.in), tc.out)
	}

}

// func Min(nums []float64) float64
// func Max(nums []float64) float64
// func Average(nums []float64) float64
// type Report struct {
// 	Min float64
// 	Max float64
// 	Average float64
// }
// func MakeReport(nums []float64) Report
