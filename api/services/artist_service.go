package services

import (
	"ensemble/clients"
	"ensemble/models"
	"ensemble/strategies"
)

// Builds a artist graph using a given graph strategy
func BuildArtistGraph(strategy models.GraphStrategy, artistName string, degreesOfSeparation int) (*models.ClientGraph, error) {

	searchArtistName := artistName + " (musician)" // suffix 'musician' to ensure artist results appear at the top of Wikipedia search results. TODO: check if this is already present in string
	searchResults, err := clients.GetSearchResults(searchArtistName)
	if err != nil {
		return nil, err
	}

	maxLayers := getMaxLayers(degreesOfSeparation)
	requestUrl := clients.GetPageUrl(searchResults[0].Title)
	scraper := models.NewWikiScraper()

	var clientGraph models.ClientGraph
	switch strategy {
	case models.InSync:
		{
			var inSyncGraph = strategies.BuildSequentialArtistGraph(artistName, requestUrl, &models.InSyncGraph{}, scraper, maxLayers)
			clientGraph = inSyncGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = strategies.BuildConcurrentArtistGraph(artistName, requestUrl, models.NewSyncGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = strategies.BuildConcurrentArtistGraph(artistName, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
