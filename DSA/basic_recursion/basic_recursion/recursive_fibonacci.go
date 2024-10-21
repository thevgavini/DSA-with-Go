package basic_recursion

func RecursiveFibonacci (n int) int {
	var recurFib func(int) int

	recurFib = func(n int) int {
		

			if n<=1 {
				return n
			}

			

		
		
		return recurFib(n-1) + recurFib(n-2)
	

		
	}

	return recurFib(n)
}