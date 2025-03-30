package main

import "fmt"

func main() {
	var myVar string
	var myVar1 bool
	var myVar2 int
	var namesList [3]string
	myVar3 := "variable type defined automatically"

	namesList[0] = "DESEC"
	namesList[1] = "Golang"
	namesList[2] = "GO"

	myVar = "word or phrase"
	myVar1 = true
	myVar2 = 1

	fmt.Println(myVar)
	fmt.Println(myVar1)
	fmt.Println(myVar2)
	fmt.Println(myVar3)
	fmt.Println(namesList[1])
}
