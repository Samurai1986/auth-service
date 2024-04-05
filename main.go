package main

import (
	"fmt"

	"github.com/Samurai1986/auth-service/controller"
)

func main() {
	// load configuration
	config := controller.NewConfig()
	fmt.Print(config)
	//TODO: init database
	
	//TODO: check migrations

	//TODO: init app routers

	//TODO: add loggers to router

	//TODO: start server

}
