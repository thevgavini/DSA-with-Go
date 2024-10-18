package arrays

func InsertionSort(arr []int)[]int {

	for i:=0;i<len(arr);i++{
		for j:=i;j>=1;j--{
			if arr[j]<arr[j-1]{
				arr[j],arr[j-1]=arr[j-1],arr[j]
			}
		}
	}

	return arr


}

