package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Create a pipe (with a reader and writer)
	reader, writer, err := os.Pipe()
	if err != nil {
		log.Fatalf("Error Creating pipe, %v\n", err)
	}

	// Start the server in a separate goroutine
	go server(writer)

	// Start the client in the main routine
	client(reader)
}

// Simulates the server writing data to the pipe
func server(writer *os.File) {
	defer writer.Close() // Close the writer when done

	messages := []string{
		"Hello, Client!",
		"How are you?",
		"Goodbye!",
	}

	for _, msg := range messages {
		// Write message to the pipe
		fmt.Fprintf(writer, "%s\n", msg)
		time.Sleep(1 * time.Second) // Simulate some delay
	}
}

// Simulates the client reading data from the pipe
func client(reader *os.File) {
	defer reader.Close() // Close the reader when done

	buffer := make([]byte, 256) // Buffer to read messages

	for {
		// Read from the pipe
		n, err := reader.Read(buffer)
		if err == io.EOF {
			// End of file means the writer is closed
			fmt.Println("Server disconnected.")
			break
		} else if err != nil {
			// Handle any other errors
			fmt.Printf("Error reading from pipe: %v\n", err)
			break
		}

		// Print the message
		fmt.Printf("Client received: %s", string(buffer[:n]))
	}
}
