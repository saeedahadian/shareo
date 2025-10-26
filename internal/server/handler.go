package server

import "net"

func handleConn(conn net.Conn) {
	defer conn.Close()

}
