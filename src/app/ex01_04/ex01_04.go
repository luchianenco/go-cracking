package ex01_04

import (
	"fmt"
	"strings"
)

const name = "Ex. 1.4: Replace space char with '%%20' result: %s";

func Execute(str string) {
	res := main(str)
	msg := fmt.Sprintf(name, res)
	fmt.Println(msg)
}

func main(str string) string {
	var newArray []string
	strArray := strings.Split(str, "");

	for i := 0; i < len(strArray); i++ {
		if strArray[i] == " " {
			newArray = append(newArray, "%")
			newArray = append(newArray, "2")
			newArray = append(newArray, "0")
		} else {
			newArray = append(newArray, strArray[i])
		}
	}


	return strings.Join(newArray, "");
}