package arrays

// ReverseArray reverses the input array in place
func ReverseArray(inputArray []int) []int {
	left, right := 0, len(inputArray)-1
	for left < right {
		inputArray[left], inputArray[right] = inputArray[right], inputArray[left]
		left++
		right--
	}
	return inputArray
}

// func main() {
// 	testArray := []int{1, 2, 3, 4, 5}
// 	fmt.Println(arrays.reverseArray(testArray))
// }



