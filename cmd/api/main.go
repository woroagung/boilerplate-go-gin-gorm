package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/woroagung/boilerplate-go-gin-gorm/internal/config"
	"github.com/woroagung/boilerplate-go-gin-gorm/internal/route"
)

func main() {
	// load env
	config.LoadEnv()

	// set gin environement
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode("release")
	}

	// set database
	config.LoadDatabase()

	// set router
	r := gin.Default()

	// register route
	route.LoadRoute(r)

	// run server
	appPort := os.Getenv("APP_PORT")

	err := r.Run(":" + appPort)

	if err != nil {
		log.Fatal("Failed to run server: ", err)
	} else {
		log.Println("Server running on port: ", appPort)
	}
}
