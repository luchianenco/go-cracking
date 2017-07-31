package main

import (
	"app/ex01_01"
	"app/ex01_03"
	"app/ex01_04"
	"fmt"
)

func main() {
	fmt.Println("Started...")
	ex01_01.Execute("12343")
	ex01_03.Execute("12343", "12342")
	ex01_04.Execute("Here we go")
	fmt.Println("Finished...")
}
