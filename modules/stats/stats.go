package stats

type Report struct {
	Min, Max, Average float64
}

func Min(nums []float64) float64 {
	if len(nums) == 0 {
		panic("slice empty")
	}

	m := nums[0]
	for _, v := range nums[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func Max(nums []float64) float64 {
	if len(nums) == 0 {
		panic("slice empty")
	}

	m := nums[0]
	for _, v := range nums[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func Average(nums []float64) float64 {
	if len(nums) == 0 {
		panic("slice empty")
	}

	var s float64

	for _, v := range nums {
		s += v
	}
	return s / float64(len(nums))
}

func MakeReport(nums []float64) Report {
	if len(nums) == 0 {
		panic("slice empty")
	}

	return Report{
		Min:     Min(nums),
		Max:     Max(nums),
		Average: Average(nums),
	}
}
