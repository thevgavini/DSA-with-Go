package basic_maths

import "fmt"

func Fib(n int)  {
	result := 0
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		result = a + b
		fmt.Println(a)
		a=b
		b=result
	}

}