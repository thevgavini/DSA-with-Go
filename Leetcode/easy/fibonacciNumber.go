package leetcode

func fib(n int) int {
    result := 0
    a := 0
    b := 1

    for i:=0;i<n;i++{
        result = a + b
        a = b
        b = result
    }

    return a
}