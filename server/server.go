package server

import (
	"github.com/KyriakosMilad/memdb/database"
	"net"
)

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

type Server struct {
	l       net.Listener
	db      *database.Database
	clients map[int]*client
}
