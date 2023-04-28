package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controllers "ensemble/controllers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/bands", controllers.GetBand)
	err := router.Run()
	log.Printf("Router error %v", err)
}
