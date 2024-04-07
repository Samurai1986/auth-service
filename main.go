package main

import (
	"fmt"

	"github.com/Samurai1986/auth-service/controller"
	"github.com/Samurai1986/auth-service/view"
	"github.com/gin-gonic/gin"
)

func main() {
	// load configuration
	config := controller.NewConfig()
	fmt.Print(config)
	//TODO: init database
	db, err := controller.InitDatabase(config)
	fmt.Print(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//TODO: check migrations

	//TODO: init app routers
	if controller.EnvironmentMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	view.Router(r)
	
	//TODO: add loggers to router

	//TODO: start server
	r.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))

}
