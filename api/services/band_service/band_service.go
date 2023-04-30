package band_service

import (
	"ensemble/clients"
	"ensemble/models"
	"ensemble/services"
	"ensemble/services/band_service/strategies"
)

// Builds a band graph using a given graph strategy
func BuildBandGraph(strategy models.GraphStrategy, bandName string, degreesOfSeparation int) (*models.ClientGraph, error) {

	searchBandName := bandName + " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results. TODO: check if this is already present in string
	searchResults, err := clients.GetSearchResults(searchBandName)
	if err != nil {
		return nil, err
	}

	maxLayers := services.GetMaxLayers(degreesOfSeparation)
	requestUrl := clients.GetPageUrl(searchResults[0].Title)
	scraper := models.NewWikiScraper()

	var clientGraph models.ClientGraph
	switch strategy {
	case models.InSync:
		{
			var inSyncGraph = strategies.BuildSequentialBandGraph(bandName, requestUrl, &models.InSyncGraph{}, scraper, maxLayers)
			clientGraph = inSyncGraph.ToClientGraph()
		}
	case models.SyncMap:
		{
			var syncMapGraph = strategies.BuildConcurrentBandGraph(bandName, requestUrl, models.NewSyncGraph(), scraper, maxLayers)
			clientGraph = syncMapGraph.ToClientGraph()
		}
	default:
		{
			var mutexGraph = strategies.BuildConcurrentBandGraph(bandName, requestUrl, models.NewMutexGraph(), scraper, maxLayers)
			clientGraph = mutexGraph.ToClientGraph()
		}
	}

	return &clientGraph, err
}
