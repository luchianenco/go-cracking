package ex01_07

import (
	"fmt"
)

const name = "Ex. 1.7: Make column and row 0 if an element in matrix NxM is equal to 0:"

func Execute(matrix [][]int) {
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

func main(matrix [][]int) [][]int {

	processed := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		processed[i] = make([]int, len(matrix[0]))
		copy(processed[i], matrix[i])
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				makeColumnAndRowZero(processed, row, col)
			}
		}
	}

	return processed
}

func makeColumnAndRowZero(processed [][]int, row int, col int) {
	for i := 0; i < len(processed[0]); i++ {
		processed[row][i] = 0
	}

	for i := 0; i < len(processed); i++ {
		processed[i][col] = 0
	}
}