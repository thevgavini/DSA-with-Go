package leetcode

func majorityElement(nums []int) int {
    hashMap := make(map[int]int)
    for _,val:= range nums {
        hashMap[val]+=1
        if hashMap[val]>len(nums)/2{
            return val
        }
    }

    return 1

}