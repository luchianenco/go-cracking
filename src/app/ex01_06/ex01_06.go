package ex01_06

import (
	"fmt"
)

const name = "Ex. 1.6: Rotate NxN image matrix by 90°:"

func Execute(matrix [][]int32) {
	res := main(matrix)

	// Original matrix show
	msg := fmt.Sprintf(name)
	fmt.Println(msg)
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Result matrix show
	fmt.Println("Result: ")
	for _, row := range res {
		fmt.Println(row)
	}
}

func main(matrix [][]int32) [][]int32 {

	rotated := make([][]int32, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		rotated[i] = make([]int32, len(matrix))
	}

	row := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		fmt.Println(rotated[row])
		current := 0
		for _, el := range matrix[i] {
			rotated[current][row] = el
			current++
		}
		row++
	}

	return rotated
}