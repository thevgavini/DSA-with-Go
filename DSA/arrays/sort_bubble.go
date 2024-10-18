package arrays

/*
Steps: 
	1. Compare each element to the next
	2. Iterate over the array, compare an element with its next, replace the left with the smallest
	3. For an array of size n, repeat this n times 

*/


func BubbleSort (arr []int)[]int {
	for j:=0; j<len(arr)-1;j++{
		for i:=0;i<len(arr)-1;i++{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1] = arr [i+1],arr[i]
			}
		}
	}
	
return arr

}





