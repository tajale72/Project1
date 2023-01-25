package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go handleMessages(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, message)
	}
}

func handleMessages(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Client 2:", string(message))
	}
}
