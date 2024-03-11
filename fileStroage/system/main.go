package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
)

// List of node IP addresses
var nodesPorts = []string{"1001", "1002", "1003"}

// Function to distribute the file to each node
func distributeFile(fileContent []byte) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(nodesPorts))
	for index, port := range nodesPorts {
		wg.Add(1)
		go func(nodeAddr string, fileContent []byte) {
			defer wg.Done()
			err := uploadFileToNode(fileContent, nodeAddr)
			if err != nil {
				errCh <- err
			}
		}(port, fileContent[index*len(fileContent)/len(nodesPorts):(index+1)*len(fileContent)/len(nodesPorts)])
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func uploadFileToNode(fileContent []byte, port string) error {
	node_server := "localhost:" + port
	conn, err := net.Dial("tcp", node_server)
	if err != nil {
		fmt.Println("Error connecting:", err)		
		return err
	}
	conn.Write([]byte(fileContent));
	defer conn.Close()

	return nil
}

func handleConnetion(conn net.Conn) {
	var buffer bytes.Buffer
	tempBuffer := make([]byte, 1024);
	n, err := (conn).Read(tempBuffer)
	if err != nil {
		if err == io.EOF {
			return 
		}
		return 
	}
	buffer.Write(tempBuffer[:n])
	err = distributeFile(buffer.Bytes())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", ":"+"1000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	} 
	fmt.Println("Listen to port:1000")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnetion(conn);
	}
}