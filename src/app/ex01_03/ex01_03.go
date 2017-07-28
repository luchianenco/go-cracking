package ex01_03

import (
	"fmt"
	"strings"
	"sort"
)

const name = "Ex. 1.3: The string \"%s\" %s permutation of string \"%s\"";

func Execute(original string, permutation string) {

	var result string;

	res := main(original, permutation)

	if res == true {
		result = "is"
	} else {
		result = "isn't"
	}

	msg := fmt.Sprintf(name, permutation, result, original)
	fmt.Println(msg)
}

func main(original string, permutation string) bool {

	if (len(original) != len(permutation)) {
		return false;
	}

	originalArray := strings.Split(original, "");
	permutationArray := strings.Split(permutation, "");
	sort.Strings(permutationArray)
	sort.Strings(originalArray)

	if (strings.Join(permutationArray, "") != strings.Join(originalArray, "")) {
		return false
	}

	return true;
}