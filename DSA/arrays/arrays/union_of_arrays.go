package arrays 

func UnionOfArrays(arr1 []int, arr2 []int) []int {

	hashMap :=make(map[int]bool)
	var arr []int


	for _,val:= range arr1 {
		if !hashMap[val]{
			hashMap[val] = true
			arr = append(arr, val)
		}
	}
	for _,val:= range arr2{
		if !hashMap[val]{
			hashMap[val] = true
			arr = append(arr, val)
		}
	}


	return arr
	




}