package strategies

import (
	"ensemble/models"
	"sync"
)

// Recursively builds a band graph using concurrent requests
func BuildConcurrentBandGraph(bandName string, bandUrl string, graph models.ConcurrentGraph, scraper models.WikiScraper, maxLayers int) models.ConcurrentGraph {
	graph.AddVertex(bandName, models.VertexData{Type: models.Band})

	var requests sync.WaitGroup
	getConcurrentBand(bandName, bandUrl, graph, scraper, 0, maxLayers, &requests)
	requests.Wait() // wait for graph requests to complete
	graph.Wait()    // wait for graph updates to complete
	return graph
}

// TODO: we don't actually use these return values in the concurrent operations
// whereas we do in the sequential operations - wonder if that's contributing to the problem?
func getConcurrentBand(bandName string, bandUrl string, graph models.ConcurrentGraph, scraper models.WikiScraper, layer int, maxLayers int, waitGroup *sync.WaitGroup) models.Graph {
	if layer > maxLayers {
		return graph
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(bandName)
	if vertexExists {
		return graph
	}

	// Add an entry to the wait group and defer removing entry until function completes
	waitGroup.Add(1)
	defer waitGroup.Done()

	metadata := scraper.GetBandMetadata(bandUrl)
	graph.UpdateVertexData(bandName, metadata.ImageUrl)

	for _, member := range metadata.Members {
		graph.AddVertex(member.Title, models.VertexData{Type: models.Artist, Url: member.Url})
		graph.AddEdge(bandName, member.Title, "member")
		if member.Url != nil {
			go getConcurrentArtist(member.Title, *member.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}
	for _, pastMember := range metadata.PastMembers {
		graph.AddVertex(pastMember.Title, models.VertexData{Type: models.Artist, Url: pastMember.Url})
		graph.AddEdge(bandName, pastMember.Title, "past member")
		if pastMember.Url != nil {
			go getConcurrentArtist(pastMember.Title, *pastMember.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}
	for _, genre := range metadata.Genres {
		graph.AddVertex(genre.Title, models.VertexData{Type: models.Genre, Url: genre.Url})
		graph.AddEdge(bandName, genre.Title, "genre")
	}
	return graph
}
