package arrays

func FindMissingNumber (nums []int) int {

	leng := len(nums)
	sum  := leng*(leng+1)/2

	sum_array := 0

	for _,val:= range nums {
		sum_array+=val
	}

	return sum-sum_array




}