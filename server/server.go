package server

import (
	"bufio"
	"fmt"
	"github.com/KyriakosMilad/memdb/database"
	"log"
	"net"
	"os"
	"reflect"
	"strings"
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

func (s *Server) removeClient(c *client) {
	if _, ok := s.clients[c.id]; !ok {
		return
	}
	c.disconnect()
	delete(s.clients, c.id)
	fmt.Println("client with id", c.id, "disconnected")
}

func (s *Server) removeAllClients() {
	for _, c := range s.clients {
		s.removeClient(c)
	}
}

func (s *Server) Run() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":3636")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	s.l = l
	// Close the listener when the application closes.
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			if strings.Split(err.Error(), ": ")[1] != "use of closed network connection" {
				panic("Can't close connection:" + err.Error())
			}
		}
	}(l)
	fmt.Println("Listening on: :3636")

	id := 0
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			if strings.Split(err.Error(), ": ")[1] == "use of closed network connection" {
				break
			} else {
				fmt.Println("Error accepting connection: ", err.Error())
				continue
			}
		}
		c := &client{conn: conn, id: id}
		s.addClient(c)
		s.write(c, "successfully connected to memdb")
		id++

		// Handle connections in a new goroutine.
		go s.handleClient(c)
	}
}

func (s *Server) handleClient(c *client) {
	scanner := bufio.NewScanner(c.conn)

ScannerLoop:
	for scanner.Scan() {
		s.write(&client{id: c.id}, scanner.Text())
		l := strings.ToLower(strings.TrimSpace(scanner.Text()))
		values := strings.Split(l, " ")

		switch {
		case len(values) == 3 && values[0] == "set":
			s.db.Set(values[1], values[2])
			s.write(c, "OK")
		case len(values) == 2 && values[0] == "get":
			k := values[1]
			val, found := s.db.Get(k)
			if !found {
				s.write(c, fmt.Sprintf("key %s not found", k))
			} else {
				s.write(c, val)
			}
		case len(values) == 2 && values[0] == "delete":
			s.db.Delete(values[1])
			s.write(c, "OK")
		case len(values) == 1 && values[0] == "exit":
			break ScannerLoop
		default:
			s.write(c, fmt.Sprintf("UNKNOWN: %s", l))
		}
	}

	// if scanner.Scan stops means client has disconnected
	s.removeClient(c)
}

func (s *Server) write(c *client, str string) {
	direction := "sent:"
	if !reflect.ValueOf(*c).Field(1).IsZero() {
		_, err := fmt.Fprintf(c.conn, "%s\n$ ", str)
		if err != nil {
			log.Fatal(err)
		}
		direction = "received:"
	}
	fmt.Println("client", c.id, direction, str)
}

func (s *Server) Stop() {
	fmt.Println("\nstopping memdb")
	fmt.Println("removing all clients")
	s.removeAllClients()
	fmt.Println("removed all clients successfully")
	fmt.Println("saving database on the desk")
	s.db.Save()
	fmt.Println("closing the tcp listener")
	err := s.l.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("closed tcp listener successfully")
	fmt.Println("exiting")
}
