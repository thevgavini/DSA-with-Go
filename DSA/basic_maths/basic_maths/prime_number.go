package basic_maths

// import "math"

// func IsPrime(x int) bool {
// 	count:=0
// 	for i := 1; i <= int(math.Sqrt(float64(x)))+1; i++ {
// 		if x%i == 0{
// 			count+=1
// 		}
// 		if count != 2{
// 			return false
// 		}
// 	}

// 	return true

// }


func IsPrime (x int ) bool {
	count := 0
	
	for i:=1; i*i<=x; i++ {
		if x%i == 0 {
			count++
			if i!=x%i {
				count++
			}
		}
		if count>2{break}
	}
	return count == 2
}