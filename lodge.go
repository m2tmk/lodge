package main

import (
  "fmt"
	"log"
  zmq "github.com/zeromq/goczmq"
)

func main() {
  tcpServer()
}

func tcpServer() {
  rep, _ := zmq.NewRep("tcp://*:54321")
  defer rep.Destroy()

  log.Println("Rep created and bound.")

  for {
    message, _ := rep.RecvMessage()
    log.Printf("message: %v", string(message[0]))

    reply := fmt.Sprintf("ok: %v", string(message[0]))

    rep.SendFrame([]byte(reply), 0)
  }
}

func fatalError(err error) {
	if err != nil {
		log.Fatal("error: ", err.Error())
	}
}
