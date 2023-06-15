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

// returns max value in list
func MaxInt(nums ...int) int {
	n := nums[0]
	for _, v := range nums {
		if v > n {
			n = v
		}
	}
	return n
}

// returns max value in slice
func MaxIntSlice(slice []int) int {
	n := slice[0]
	for _, v := range slice {
		if v > n {
			n = v
		}
	}
	return n
}
