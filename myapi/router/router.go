package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"myapi/controller"
	"myapi/db"
)

func (r *Router) InitializeRouter() *Router {
	d, _ := db.PostgressClient()

	db := &db.Service{Db: d, Name: "romit"}
	r.controllersvc = &controller.Service{DB: db}

	return r
}

func (r *Router) InitializeEndpoints(router *gin.Engine) {
	router.GET("/", r.Hello)
	router.GET("/verifytoken", r.Verify)
	router.GET("/token", r.GenerateToken)
	router.POST("/finance", r.StoreFinancial)

}

func (r *Router) Hello(c *gin.Context) {
	res, err := r.controllersvc.Hello()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, res)
}
