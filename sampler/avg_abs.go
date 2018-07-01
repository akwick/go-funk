package sampler

// AvgAbs calculates the absolute value of the average
// of the values in the given slice
func AvgAbs(in []int) int {
	sum := 0
	for _, i := range in {
		sum += i
	}
	avg := sum / len(in)
	if avg > 0 {
		return avg
	} else {
		return avg * (-1)
	}
}
