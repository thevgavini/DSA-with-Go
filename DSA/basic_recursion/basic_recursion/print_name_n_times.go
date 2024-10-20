package basic_recursion
import "fmt"

func PrintNTimes(i int, n int){

	if (i>n){
		return
	}
	fmt.Println("name")
	PrintNTimes(i+1, n)

}
