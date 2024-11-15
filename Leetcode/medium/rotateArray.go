package leetcode

func rotate(nums []int, k int)  {

    k=k%len(nums)
    if k<0 {
        k+=len(nums)
    }
   var reverseArray func([]int) 
	reverseArray = func(arr []int) {
		for i:=0; i<len(arr)/2; i++{
			arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
		}
	}
	reverseArray(nums[len(nums)-k:])
	reverseArray(nums[:len(nums)-k])
	reverseArray(nums)
}