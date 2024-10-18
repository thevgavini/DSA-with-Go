package arrays

//returns the max and min of an array
func MinMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max, min

}
