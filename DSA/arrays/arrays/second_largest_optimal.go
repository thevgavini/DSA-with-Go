package arrays

func SecondLargestOptimal (arr []int) (int, int, int, int) {

	max:= -int(^uint(0)>>1)
	second_max:=-int(^uint(0)>>1)

	least:= -max
	second_least := -max

	for _,val:= range arr {
		
		if val>max {
			max = val
		}

		if val>second_max && val!=max {
			second_max = val
		}

		if val<least {
			least = val
		}

		if val<second_least && val!=least{
			second_least = val
		}
	}

	return max,second_max,least,second_least
}