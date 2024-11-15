package arrays

func LeftRotateByOneBF(arr []int) []int {

	var tempArray []int

	for i:=1; i<len(arr)-1; i++ {
		tempArray=append(tempArray, arr[i])
	}
	tempArray = append(tempArray, arr[0])
	return tempArray

}