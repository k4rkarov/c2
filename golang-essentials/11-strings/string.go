package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//shows the path where the file I want to read is
	myFile := filepath.FromSlash(filepath.Join("..", "10-files", "myFile.txt"))
	fmt.Println(myFile)

	//Opens that file and reads its content
	content, err := os.ReadFile(myFile)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(string(content))

	//splitting the string within myFile.txt
	array_content := strings.Split(string(content), " ")
	fmt.Println(array_content[0])

	//replacing the space for an @
	result := strings.ReplaceAll(string(content), " ", "@")
	fmt.Println(result)

	//asking for user input and printing it
	var res string
	fmt.Print("Type something: ")
	fmt.Scanln(&res)
	fmt.Println(res)

	fmt.Print("Type something: ")
	reader := bufio.NewReader(os.Stdin)
	res, _ = reader.ReadString('\n')
	fmt.Println(res)

	//counting words
	words := strings.Split(res, " ")
	fmt.Println("Number of words: ", len(words))
}
