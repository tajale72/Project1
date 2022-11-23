package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *Router) GenerateToken(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	res, _ := r.controllersvc.GenerateToken(body)
	c.JSON(http.StatusOK, res)
}

func (r *Router) Verify(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	if len(token) == 0 {
		log.Println("error from token validation")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return

	}
	updatedtoken := strings.Split(token, " ")
	res, _ := r.controllersvc.VerifyToken(updatedtoken[1])
	c.JSON(http.StatusOK, res)
}
