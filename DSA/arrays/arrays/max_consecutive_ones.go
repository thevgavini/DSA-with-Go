package arrays

func MaxConsecutiveOnes(arr []int) int {
	counter :=0
	max:=0

	for _,val := range arr {

		if val == 1 {
			counter ++ 
			if max<counter{
				max = counter
			}
		} else {
			counter = 0
		}

	}

	return max
}