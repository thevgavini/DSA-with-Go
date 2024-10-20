package leetcode

func countPrimes(n int) int {

	if n < 2 {
	return 0
}
var eratoSieve func(int) []int
eratoSieve = func(n int)[]int {
arr:= make([]int,n+1)
//initializing arr with 1s
for i:=2; i<=n; i++ {
	arr[i] = 1
}
for k:=2;k*k<=n; k++ {
	if arr[k] == 1{
		for j:=k*k; j<=n; j+=k {
			arr[j] = 0
		}
	}

}
return arr
}
list:=eratoSieve(n)
count:=0
for i:=2; i<n; i++ {
count+=list[i]
} 
return count
}