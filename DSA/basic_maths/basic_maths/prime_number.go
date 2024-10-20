package basic_maths


func IsPrime (x int ) bool {
	count := 0
	for i:=1; i*i<=x; i++ {
		if x%i == 0 {
			count++
			if i!=x/i {
				count++
			}
		}
		if count>2{break}
	}
	return count == 2
}