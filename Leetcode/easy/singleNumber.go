package leetcode

func singleNumber(nums []int) int {
    res:=0
    for _,val:= range nums{
        res=res^val
    }
    return res
}