package services

import (
	"ensemble/clients"
	"ensemble/models"
	"ensemble/strategies"
)

// Builds a artist graph using a given graph strategy
func BuildArtistGraph(strategy models.GraphStrategy, artistName string, degreesOfSeparation int) (*models.ClientGraph, error) {

	searchResults, err := clients.GetSearchResults(artistName)
	if err != nil {
		return nil, err
	}

	searchResultTitle := searchResults[0].Title
	maxLayers := getMaxLayers(degreesOfSeparation)
	requestUrl := clients.GetPageUrl(searchResultTitle)
	scraper := models.NewWikiScraper()

	var clientGraph models.ClientGraph
	switch strategy {
	case models.Sequential:
		{
			var sequentialGraph = strategies.BuildSequentialArtistGraph(searchResultTitle, requestUrl, &models.SequentialGraph{}, scraper, maxLayers)
			clientGraph = sequentialGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = strategies.BuildConcurrentArtistGraph(searchResultTitle, requestUrl, models.NewSyncMapGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = strategies.BuildConcurrentArtistGraph(searchResultTitle, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
