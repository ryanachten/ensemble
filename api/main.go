package main

import (
	"github.com/gin-gonic/gin"

	controllers "ensemble/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/bands", controllers.GetBand)
	router.Run("localhost:8080")
}
