package stats

import "sort"

func Min(nums []float64) float64 {
	sort.Float64s(nums)
	return nums[0]
}
func Max(nums []float64) float64 {
	sort.Float64s(nums)
	return nums[len(nums)-1]
}
func Average(nums []float64) float64 {
	sum := 0.0
	for num := range nums {
		num += sum
	}
	return sum / float64(len(nums))
}

type Report struct {
	Min     float64
	Max     float64
	Average float64
}

func MakeReport(nums []float64) {

}
