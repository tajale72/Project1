package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{}

func main() {
	// Serve the HTML file
	http.HandleFunc("/", serveHome)
	// Handle WebSocket connections
	http.HandleFunc("/ws", handleConnections)
	// Start the server on localhost port 8000
	fmt.Println("Server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer ws.Close()
	// Register new client
	clients[ws] = true
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
