package templates

import (
	"eighty/assets"
	"fmt"
	"net/http"
)

func Response(statusCode int, name string) http.Handler {
	data, _ := assets.Asset(fmt.Sprintf("templates/%s", name))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write(data)
	})
}
