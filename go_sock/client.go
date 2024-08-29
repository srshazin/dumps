package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func sock_client() {
	conn, error := net.Dial("tcp", "192.168.10.110:8000")
	if error != nil {
		log.Fatal(error)
	}
	defer conn.Close()

	go readMessages(conn)
	reader := bufio.NewReader(os.Stdin)
	// take user's input
	for {
		// Print the prompt

		// Read the input from the user
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		message = message[:len(message)-1]
		// Send the message to the server
		_, err = fmt.Fprintln(conn, message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
		fmt.Printf("\033[32m<%v>:\033[0m ", conn.LocalAddr().String())
	}

}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		byteMessage, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Connection closed by the server")
			os.Exit(1)
		}

		var message Message
		error := json.Unmarshal(byteMessage, &message)
		if error != nil {
			println(error.Error())
		}
		playMP3()
		clearLine()
		fmt.Printf("<%v>: %v", message.Sender, message.Message)
		fmt.Printf("\033[32m<%v>:\033[0m ", conn.LocalAddr().String())
	}
}
