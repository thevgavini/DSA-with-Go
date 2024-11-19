package arrays

func MaxLengthSumBF (arr []int, ReqSum int) int {
	sum:=0
	maxLen:=0

	for i:=0; i<len(arr); i++ {
		sum = 0
		for j:=i; j<len(arr); j++{
			sum+=arr[j]
			if sum==ReqSum{
				if maxLen<j-i+1{
					maxLen = (j-i+1)
				}
			} 
		}
	}

	return maxLen
}