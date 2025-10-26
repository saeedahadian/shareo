package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"golang.org/x/term"
)

var (
	msgChan chan string = make(chan string)
)

const (
	inputPrompt = "\033[32m> \033[0m"
	cmdPrompt   = "\033[36m/ \033[0m"
	serverPrompt = "\033[33m< \033[0m"
)

type inputMode int

const (
	normalMode inputMode = iota
	commandMode
	serverMode
)

const (
	ctrlC      byte = 3
	ctrlU      byte = 21
	enter      byte = 13
	backspace1 byte = 127
	backspace2 byte = 8
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

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	mode := normalMode
	buffer := []rune{}

	showPrompt(mode)

	buf := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		char := buf[0]

		switch char {
		case ctrlC:
			fmt.Println("\r")
			return
		case ctrlU:
			buffer = []rune{}
			clearLine()
			showPrompt(mode)
		case enter:
			fmt.Print("\r\n")
			if len(buffer) > 0 {
				line := string(buffer)
				if mode == commandMode {
					line = "/" + line
				}
				msgChan <- line
				buffer = []rune{}
			}
			mode = normalMode
			// showPrompt(mode)
		case backspace1, backspace2:
			if len(buffer) > 0 {
				buffer = buffer[:len(buffer)-1]
				fmt.Printf("\b \b")
			} else if mode == commandMode {
				mode = normalMode
				clearLine()
				showPrompt(mode)
			}
		case '/':
			if len(buffer) == 0 && mode == normalMode {
				mode = commandMode
				clearLine()
				showPrompt(mode)
			} else {
				buffer = append(buffer, rune(char))
				fmt.Printf("%c", char)
			}
		case '>':
			if len(buffer) == 0 && mode == commandMode {
				mode = normalMode
				clearLine()
				showPrompt(mode)
			} else {
				buffer = append(buffer, rune(char))
				fmt.Printf("%c", char)
			}
		default:
			buffer = append(buffer, rune(char))
			fmt.Printf("%c", char)
		}
	}
}

func handleServerResp(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		showPrompt(serverMode)
		fmt.Printf("%s\r\n", text)
		showPrompt(normalMode)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func showPrompt(mode inputMode) {
	switch mode {
	case normalMode:
		fmt.Print(inputPrompt)
	case commandMode:
		fmt.Print(cmdPrompt)
	case serverMode:
		fmt.Print(serverPrompt)
	}
}

func clearLine() {
	fmt.Print("\r\033[K")
}
