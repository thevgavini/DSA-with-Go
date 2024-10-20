package basic_maths

func PrimeCount(n int)int{

	arr := EratoSieve(n)
	count:=0
	for i:=2;i<=n;i++{
		count+=arr[i]
		arr[i]=count
	}

	return arr[n]
}