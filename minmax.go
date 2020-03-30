package main

// Min creates a min value.
func Min(a []int) int {
	min := a[0]
	for _, value := range a {
		if value <= min {
			min = value
		}
	}
	return min
}

// Max creates a max value
func Max(a []int) (max int) {
	for _, value := range a {
		if value >= max {
			max = value
		}
	}
	return
}
