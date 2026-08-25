package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8180")
	if err != nil {
		log.Fatalf("could not listen 8180: %v", err)
	}
	fmt.Println("TCP server ishga tushdi 8180 portda!!!")
	// Local IP address: 127.0.0.1 -> localhost
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic("could accept")
		}
		go func(c net.Conn) {
			defer c.Close()
			msg, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Println("Client: ", msg)

			c.Write([]byte("Salom Client\n"))

		}(conn)
	}
}
