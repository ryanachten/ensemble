package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "ensemble/models"
	services "ensemble/services"
)

func GetBand(c *gin.Context) {
	name := c.Query("name")
	bandName := name + " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results
	rawGraph := services.GetBandGraph(bandName)
	clientGraph := models.FormatClientGraph(rawGraph)
	c.IndentedJSON(http.StatusOK, clientGraph)
}
