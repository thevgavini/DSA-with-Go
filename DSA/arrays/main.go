package main

import (

	"fmt"
	"DSA/arrays/arrays"

)

//MinMax

func main() {
	// testArray := []int{9,8,4,-2,0,1,2,3,5,4}
	// testArraySorted := []int{-2,0,1,2,3,4}
	// testArraySortedDuplicates:=[]int{1,1,2,2,2,3,3}
	testArrayInOrder := []int{1,2,3,4,5,6,7}


// 	min, max := arrays.MinMax(testArray)
// 	fmt.Printf("Min: %d, Max: %d\n", min, max)
// // }

//Reverse an array
// 	fmt.Println(arrays.ReverseArray(testArray))

//Bubble sort
// fmt.Println(arrays.BubbleSort(testArray))

//Selection sort
// fmt.Println(arrays.SelectionSort(testArray))

//Insertion sort
// fmt.Println(arrays.InsertionSort(testArray))

// fmt.Println(arrays.PrintLcis(testArray))

// fmt.Println(arrays.SecondLargestOptimal(testArray))
// fmt.Println(arrays.CheckIfSorted(testArraySorted))
// fmt.Println(arrays.RemoveDuplicates(testArraySortedDuplicates))

// fmt.Println(arrays.LeftRotateByOneBF(testArray))

// fmt.Println(arrays.Rotate(testArrayInOrder, 3))
fmt.Println(arrays.RotateRightOptimal(testArrayInOrder,3))

}




