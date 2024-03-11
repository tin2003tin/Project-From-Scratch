package main

import (
	"fmt"
	"net"
)

// Function to handle incoming connections on a server
func handleConnection(port string) {
	// Start listening for incoming connections on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s\n", port)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection (for now, just print received data)
		handleClient(conn)
	}
}

// Function to handle individual client connections
func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read data from the connection
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Print received data
	fmt.Printf("Received data: %s\n", buffer[:n])
}

func main() {
	// Ports to listen on
	ports := []string{"1001", "1002", "1003"}

	// Start a server for each port
	for _, port := range ports {
		go handleConnection(port)
	}

	// Keep the main goroutine running
	select {}
}
