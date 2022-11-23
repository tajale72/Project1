package router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Router struct {
// 	db *db.Business
// }

func InsertUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	id, err := a.db.InsertUser(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, "data not inserted")
	}
	c.JSONP(http.StatusAccepted, id)
}

func GetUser(c *gin.Context) {
	res, err := a.db.GetUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSONP(http.StatusAccepted, res)
}
