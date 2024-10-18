package leetcode

func selfDividingNumbers(left int, right int) []int {
    var test []int
    for i := left; i <= right; i++ {
        sdn := true
        temp := i
        for temp != 0 {
            if temp%10 == 0 || i%(temp%10) != 0 {
                sdn = false
                break
            }
            temp = temp / 10
        }
        if sdn {
            test = append(test, i)
        }
    }
    return test
}

