package main

import (
	"commons/commons"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"log"
	"net"
	"os"
	"time"
)

var (
	message     commons.Message
	waitingTime = 10
)

const (
	SERVER = "127.0.0.1"
	PORT   = "9090"
)

func init() {
	message.AgentHostname, _ = os.Hostname()
	message.AgentCWD, _ = os.Getwd()
	message.AgentId = idGenerator()
}

func main() {
	log.Println("The agent has started!")

	for {
		channel := connectServer()

		gob.NewEncoder(channel).Encode(message)
		gob.NewDecoder(channel).Decode(&message)

		channel.Close() // ✅ Close the connection immediately
		log.Println("Sleeping for", waitingTime, "seconds...")
		time.Sleep(time.Duration(waitingTime) * time.Second)
	}
}

func connectServer() net.Conn {
	for {
		channel, err := net.Dial("tcp", SERVER+":"+PORT)
		if err != nil {
			log.Println("Error connecting to server:", err.Error()) // Log instead of fatal
			time.Sleep(5 * time.Second)                             // Wait before retrying
			continue
		}
		return channel
	}
}

func idGenerator() string {
	myTime := time.Now().String()
	hasher := md5.New()
	hasher.Write([]byte(message.AgentHostname + myTime))
	return hex.EncodeToString(hasher.Sum(nil))
}
