package main

import (
	"github.com/KyriakosMilad/memdb/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	s := server.InitServer()
	go s.Run()

	select {
	case <-stop:
		s.Stop()
	}
}
