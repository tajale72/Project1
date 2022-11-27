package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"myapi/router"
)

func main() {
	//This is how we start this application
	AuthorizeJWT()
	router.Start()

}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		log.Println(tokenString)
		// token, err := service.JWTAuthService().ValidateToken(tokenString)
		// if token.Valid {
		// 	claims := token.Claims.(jwt.MapClaims)
		// 	fmt.Println(claims)
		// } else {
		// 	fmt.Println(err)
		// 	c.AbortWithStatus(http.StatusUnauthorized)

	}
}
