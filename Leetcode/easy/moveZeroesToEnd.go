package leetcode

func moveZeroes(nums []int)  {
    prev:=0
    current:=1
    for current<len(nums) {
        if nums[prev]==0{
            if nums[current]!=0{
                nums[prev],nums[current]=nums[current],nums[prev]
                prev++
            }else{
                current++
            }
        }else{
            prev++
            current++
        }

        }
    }
