package arrays

func PrintLcis(nums []int)[]int {
	count:=1
	max:=1
	index:=0

	for i:=1;i<len(nums);i++{

		if nums[i-1]<nums[i]{
			count++
		}else{
			if count>max{
				max=count
				index=i-1
			}
			count=1
		}

	}

	return nums[index-max+1:index+1]
}