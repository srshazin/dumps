package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)


func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("New client <%v> Connected!\n", conn.RemoteAddr())
	// Store the client to the clients map
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	reader := bufio.NewReader(conn)

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
			_, error := fmt.Fprintln(client, strings.Trim(message, "\n"))
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
