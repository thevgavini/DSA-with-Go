package basic_recursion

//the following function can be used when passing only string as the argument. However, it requires a global variable to be set, and doesn't reset the global variable later


// func IsPalindrome(str string) bool {
// 	if str[counter]!=(str[len(str)-1-counter]){return false}
// 	if counter == len(str)/2 {return true}
// 	counter ++
// 	return IsPalindrome(str)

// }

func IsPalindrome (str string, i int) bool {

	if str[i]!=str[len(str)-i-1] {
		return false
	}

	if i == (len(str)/2){
		return true
	}

	return IsPalindrome(str,i+1)
}