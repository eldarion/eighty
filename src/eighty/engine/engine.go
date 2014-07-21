package engine

import (
	"eighty/templates"
	"net/http"
)

type Engine struct {
}

func New() *Engine {
	e := &Engine{}
	return e
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vhost := e.LookupVhost(r)
	var handler http.Handler
	switch vhost.Mode {
	default:
		handler = templates.Response(http.StatusNotFound, "not_configured.html")
	}
	handler.ServeHTTP(w, r)
}
