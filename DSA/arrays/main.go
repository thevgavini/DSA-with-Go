package main

import (

	"fmt"
	"DSA/arrays/arrays"

)

//MinMax

func main() {
	testArray := []int{9,8,4,-2,0,1,2,3,5,4}


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

fmt.Println(arrays.SecondLargestOptimal(testArray))

}




