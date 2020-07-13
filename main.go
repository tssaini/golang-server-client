package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage go run . server")
	}
	port := 8080
	switch os.Args[1] {
	case "server":
		server(port)
	case "client":
		client(port)
	default:
		log.Fatalf("Invalid arg %v", os.Args[1])
	}
}

func server(port int) {
	log.Printf("Starting server on port %v", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Connection closed to %v", conn.RemoteAddr())
			return
		}
		fmt.Print("Message Received:", string(message))
	}
}

func client(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%v", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		_, err := fmt.Fprintf(conn, text+"\n")
		if err != nil {
			log.Fatalf("Unable to send message to sever")
		}
	}
}
