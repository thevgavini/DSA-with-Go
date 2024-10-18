package basic_maths

func Gcd(x int, y int ) int {
	for i:=min(x,y); i>0 ; i-- {
		if x%i==0 && y%i == 0 {
			return i
		} 
	}
	return 1
}


func GcdEuclidean (x int, y int ) int {

	for (x>0 && y>0) {

		if (x>y){
			x=x%y
		}else {
			y=y%x
		}

	}

	if x>0 {
		return x
	}

	return y
}