package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var clients = make(map[net.Conn]bool)
var mu sync.Mutex

func clientHandler(conn net.Conn) {
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

		go clientHandler(conn)
	}
}

func sock_client() {
	conn, error := net.Dial("tcp", "192.168.158.190:8000")
	if error != nil {
		log.Fatal(error)
	}
	defer conn.Close()

	go readMessages(conn)
	// Send user input to the server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		_, err := fmt.Fprintln(conn, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by the server")
			return
		}
		fmt.Printf("<%v>: %v", conn.RemoteAddr(), message)
	}
}

func main() {

	fmt.Printf("Eneter mode:\n1\tfor server\n2\tfor client\n")
	var i int
	fmt.Scan(&i)
	if i == 1 {
		socket_server()
	}
	if i == 2 {
		sock_client()
	}

}
