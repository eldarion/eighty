package engine

import (
	"net/http"
)

type Vhost struct {
	Mode string
}

func (e *Engine) LookupVhost(r *http.Request) *Vhost {
	return &Vhost{}
}
