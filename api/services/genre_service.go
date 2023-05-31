package services

import (
	"ensemble/clients"
	"ensemble/models"
	"ensemble/strategies"
)

func BuildGenreGraph(strategy models.GraphStrategy, genreName string, degreesOfSeparation int) (*models.ClientGraph, error) {
	searchResults, err := clients.GetSearchResults(genreName)
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
			var inSyncGraph = strategies.BuildSequentialGenreGraph(genreName, requestUrl, &models.InSyncGraph{}, scraper, maxLayers)
			clientGraph = inSyncGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = strategies.BuildConcurrentGenreGraph(genreName, requestUrl, models.NewSyncGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = strategies.BuildConcurrentGenreGraph(genreName, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
