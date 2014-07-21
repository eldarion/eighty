package management

import (
	"eighty/engine"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ManagementApi struct {
	Engine *engine.Engine
}

func Handler(e *engine.Engine) http.Handler {
	r := gin.New()
	api := ManagementApi{Engine: e}
	api.Setup(r)
	return r
}

func (api *ManagementApi) Setup(r *gin.Engine) {
	r.GET("/vhosts", api.vhostsList)
}

func (api *ManagementApi) vhostsList(c *gin.Context) {
	var payload struct {
		Vhosts map[string]string
	}
	payload.Vhosts = make(map[string]string)
	for host, vhost := range api.Engine.Vhosts {
		payload.Vhosts[host] = vhost.Mode
	}
	c.JSON(200, payload)
}
