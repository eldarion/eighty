package main

import (
	"eighty/engine"
	"eighty/handler"
	"eighty/log"
	"eighty/server"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var bind string
	flag.StringVar(&bind, "bind", ":80", "socket to bind (e.g., localhost:8000 or einhorn@0)")
	flag.Parse()
	engine := engine.New()
	server := &server.HttpServer{
		Bind:    bind,
		Handler: handler.EngineHandler(engine),
	}
	setupSignals(server)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Infof("stopping server")
}

func setupSignals(server *server.HttpServer) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR2)
	go func() {
		s := <-sigc
		switch s {
		case os.Interrupt:
			server.Shutdown <- true
		case syscall.SIGTERM:
			server.Shutdown <- true
		case syscall.SIGUSR2:
			server.Shutdown <- true
		}
	}()
}
