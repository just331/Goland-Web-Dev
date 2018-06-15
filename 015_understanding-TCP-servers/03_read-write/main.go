package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		go hangle(conn)
	}
}

func hangle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()


	fmt.Println("code got here!") // never get here/
}