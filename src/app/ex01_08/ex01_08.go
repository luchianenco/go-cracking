package ex01_08

import (
	"fmt"
	"strconv"
	"strings"
)

const name = "Ex. 1.8: Check if a \"%s\" is a rotation of \"%s\": %s"

func Execute(original string, substring string) {
	res := main(original, substring)

	resStr := strconv.FormatBool(res)
	msg := fmt.Sprintf(name, substring, original, resStr)
	fmt.Println(msg)

}

func main(original string, substring string) bool {

	originalArray := strings.Split(original, "")
	subArray := strings.Split(substring, "")

	for _, char := range subArray {
		i := contains(originalArray, char)
		if i > -1 {
			originalArray[i] = ""
		} else {
			return false
		}
	}

	return true
}

func contains(string []string, value string) int {
	for i, char := range string {
		if char == value {
			return i
		}
	}

	return -1
}