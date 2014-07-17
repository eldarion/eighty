package main

import (
	"eighty/log"
	"eighty/server"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 80, "port to listen on for HTTP connections")
	flag.Parse()
	server := &server.HttpServer{
		Addr: fmt.Sprintf(":%d", port),
	}
	setupSignals(server)
	server.ListenAndServe()
	log.Infof("stopping server")
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
