package leetcode


func IsHappy(n int) bool {
    if n==1{
        return true
    }

    hashMap:=make(map[int]bool)

    var squareCount func(int) bool

    squareCount = func (n int) bool{
		hashMap[n] = true
        if n==1 { return true}
        sum:=0
        for n!=0 {
            digit:=n%10
            sum+=digit*digit
            n=n/10
        }
		if hashMap[sum] {return false}
        return squareCount(sum)

    }
    return squareCount(n)

}