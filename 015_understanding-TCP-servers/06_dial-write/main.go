package main

import (
	"net"
	"log"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you!")
}
