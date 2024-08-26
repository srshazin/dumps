package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)
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
