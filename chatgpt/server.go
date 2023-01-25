package main

import (
	"bufio"
	"fmt"
	"net"
)

var clients = make(map[net.Conn]bool)
var broadcast = make(chan string)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	go handleMessages()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		clients[conn] = true

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}

		broadcast <- message
	}
}

func handleMessages() {
	for {
		message := <-broadcast

		for client := range clients {
			_, err := client.Write([]byte(message))
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
