package leetcode

func findLengthOfLCIS(nums []int) int {
    count:=1
    max:=1
    for i:=1;i<len(nums);i++{
        if nums[i-1]<nums[i]{
            count++
        }else{
            if count>max{
                max = count
            }
            count=1
        }
    }
    if count>max{
        max = count
    }

    return max
    
}