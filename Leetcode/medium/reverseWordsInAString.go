package leetcode

import "strings"

func reverseWords(s string) string {
	stringSplit:=strings.Fields(s)
	stringReturn := ""
	for i:=len(stringSplit)-1; i>=0; i--{
        if i==0{
            stringReturn+=stringSplit[i]
        }else{stringReturn+=stringSplit[i]+" "}

	}
	return stringReturn
}