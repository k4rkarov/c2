package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal("Error when trying to conect to the server.", err.Error())
	}

	for {
		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		connection.Write([]byte(msg))

		res, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Println("SERVER: ", res)

	}

}
