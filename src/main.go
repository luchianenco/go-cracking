package main

import (
	"app/ex01_01"
	"app/ex01_03"
	"app/ex01_04"
	"app/ex01_05"
	"app/ex01_06"
	"app/ex01_07"
	"fmt"
)

func main() {

	matrix1 := [][]int32{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	matrix2 := [][]int{{1, 2, 3, 4}, {5, 0, 7, 0}, {9, 10, 11, 12}}

	fmt.Println("Started...")
	ex01_01.Execute("12343")
	ex01_03.Execute("12343", "12342")
	ex01_04.Execute("Here we go")
	ex01_05.Execute("aabbbbcccccfedddd")
	ex01_06.Execute(matrix1)
	ex01_07.Execute(matrix2)
	fmt.Println("Finished...")
}
