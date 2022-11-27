package router

import (
	"log"

	"github.com/gin-gonic/gin"

	"myapi/controller"
	"myapi/db"
)

func (r *Router) InitializeServices() *Router {
	//Creating a monog client
	mongoclient, err := db.Mongo()
	Error(err)
	//Creatng a postgres client
	d, err := db.PostgressClient()
	Error(err)
	//pointing to the methods of the structs
	db := &db.Service{Db: d, Name: "romit", Mongoclient: mongoclient}

	r.controllersvc = &controller.Service{DB: db}
	return r
}

func (r *Router) InitEndpoints(auth *gin.RouterGroup) {
	auth.POST("/finance", r.StoreFinancial)
	auth.GET("/finance/:id", r.GetFinancial)

}

func (r *Router) InitTokenEndpoints(router *gin.Engine) {
	router.GET("/", r.Hello)
	router.POST("/verifytoken", r.Verify)
	router.POST("/token", r.GenerateToken)

}

//Error is a function which takes in an error as a paramter and return an err
func Error(err error) error {
	if err != nil {
		log.Println("error from the router: " + err.Error())
		return err
	}
	return nil
}
