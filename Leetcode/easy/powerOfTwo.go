package leetcode

import "math"

func isPowerOfTwo(n int) bool {
    p:= math.Log2(float64(n))
    if (math.Ceil(p)-p) != 0 {
        return false
    }
    return true
}