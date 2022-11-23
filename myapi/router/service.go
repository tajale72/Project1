package router

import (
	"strings"

	"github.com/gin-gonic/gin"

	"myapi/controller"
)

const (
	Authorization = "authorization"
)

type Router struct {
	controllersvc controller.ControllerInterface
}

func Start() {
	var r Router
	router := gin.Default()

	//auth := router.Group("/v1", r.Middleware)
	r.InitializeEndpoints(router)
	r.InitializeRouter()
	router.Run()

}

func (r *Router) Middleware(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	updatedtoken := strings.Split(token, " ")
	r.controllersvc.VerifyToken(updatedtoken[1])
}
