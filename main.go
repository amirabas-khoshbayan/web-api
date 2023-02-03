package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"web-api/config"
)

func main() {

	//get config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	//init server
	server := echo.New()
	//routing

	//middleware

	//start server
	server.Start(":" + config.AppConfig.Server.Port)
}
