package leetcode


func twoSum(nums []int, target int) []int {
    numAndIndexMap := make(map[int]int)
    for index, num := range nums {
        diff:= target-num
        if idx, exists:= numAndIndexMap[diff]; exists {
            return []int{idx,index}
        }
        numAndIndexMap[num]=index
    } 
    return nil
} 