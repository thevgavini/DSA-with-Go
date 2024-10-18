package leetcode

import (
    "math"
)

func checkPerfectNumber(num int) bool {
    sum := 1
    if num == 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num%i == 0 {
            sum += i
            if i != num/i {
                sum += num / i
            }
        }
    }
    return sum == num
}
