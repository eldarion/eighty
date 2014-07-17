package main

import (
	"eighty/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := 80
	server := &server.HttpServer{
		Addr: fmt.Sprintf(":%d", port),
	}
	setupSignals(server)
	server.ListenAndServe()
	log.Printf("stopping server")
}

func setupSignals(server *server.HttpServer) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-sigc
		switch s {
		case os.Interrupt:
			server.Shutdown <- true
		case syscall.SIGTERM:
			server.Shutdown <- true
		}
	}()
}
