package router

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	controller "interview/internal/controller"
)

func GetUser(c *gin.Context) {
	origin := c.Request.Header.Get("User-Agent")
	name := c.Param("name")
	res, err := controller.GetUser(name)
	if err != nil {
		log.Println("error from controller", err)
		c.JSON(http.StatusBadRequest, err)
	}
	log.Println("origin", origin)
	c.JSON(http.StatusOK, res)
}

func InsertUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	res, err := controller.InsertUser(body)
	if err != nil {
		log.Println("error inserting the data", err)
		c.JSON(http.StatusBadRequest, err)
	}
	log.Println("inserted users")
	c.JSON(http.StatusOK, res)
}
