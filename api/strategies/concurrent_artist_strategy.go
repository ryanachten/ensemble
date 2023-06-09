package strategies

import (
	"ensemble/models"
	"sync"
)

// Recursively builds an artist graph using concurrent requests
func BuildConcurrentArtistGraph(artistName string, artistUrl string, graph models.ConcurrentGraph, scraper models.WikiScraper, maxLayers int) models.ConcurrentGraph {
	graph.AddVertex(artistName, models.VertexData{Type: models.Artist})

	var requests sync.WaitGroup
	getConcurrentArtist(artistName, artistUrl, graph, scraper, 0, maxLayers, &requests)
	requests.Wait() // wait for graph requests to complete
	graph.Wait()    // wait for graph updates to complete
	return graph
}

func getConcurrentArtist(artistName, artistUrl string, graph models.ConcurrentGraph, scraper models.WikiScraper, layer int, maxLayers int, waitGroup *sync.WaitGroup) models.Graph {
	if layer > maxLayers {
		return graph
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(artistName)
	if vertexExists {
		return graph
	}

	// Add an entry to the wait group and defer removing entry until function completes
	waitGroup.Add(1)
	defer waitGroup.Done()

	metadata := scraper.GetArtistMetadata(artistUrl)
	graph.UpdateVertexData(artistName, metadata.ImageUrl)

	for _, currentBand := range metadata.MemberOf {
		graph.AddVertex(currentBand.Title, models.VertexData{Type: models.Band, Url: currentBand.Url})
		graph.AddEdge(artistName, currentBand.Title, "member of")
		if currentBand.Url != nil {
			go getConcurrentBand(currentBand.Title, *currentBand.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}
	for _, formerBand := range metadata.FormerlyOf {
		graph.AddVertex(formerBand.Title, models.VertexData{Type: models.Band, Url: formerBand.Url})
		graph.AddEdge(artistName, formerBand.Title, "formerly of")
		if formerBand.Url != nil {
			go getConcurrentBand(formerBand.Title, *formerBand.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}

	return graph
}
