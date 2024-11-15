package arrays

func RotateRightOptimal(nums []int, k int) []int{

	var reverseArray func([]int) = func(arr []int) {
		for i:=0; i<len(arr)/2; i++{
			arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
		}
	}

	reverseArray(nums[len(nums)-k:])
	reverseArray(nums[:len(nums)-k])
	reverseArray(nums)

	return nums


}