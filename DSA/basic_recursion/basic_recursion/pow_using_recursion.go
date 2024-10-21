package basic_recursion

func PowRecursion(x float64, n int) float64 {
	
	ans := 1.0
	temp:=n
	if n<0{
		n = -1*n
	}

	var recurPow func(float64,int) float64
	recurPow = func(x float64, n int) float64 {
		if n%2==0 {
			x = x*x
			n = n/2
		}else {
			ans = ans * x
			n = n-1
		}

		if n == 0 {
			if temp>0{
				return ans
			}else{
				return 1/ans
			}
		}

		return recurPow(x,n)
	}

	return recurPow(x,n)
   
}