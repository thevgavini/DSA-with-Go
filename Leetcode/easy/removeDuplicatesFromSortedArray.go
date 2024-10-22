package leetcode

func removeDuplicates(nums []int) int {
	prev:=0
	for i:=1;i<len(nums);i++{
	 if nums[prev]!=nums[i]{
		 prev++
		 nums[prev]=nums[i]
	 }
	}
	return prev+1
 }