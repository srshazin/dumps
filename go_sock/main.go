package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

var clients = make(map[net.Conn]bool)
var mu sync.Mutex

type Message struct {
	Message string
	Sender  string
}

func clearLine() {
	fmt.Print("\033[2K\r")
}

func playMP3() error {
	var filePath = "top.mp3"
	// Open the MP3 file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open MP3 file: %v", err)
	}
	defer file.Close()

	// Create MP3 decoder
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		return fmt.Errorf("failed to decode MP3 file: %v", err)
	}

	// Initialize Oto context for audio playback
	context, err := oto.NewContext(44100, 2, 2, 4096) // Sample rate: 44100 Hz, 2 channels, 16-bit samples
	if err != nil {
		return fmt.Errorf("failed to create Oto context: %v", err)
	}
	defer context.Close()

	// Create a new player for the audio
	player := context.NewPlayer()
	defer player.Close()

	// Play the decoded MP3 file
	if _, err := io.Copy(player, decoder); err != nil {
		return fmt.Errorf("error playing MP3: %v", err)
	}

	return nil
}

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

	// test()

}
