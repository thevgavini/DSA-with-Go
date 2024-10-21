package leetcode

func myPow(x float64, n int) float64 {
    ans:= 1.0000
    temp:=1

    if n<0{
        temp=-1
        n=-1*n
    }

    for n!=0 {
        if n%2 == 0 {
            x = x*x
            n=n/2
        } else{
            ans = ans * x
            n = n-1
        } 
    }

    if temp>0{
        return ans
    }else{
        return 1/ans
    }
   
}