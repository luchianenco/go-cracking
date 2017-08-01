package ex01_05

import (
	"fmt"
	"strings"
	"strconv"
)

const name = "Ex. 1.5: Compress string %s, result: %s"

var newArray []string

func Execute(str string) {
	res := main(str)
	msg := fmt.Sprintf(name, str, res)
	fmt.Println(msg)
}

func main(str string) string {
	strArray := strings.Split(str, "")
	currentChar := strArray[0];
	charCounter := 1

	for i := 1; i < len(strArray); i++ {
		if currentChar == strArray[i] {
			charCounter++
			continue
		}

		appendNewValues(currentChar, charCounter)
		currentChar = strArray[i]
		charCounter = 1
	}

	appendNewValues(currentChar, charCounter)

	return strings.Join(newArray, "")
}

func appendNewValues(currentChar string, counter int) {
	newArray = append(newArray, currentChar)

	if counter > 1 {
		newArray = append(newArray, strconv.Itoa(counter))
	}
}