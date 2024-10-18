package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {


	stringSplit:=strings.Fields(s)
    leng := len(stringSplit)-1
	stringReturn := stringSplit[leng]


	for i:=leng-1; i>=0; i--{

        stringReturn+=" "+stringSplit[i]

	}
  	
	return stringReturn
}

func main() {

	fmt.Println(reverseWords("hello sodhara"))

}
