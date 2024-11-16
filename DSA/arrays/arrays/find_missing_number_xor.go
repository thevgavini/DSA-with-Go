package arrays 

func FindMissingNumberXOR(arr []int) int{

	xor1:=0
	xor2:=0
	for i:=0; i<len(arr); i++ {
		xor1 = xor1^i
		xor2 = xor2^arr[i]
	}
	xor1 = xor1^len(arr)

	return xor1^xor2
}