package leetcode

func findMaxConsecutiveOnes(nums []int) int {

    counter :=0
	max:=0

    for _,val := range nums {

        if val == 1 {
            counter++
            if counter > max {
                max = counter
            }
        } else {
            counter = 0
        }

    }
     return max

}