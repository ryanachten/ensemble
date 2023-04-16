package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "ensemble/services"
)

func GetBand(c *gin.Context) {
	name := c.Query("name")
	bandName := name + " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results
	metadata := services.GetBandGraph(bandName)
	c.IndentedJSON(http.StatusOK, metadata)
}
