package basic_maths

func IsPalindrome(x int) bool {

	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}
	var rev int
	temp := x
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	return temp == rev
}
