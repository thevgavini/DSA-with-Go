package basic_maths

func PrimeFactors(n int) []int {

	var arr []int

	for i:=2;i*i<=n;i++{
		if n%i==0{
			arr =append(arr,i)

			for n%i==0 {
				n=n/i
			}	
		}
	}

	arr = append(arr, n)

	return arr
}