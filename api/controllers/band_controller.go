package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ensemble/services"
)

func GetBand(c *gin.Context) {
	params := ParseQueryParameters(c)
	clientGraph, err := services.BuildBandGraph(params.Strategy, params.Name, params.DegreesOfSeparation)

	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, clientGraph)
}
