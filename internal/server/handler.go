package server

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer close(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("%s: %s\n", conn.RemoteAddr().String(), message)
		fmt.Fprintln(conn, message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func close(conn net.Conn) {
	fmt.Printf("\033[31mExited\033[0m %s\n", conn.RemoteAddr().String())
	conn.Close()
}
