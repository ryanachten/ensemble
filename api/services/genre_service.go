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

	searchResultTitle := searchResults[0].Title
	maxLayers := getMaxLayers(degreesOfSeparation)
	requestUrl := clients.GetPageUrl(searchResultTitle)
	scraper := models.NewWikiScraper()

	var clientGraph models.ClientGraph
	switch strategy {
	case models.Sequential:
		{
			var sequentialGraph = strategies.BuildSequentialGenreGraph(searchResultTitle, requestUrl, &models.SequentialGraph{}, scraper, maxLayers)
			clientGraph = sequentialGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = strategies.BuildConcurrentGenreGraph(searchResultTitle, requestUrl, models.NewSyncGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = strategies.BuildConcurrentGenreGraph(searchResultTitle, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
