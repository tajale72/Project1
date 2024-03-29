package main

import (
	router "interview/internal/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func setup(processes nacelle.ProcessContainer, services nacelle.ServiceContainer) error {
// 	processes.RegisterInitializer(api.NewProcess(), nacelle.WithInitializerName("process"))
// 	return nil
// }

type ConfigCORS struct {
	AllowOrigins     []string      `env:"ALLOW_ORIGINS" default:"[\"*\"]"`
	AllowMethods     []string      `env:"ALLOW_METHODS" default:"[\"*\"]"`
	AllowHeaders     []string      `env:"ALLOW_HEADERS" default:"[\"*\"]"`
	ExposeHeaders    []string      `env:"EXPOSE_HEADERS"`
	AllowCredentials bool          `env:"ALLOW_CREDENTIALS"`
	MaxAge           time.Duration `env:"MAX_AGE"`
}

func main() {
	//nacelle.NewBootstrapper("gin-router", setup).BootAndExit()

	r := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	r.Use(cors.Default())
	r.GET("/user/:name", router.GetUser)
	r.POST("/user", router.InsertUser)

	r.Run()

}
