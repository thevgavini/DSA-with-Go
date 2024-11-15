package arrays 

// # Algorithm:
// 	1. Initialise two pointers, prev and current
// 	2. Start prev at 0, and current at 1
// 	3. For any value in the array where arr[prev]!=arr[current]:
// 		1. Increment prev
// 		2. Swap values at prev and current

func RemoveDuplicates (arr []int) []int {

	prev:=0

	for current:=1; current<len(arr)-1; current ++ {
		if arr[prev]!=arr[current]{
			prev++
			arr[prev]=arr[current]
		}
	}
	return arr[:prev+1]
}