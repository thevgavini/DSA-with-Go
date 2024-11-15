package arrays

func Rotate(nums []int, k int)  []int{

	//Normalizing k for cases where k>len(nums) and k<0
    k = k % len(nums)
    if k < 0 {
        k += len(nums)
    }


    arb:=len(nums)-k
    var tempArray []int

    for j:=arb; j<len(nums); j++{
        tempArray = append(tempArray,nums[j])
    }

    for i:=0; i<arb; i++ {
        tempArray = append(tempArray, nums[i])
    }

    for k=0; k<len(tempArray); k++ {
        nums[k] = tempArray[k]
    }

	return nums
    
}