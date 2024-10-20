package basic_recursion


func FactorialOfANumber(n int) int {
	if n<1 {
		return 1
	}

	counter = n * FactorialOfANumber(n-1)
	return counter
}