package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8180")
	if err != nil {
		log.Fatalf("could not dial server: %v\n", err)
	}
	conn.Write([]byte("Assalomu alaykum!!!\n"))
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(msg)
	conn.Close()
}
