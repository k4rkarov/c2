package main

import (
	"log"
	"os"
)

func main() {
	/*//creating files
	os.Create("myFile.txt")

	//creating a file on another folder
	//fmt.Println(os.TempDir())
	//os.Create(os.TempDir() + "\\" + "myFile.txt")

	file, err := os.OpenFile("myFile.txt", os.O_APPEND, 660)

	if err != nil {
		log.Panic(err.Error())
	}
	//it is mandatory to close the file after opening it
	defer file.Close()

	file.Write([]byte("DESEC Security"))
	*/

	// Open the file with proper flags
	file, err := os.OpenFile("myFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)

	if err != nil {
		log.Panic(err.Error())
	}
	// It is mandatory to close the file after opening it
	defer file.Close()

	_, err = file.Write([]byte("DESEC Security\n"))
	if err != nil {
		log.Panic(err.Error())
	}
}
