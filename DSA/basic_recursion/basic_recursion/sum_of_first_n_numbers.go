package basic_recursion

// import "fmt"

func SumOfFirstN (n int) int {

	if n<1 {
		return counter
	}
	counter+=n
	n--
	return SumOfFirstN(n)
}