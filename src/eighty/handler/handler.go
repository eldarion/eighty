package handler

import (
	"eighty/engine"
	"eighty/management"
	"eighty/templates"
	"net/http"
)

func EngineHandler(e *engine.Engine) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vhost := e.LookupVhost(r)
		var handler http.Handler
		if vhost != nil {
			switch vhost.Mode {
			case "management":
				handler = management.Handler(e)
			default:
				handler = templates.Response(http.StatusNotFound, "not_configured.html")
			}
		} else {
			handler = templates.Response(http.StatusNotFound, "not_configured.html")
		}
		handler.ServeHTTP(w, r)
	})
}
