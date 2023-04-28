package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ensemble/models"
	services "ensemble/services"
)

func GetBand(c *gin.Context) {
	bandName := c.Query("name")
	mode := c.Query("mode")
	degreesOfSeparationQuery := c.Query("degreesOfSeparation")

	degreesOfSeparation, err := strconv.Atoi(degreesOfSeparationQuery)
	if err != nil {
		degreesOfSeparation = 1
	}
	strategy := models.ParseStrategyString(mode)

	clientGraph, err := services.BuildBandGraph(strategy, bandName, degreesOfSeparation)

	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, clientGraph)
}
