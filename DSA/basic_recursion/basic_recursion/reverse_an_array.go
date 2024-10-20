package basic_recursion

func ReverseArray(arr []int, start int, end int )  []int {

	if start < end {
		arr[start],arr[end] = arr[end], arr[start]
		ReverseArray(arr, start+1, end-1)
	}
	return arr
}