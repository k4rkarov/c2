package main

import "fmt"

func main() {
	ifcase()
	switchcase()
	forcase()
}

func ifcase() {
	age := 19

	if age <= 14 {
		fmt.Println("Just a kid")
	} else if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Underaged")
	}
}

func switchcase() {
	age := 19

	switch age {
	case 14:
		fmt.Println("Just a kid")
	case 18:
		fmt.Println("Adult")
	default:
		fmt.Println("Underaged")
	}
}

func forcase() {
	var namesList []string

	namesList = append(namesList, "Desec")
	namesList = append(namesList, "Golang")
	namesList = append(namesList, "Go")

	for counter := 1; counter <= 10; counter++ {
		fmt.Print(counter)
	}

	for index, name := range namesList {
		fmt.Printf("\nIn the position %d I have the name %s", index, name)
		if name == "Go" {
			fmt.Println("\nI am coding in ", name)
		}
	}
}
