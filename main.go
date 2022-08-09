package main

import (
	"flag"
	"github.com/KyriakosMilad/memdb/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.String("port", "3636", "memdb listening port")
	flag.Parse()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	s := server.InitServer()
	go s.Run(*port)

	select {
	case <-stop:
		s.Stop()
	}
}
