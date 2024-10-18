package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

func main() {
	var finalArray []int
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	stringsNum := strings.Split(input, ",")

	for _, s := range stringsNum{
		num,_ := strconv.Atoi(s)
		finalArray = append(finalArray, num)
	}
	fmt.Println(finalArray)


}
