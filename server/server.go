package server

import (
	"fmt"
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

func InitServer() *Server {
	return &Server{
		db:      database.InitDatabase(),
		clients: map[int]*client{},
	}
}

func (s *Server) addClient(c *client) {
	s.clients[c.id] = c
	fmt.Println("client with id", c.id, "joined")
}
