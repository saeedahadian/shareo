package server

import (
	"fmt"
	"net"
	"os"
)

func Start(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Sharing server listening on :%d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("\033[32mAccepted\033[0m %s\n", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
