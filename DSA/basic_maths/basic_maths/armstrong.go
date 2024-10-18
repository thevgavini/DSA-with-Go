package basic_maths

import "math"

func IsArmstrong(x int) bool {
    len := int(math.Floor(math.Log10(float64(x))) + 1)
    temp := x
    res := 0
    for x != 0 {
        res += int(math.Pow(float64(x % 10), float64(len)))
        x = x / 10
    }
    return temp == res
}