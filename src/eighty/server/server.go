package server

import (
	"log"
	"net"
	"net/http"
	"sync"
)

type HttpServer struct {
	Addr            string
	Handler         http.Handler
	Shutdown        chan bool
	wg              sync.WaitGroup
	shutdownHandler func()
	server          http.Server
}

func (s *HttpServer) ListenAndServe() error {
	s.Shutdown = make(chan bool)
	netListener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	listener := &HttpServerListener{netListener, true, s}
	err = s.Serve(listener, s.Handler)
	return err
}

func (s *HttpServer) Serve(listener net.Listener, handler http.Handler) error {
	s.shutdownHandler = func() {
		log.Printf("shutdown requested")
		listener.Close()
	}
	s.listenForShutdown()
	s.server = http.Server{
		Handler: handler,
		ConnState: func(conn net.Conn, state http.ConnState) {
			switch state {
			case http.StateNew:
				s.StartRoutine()
			case http.StateClosed, http.StateHijacked:
				s.FinishRoutine()
			}
		},
	}
	err := s.server.Serve(listener)
	if err != nil {
		if _, ok := err.(listenerAlreadyClosed); ok {
			s.wg.Wait()
			return nil
		}
	}
	return err
}

func (s *HttpServer) SetKeepAlivesEnabled(v bool) {
	s.server.SetKeepAlivesEnabled(v)
}

func (s *HttpServer) StartRoutine() {
	s.wg.Add(1)
}

func (s *HttpServer) FinishRoutine() {
	s.wg.Done()
}

func (s *HttpServer) listenForShutdown() {
	go func() {
		<-s.Shutdown
		s.shutdownHandler()
	}()
}
