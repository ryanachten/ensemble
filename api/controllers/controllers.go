package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	models "ensemble/models"
	services "ensemble/services"
)

func GetBand(c *gin.Context) {
	bandName := c.Query("name")
	degreesOfSeparationQuery := c.Query("degreesOfSeparation")
	degreesOfSeparation, err := strconv.Atoi(degreesOfSeparationQuery)
	if err != nil {
		degreesOfSeparation = 1
	}
	rawGraph := services.SearchBandGraph(bandName, degreesOfSeparation)
	clientGraph := models.FormatClientGraph(rawGraph)
	c.IndentedJSON(http.StatusOK, clientGraph)
}
