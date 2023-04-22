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
	mode := c.Query("mode")
	degreesOfSeparationQuery := c.Query("degreesOfSeparation")
	degreesOfSeparation, err := strconv.Atoi(degreesOfSeparationQuery)
	if err != nil {
		degreesOfSeparation = 1
	}

	// Switch on `mode` query string for performance testing
	var clientGraph *models.ClientGraph
	switch mode {
	case "mutex":
		clientGraph, err = getMutexGraph(bandName, degreesOfSeparation)
	case "nonsync":
		clientGraph, err = getNonSyncGraph(bandName, degreesOfSeparation)
	default:
		clientGraph, err = getSyncGraph(bandName, degreesOfSeparation)
	}
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, clientGraph)
}

// Graph built using the sync graph
func getSyncGraph(bandName string, degreesOfSeparation int) (*models.ClientGraph, error) {
	rawGraph, err := services.SearchSyncBandGraph(bandName, degreesOfSeparation)
	if err != nil {
		return nil, err
	}
	clientGraph := models.ConvertSyncToClientGraph(rawGraph)
	return &clientGraph, err
}

// Graph built using the mutex graph
func getMutexGraph(bandName string, degreesOfSeparation int) (*models.ClientGraph, error) {
	rawGraph, err := services.SearchMutexBandGraph(bandName, degreesOfSeparation)
	if err != nil {
		return nil, err
	}
	clientGraph := models.ConvertMutexToClientGraph(rawGraph)
	return &clientGraph, err
}

// Graph built using the original graph
func getNonSyncGraph(bandName string, degreesOfSeparation int) (*models.ClientGraph, error) {
	rawGraph, err := services.SearchBandGraph(bandName, degreesOfSeparation)
	if err != nil {
		return nil, err
	}
	clientGraph := models.ConvertToClientGraph(rawGraph)
	return &clientGraph, err
}
