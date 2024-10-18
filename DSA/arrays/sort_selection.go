package arrays
import "fmt"

/*
Steps: 
	1. Select the first element, find the smallest element, and swap 
	2. Increment the selector, and repeat the process with the remaining slice
*/

func SelectionSort(arr []int)[]int {


	for i:=0; i< len(arr); i++{
		for j:=i; j<len(arr); j++{
			if arr[j]<arr[i]{
				arr[j],arr[i]=arr[i],arr[j]
			}
	
		}
		fmt.Println(arr)
	}
	return arr
	
}

