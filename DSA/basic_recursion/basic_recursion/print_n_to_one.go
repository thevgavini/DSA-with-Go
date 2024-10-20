package basic_recursion

import "fmt"

func PrintNToOne(n int ){

	if n<1 {
		return
	}
	fmt.Println(n)
	n--
	PrintNToOne(n)
}