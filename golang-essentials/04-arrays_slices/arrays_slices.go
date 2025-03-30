package main

import "fmt"

func main() {
	//This is how slice works
	var namesList []string

	//you canto do it like that: namesList[0] = "DESEC"
	namesList = append(namesList, "DESEC")
	namesList = append(namesList, "Golang")
	namesList = append(namesList, "Go")

	fmt.Println(namesList[2])
}
