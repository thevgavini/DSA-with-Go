package basic_recursion

import "fmt"


func PrintN(n int) {

    count:=0

    var printRecursive func(int)

    printRecursive = func(n int){
        fmt.Println(n)
        count++
        if count == 3 {
            return
        }

        printRecursive(n)
    }
	
    printRecursive(n)
}
