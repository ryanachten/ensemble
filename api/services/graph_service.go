package services

import (
	clients "ensemble/clients"
	models "ensemble/models"
	"fmt"
	"net/url"
)

const MAX_LAYERS = 10 // Hard limit to prevent hypothetical endless recursive searching

// Builds a band graph using a given graph strategy
func BuildBandGraph(strategy models.GraphStrategy, bandName string, degreesOfSeparation int) (*models.ClientGraph, error) {
	searchBandName := bandName + " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results. TODO: check if this is already present in string
	searchResults, err := clients.GetSearchResults(searchBandName)
	if err != nil {
		return nil, err
	}

	encodedTitle := url.QueryEscape(searchResults[0].Title)
	requestUrl := fmt.Sprintf("https://en.wikipedia.org/w/index.php?title=%s", encodedTitle)

	maxLayers := degreesOfSeparation
	if maxLayers > MAX_LAYERS {
		maxLayers = MAX_LAYERS
	}

	scraper := models.NewWikiScraper()

	var clientGraph models.ClientGraph
	switch strategy {
	case models.InSync:
		{
			var inSyncGraph = buildSequentialBandGraph(bandName, requestUrl, &models.InSyncGraph{}, scraper, maxLayers)
			clientGraph = inSyncGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = buildConcurrentBandGraph(bandName, requestUrl, models.NewSyncGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = buildConcurrentBandGraph(bandName, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
