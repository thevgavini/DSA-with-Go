package basic_maths

func ReverseNumber(x int) int {
    var rev int 
    for x!=0 {
        rev=rev*10+x%10
        x=x/10
    }

	return rev
}