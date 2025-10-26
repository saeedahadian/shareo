package client

import (
	"fmt"
	"net"
	"os"
)

func Connect(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Connected %s", conn.RemoteAddr().String())
}
