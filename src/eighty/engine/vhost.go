package engine

import (
	"net/http"
)

type Vhost struct {
	Mode string
}

func (e *Engine) LookupVhost(r *http.Request) *Vhost {
	if _, ok := e.Vhosts[r.Host]; ok {
		return e.Vhosts[r.Host]
	}
	return nil
}
