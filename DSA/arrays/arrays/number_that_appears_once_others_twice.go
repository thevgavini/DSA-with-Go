package arrays

func NumberThatAppearsOnce (arr []int) int {
	res := 0
	for _,val := range arr {
		res = res ^ val
	}

	return res
}