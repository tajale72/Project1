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

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	Chatgpt()
// }

// func Chatgpt() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("you can ask here: ")
// 	sentence, _ := reader.ReadString('\n')
// 	fmt.Println("You entered:", sentence)

// 	splitting_sentence := strings.Split(sentence, " ")
// x
// 	for k, _ := range splitting_sentence {
// 		if splitting_sentence[k] == "what" {
// 			fmt.Println("what:", splitting_sentence[k+2])
// 			if splitting_sentence[k+1] == "is" {
// 				fmt.Println("is:", splitting_sentence[k+1])
// 				if splitting_sentence[k+2] == "apple" {
// 					fmt.Println("An apple is a fruit, which is very sweet and has a lot of health benefits")
// 				}
// 				if splitting_sentence[k+2] == "banana" {
// 					fmt.Println("An banana is a fruit, which is very high is carb and has potassium")
// 				}
// 			}

// 		}
// 	}
// 	return ""
// }
