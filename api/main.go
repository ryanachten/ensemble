package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controllers "ensemble/controllers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/bands", controllers.GetBand)
	router.Run("localhost:8080")
}
