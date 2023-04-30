package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ensemble/services/genre_service"
)

func GetGenre(c *gin.Context) {
	params := ParseQueryParameters(c)
	clientGraph, err := genre_service.BuildGenreGraph(params.Strategy, params.Name, params.DegreesOfSeparation)

	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, clientGraph)
}
