package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Store the client to the clients map
	mu.Lock()
	clients[conn] = true
	mu.Unlock()
	fmt.Printf("New client <%v> Connected! Total clients: %v\n", conn.RemoteAddr(), len(clients))
	reader := bufio.NewReader(conn)
	welcomeMessageByte, _ := json.Marshal(Message{
		Message: "Connection established!\n",
		Sender:  "Admin",
	})
	conn.Write(append(welcomeMessageByte, '\n'))
	for {
		message, error := reader.ReadString('\n')

		if error != nil {
			// Delete the client if there's an error
			fmt.Printf("Client <%v> disconnected. \n", conn.RemoteAddr())

			// then delete that client from the clients map
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			return
		}
		broadCastMessage(message, conn)

	}

}

func broadCastMessage(message string, conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {

		// send the message to the client
		if conn != client {
			jsonByte, _ := json.Marshal(Message{
				Message: message,
				Sender:  conn.RemoteAddr().String(),
			})

			_, error := client.Write(append(jsonByte, '\n'))
			if error != nil {
				fmt.Printf("Error sending message to client: %v\n", error)
			}
		}

	}
}

func socket_server() {
	listener, error := net.Listen("tcp", "0.0.0.0:8000")
	if error != nil {
		log.Fatal(error)
	}

	defer listener.Close()
	fmt.Println("Socket server started on port 8000")

	for {
		conn, error := listener.Accept()
		if error != nil {
			log.Fatal(error)
		}

		go handleConnection(conn)
	}
}
