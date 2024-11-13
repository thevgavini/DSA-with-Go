package arrays

func SecondLargest(arr []int) (int,int,int,int) {
	
	min:=arr[0]
	max:=arr[0]



	for _,val:= range arr {
		if val<min{
			min = val
		}
		if val>max{
			max = val
		}
	}
	second_min := int(^uint(0)>>1)
	second_max := -second_min

	for _,val := range arr {
		if val<second_min && val!=min{
			second_min = val
		}

		if val>second_max && val!=max {
			second_max = val
		}


	}

	return max,second_max,min,second_min
}