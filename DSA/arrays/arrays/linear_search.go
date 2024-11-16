package arrays 

func LinearSearch(arr []int, key int) int {

	for index, val:= range arr {
		if val==key{
			return index
		}
	}

	return -1
}