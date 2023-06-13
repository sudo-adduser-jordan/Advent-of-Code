package mathy

// returns min value in list
func MinInt(nums ...int) int {
	n := nums[0]
	for _, v := range nums {
		if v < n {
			n = v
		}
	}
	return n
}
