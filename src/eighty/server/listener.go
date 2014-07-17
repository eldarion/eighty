package server

import (
	"net"
)

type HttpServerListener struct {
	net.Listener
	open   bool
	server *HttpServer
}

func (l *HttpServerListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		if !l.open {
			err = listenerAlreadyClosed{err}
		}
		return nil, err
	}
	return conn, nil
}

func (l *HttpServerListener) Close() error {
	if !l.open {
		return nil
	}
	l.open = false
	l.server.SetKeepAlivesEnabled(false)
	err := l.Listener.Close()
	return err
}

type listenerAlreadyClosed struct {
	error
}
