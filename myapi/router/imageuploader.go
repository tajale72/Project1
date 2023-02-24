package router

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (r *Router) ImageUploader(c *gin.Context) {

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		log.Println("error gettig the image", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	f, err := os.Create("image.jpg")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()

	io.Copy(f, file)
	c.String(http.StatusOK, "File uploaded successfully")

}
