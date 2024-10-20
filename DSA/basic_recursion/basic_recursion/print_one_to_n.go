//without passing the counter through function call/declaring an inner function

package basic_recursion

import "fmt"


func PrintOneToN (n int ) {
	if 	counter>n {
		return
	}
	fmt.Println(counter)
	counter++
	PrintOneToN(n)
}