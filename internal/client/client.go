package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	msgChan chan string = make(chan string)
)

const (
	inputPrompt = "\033[32m> \033[0m"
)

func Connect(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Connected %s\n", conn.RemoteAddr().String())

	go handleServerResp(conn)
	go handleUserInput(msgChan)

	for message := range msgChan {
		fmt.Fprintln(conn, message)
	}
}

func handleUserInput(msgChan chan string) {
	defer close(msgChan)
	fmt.Print(inputPrompt)
	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		msgChan <- inputScanner.Text()
	}

	if err := inputScanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleServerResp(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Server: %s\n", text)
		fmt.Print(inputPrompt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
