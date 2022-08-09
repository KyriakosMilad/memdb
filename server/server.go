package server

import "net"

type client struct {
	id   int
	conn net.Conn
}
