package main

import (
	"desec/models"
	"fmt"
)

func main() {
	printPersonDetails("David", "Canada", 100, false)

	person := models.Person{}

	person.Name = "Alan"
	person.Address = "Brazil"
	person.Age = 33
	person.Smoker = false

	println(person.Name)
}

func printPersonDetails(name, address string, age int, smoker bool) {
	fmt.Println("Name: ", name)
	fmt.Println("Address: ", address)
	fmt.Println("Age: ", age)
	fmt.Println("Smoker: ", smoker)
}
