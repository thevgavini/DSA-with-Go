package basic_maths

import "math"

func IsPrime(x int) bool {
	count:=1
	for i := 1; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0{
			count+=1
		}
		if count>2{
			return false
		}
	}

	return true

}