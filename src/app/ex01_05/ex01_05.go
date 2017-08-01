package ex01_05

import (
	"fmt"
	"strings"
	"strconv"
)

const name = "Ex. 1.5: Compress string %s, result: %s";

func Execute(str string) {
	res := main(str)
	msg := fmt.Sprintf(name, str, res)
	fmt.Println(msg)
}

func main(str string) string {
	var newArray []string
	strArray := strings.Split(str, "");
	currentChar := strArray[0];
	charCounter := 1

	for i := 1; i < len(strArray); i++ {
		if currentChar == strArray[i] {
			charCounter++
			continue
		}

		newArray = append(newArray, currentChar)

		if charCounter > 1 {
			newArray = append(newArray, strconv.Itoa(charCounter))
		}

		currentChar = strArray[i]
		charCounter = 1;
	}

	newArray = append(newArray, currentChar)

	if charCounter > 1 {
		newArray = append(newArray, strconv.Itoa(charCounter))
	}


	return strings.Join(newArray, "");
}