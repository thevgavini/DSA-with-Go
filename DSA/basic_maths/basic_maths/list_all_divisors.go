package basic_maths

import "math"

func Divisors(x int) []int {
	var divisors []int

	for i:= 1; i<=int(math.Sqrt(float64(x))); i++ {
		if x%i==0{
			divisors = append(divisors, i)
			if i!=x/i{
				divisors = append(divisors, x/i)
			}
			
		}
	}
	return divisors
}