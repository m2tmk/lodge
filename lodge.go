package main

import (
  "fmt"
	"log"
	"net"
  zmq "github.com/zeromq/goczmq"
)

func main() {
  tcpServer()
  //udpServer()
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

func udpServer() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:54321")

	fatalError(err)

	conn, err := net.ListenUDP("udp", addr)

	fatalError(err)

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		rlen, addr, err := conn.ReadFromUDP(buf)

		fatalError(err)

		s := string(buf[:rlen])

		log.Printf("Received: [%v]: %v\n", addr, s)
	}
}

func fatalError(err error) {
	if err != nil {
		log.Fatal("error: ", err.Error())
	}
}
