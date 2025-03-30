package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal("Error when configuring the listener.", err.Error())
	}

	log.Println("Server is listening on port 9090...")

	connection, err := listener.Accept()
	if err != nil {
		log.Fatal("Error accepting connection.", err.Error())
	}

	for {
		msg, _ := bufio.NewReader(connection).ReadString('\n')
		connection.Write([]byte(msg))
	}

}
