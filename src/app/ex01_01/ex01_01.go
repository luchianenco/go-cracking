package ex01_01

import (
	"fmt"
	"strconv"
)

const name = "Ex. 1.1: All Unique chars in String: ";

func Execute(str string) {

	res := main(str)
	resultStr := strconv.FormatBool(res)
	fmt.Println(name + resultStr)
}

func main(str string) bool {

	var chars [256]bool

	for _, char := range str {
		if (chars[char] == true) {
			return false;
		} else {
			chars[char] = true
		}
	}

	return true;
}