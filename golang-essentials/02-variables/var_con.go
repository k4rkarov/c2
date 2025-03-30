package main

import "fmt"

const myConst string = "const in GLOBAL SCOPE"

var myVariable string

func main() {
	myVariable = "variable in func MAIN"

	fmt.Println(myVariable)
	fmt.Println(myConst)

	print()
}

func print() {
	myVariable = "variable in func PRINT"
	fmt.Println(myVariable)
	fmt.Println(myConst)
}
