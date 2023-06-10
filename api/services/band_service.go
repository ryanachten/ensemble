package services

import (
	"ensemble/clients"
	"ensemble/models"
	"ensemble/strategies"
	"strings"
)

// Builds a band graph using a given graph strategy
func BuildBandGraph(strategy models.GraphStrategy, bandName string, degreesOfSeparation int) (*models.ClientGraph, error) {

	suffix := " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results. TODO: check if this is already present in string
	searchBandName := bandName + suffix
	searchResults, err := clients.GetSearchResults(searchBandName)
	if err != nil {
		return nil, err
	}

	maxLayers := getMaxLayers(degreesOfSeparation)
	requestUrl := clients.GetPageUrl(searchResults[0].Title)
	scraper := models.NewWikiScraper()
	formattedTitle, _ := strings.CutSuffix(searchResults[0].Title, suffix)

	var clientGraph models.ClientGraph
	switch strategy {
	case models.InSync:
		{
			var inSyncGraph = strategies.BuildSequentialBandGraph(formattedTitle, requestUrl, &models.InSyncGraph{}, scraper, maxLayers)
			clientGraph = inSyncGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = strategies.BuildConcurrentBandGraph(formattedTitle, requestUrl, models.NewSyncGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = strategies.BuildConcurrentBandGraph(formattedTitle, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
