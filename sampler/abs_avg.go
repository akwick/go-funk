package sampler

// AbsAvg calculates the average of the absolute values in the given slice
func AbsAvg(in []int) int {
	sum := 0
	for _, i := range in {
		if i < 0 {
			pos := i * (-1)
			sum += pos
		} else {
			sum += i
		}
	}
	return sum / len(in)
}
