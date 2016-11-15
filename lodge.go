package main

import (
  "log"
  "net"
)

func main() {
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

func fatalError(err error){
  if err != nil {
    log.Fatal("error: ", err.Error())
  }
}
