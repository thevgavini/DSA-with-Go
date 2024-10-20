/*
Eratosthenes' sieve:
	Step 1: Initialize an empty slice, and append 1 until n+1 index
	Step 2: Iterate over the slice, and check if the value is 1
	Step 3: If the value is 1, update all of its multiples with 0

	Optimisations:
		1:
			When updating multiples to 0, start with the square of the number,
			and not the second multiple, as all the multiples before the number would've been covered:
				no need to start from 7*2, as 2*7 has already been marked during 2's pass
		2:
			When iterating over the entire array, iterate only till the sqrt of the length of the array,
			as the squared value is being visited by the pass in the internal loop

*/

package basic_maths

func EratoSieve(n int) []int {

	arr:= make([]int,n+1)

	//Initialize all values to 1
	for i := 2; i <= n; i++ {
		arr[i] = 1
	}
	//Traverse the array
	for k := 2; k*k <= n; k++ {
		//If the number is prime
		if arr[k] == 1 {
			//Mark its multiples as composite
			for l := k * k; l <=n; l += k {
				arr[l] = 0
			}
		}
	}

	return arr
}
