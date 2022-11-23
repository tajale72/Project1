package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"myapi/controller"
	"myapi/db"
)

func PostgressClient() (*sql.DB, error) {
	conn := "host=localhost port=5432 user=romittajale dbname=finance sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")
	return db, nil
}

func (r *Router) InitializeRouter() *Router {
	mongoclient, _ := db.Mongo()
	d, _ := db.PostgressClient()

	db := &db.Service{Db: d, Name: "romit", Mongoclient: mongoclient}
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
