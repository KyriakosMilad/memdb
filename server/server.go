package server

import "net"

type client struct {
	id   int
	conn net.Conn
}

func (c *client) disconnect() {
	err := c.conn.Close()
	if err != nil {
		panic(err)
	}
}
