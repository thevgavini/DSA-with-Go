package leetcode

func isUgly(n int) bool {
    
    var arr []int

    if n < 1 {
        return false
    }

    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            arr = append(arr, i)
            for n%i == 0 {
                n = n / i
            }

            if i != 2 && i != 3 && i != 5 {
                return false
            }
        }
    }

    if n > 1 {
        if n != 2 && n != 3 && n != 5 {
            return false
        }
    }

    return true
}
