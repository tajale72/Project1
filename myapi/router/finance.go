package router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) StoreFinancial(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	res, err := r.controllersvc.StoreFinancial(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, res)

}
