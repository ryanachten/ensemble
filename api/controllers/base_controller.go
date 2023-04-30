package controllers

import (
	"ensemble/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerParameters struct {
	Name                string
	DegreesOfSeparation int
	Strategy            models.GraphStrategy
}

// Parses common query parameters
func ParseQueryParameters(c *gin.Context) ControllerParameters {
	name := c.Query("name")

	degreesOfSeparationQuery := c.Query("degreesOfSeparation")
	degreesOfSeparation, err := strconv.Atoi(degreesOfSeparationQuery)
	if err != nil {
		degreesOfSeparation = 1
	}

	mode := c.Query("mode")
	strategy := models.ParseStrategyString(mode)

	return ControllerParameters{
		Name:                name,
		DegreesOfSeparation: degreesOfSeparation,
		Strategy:            strategy,
	}
}
