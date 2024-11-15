package arrays

func LeftRotateByOne (arr []int) []int {
	temp:=arr[0]
	for i:=0; i<len(arr)-2; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = temp
	return arr
}
