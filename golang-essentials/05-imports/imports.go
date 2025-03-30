package main

import (
	. "fmt"
	msg "log"
)

func main() {
	Println("FMT print")

	//shows a log message and exits application
	//log.Fatal("Application error")

	//shows a log message and continues the execution
	msg.Println("LOG print")

}
