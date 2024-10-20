package leetcode

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	var eratoSieve func(int) []bool
	eratoSieve = func(n int) []bool {
		isPrime := make([]bool, n+1)
		// initializing arr with 1s
		for i := 2; i <= n; i++ {
			isPrime[i] = true
		}

		for k := 2; k*k <= n; k++ {
			if isPrime[k] == true {
				for j := k * k; j <= n; j += k {
					isPrime[j] = false
				}
			}
		}
		return isPrime
	}

	list := eratoSieve(n)
	count := 0
	for i := 2; i < n; i++ {
		if list[i] {
			count++
		}
	} 
	return count
}
