package main

import (
	"fmt"
	"net"
	"sync"
)

var clients = make(map[net.Conn]bool)
var mu sync.Mutex

func main() {

	fmt.Printf("Eneter mode:\n\t*1\tfor server\n\t*2\tfor client\n")
	var i int
	fmt.Print("Your choice: ")
	fmt.Scan(&i)
	if i == 1 {
		socket_server()
	}
	if i == 2 {
		sock_client()
	}

}
