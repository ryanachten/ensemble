package strategies

import (
	"ensemble/models"
	"sync"
)

// Recursively builds a genre graph using concurrent requests
func BuildConcurrentGenreGraph(genreName string, genreUrl string, graph models.ConcurrentGraph, scraper models.WikiScraper, maxLayers int) models.ConcurrentGraph {
	graph.AddVertex(genreName, models.VertexData{Type: models.Genre})

	var requests sync.WaitGroup
	getConcurrentGenre(genreName, genreUrl, graph, scraper, 0, maxLayers, &requests)
	requests.Wait() // wait for graph requests to complete
	graph.Wait()    // wait for graph updates to complete
	return graph
}

func getConcurrentGenre(genreName string, genreUrl string, graph models.ConcurrentGraph, scraper models.WikiScraper, layer int, maxLayers int, waitGroup *sync.WaitGroup) models.Graph {
	if layer > maxLayers {
		return graph
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(genreName)
	if vertexExists {
		return graph
	}

	// Add an entry to the wait group and defer removing entry until function completes
	waitGroup.Add(1)
	defer waitGroup.Done()

	metadata := scraper.GetGenreMetadata(genreUrl)
	graph.UpdateVertexData(genreName, metadata.ImageUrl)

	for _, link := range metadata.StylisticOrigins {
		graph.AddVertex(link.Title, models.VertexData{Type: models.Genre, Url: link.Url})
		graph.AddEdge(genreName, link.Title, "stylistic origin")
		if link.Url != nil {
			go getConcurrentGenre(link.Title, *link.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}
	for _, link := range metadata.DerivativeForms {
		graph.AddVertex(link.Title, models.VertexData{Type: models.Genre, Url: link.Url})
		graph.AddEdge(genreName, link.Title, "derivative form")
		if link.Url != nil {
			go getConcurrentGenre(link.Title, *link.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}
	for _, link := range metadata.Subgenres {
		graph.AddVertex(link.Title, models.VertexData{Type: models.Genre, Url: link.Url})
		graph.AddEdge(genreName, link.Title, "subgenre")
		if link.Url != nil {
			go getConcurrentGenre(link.Title, *link.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}
	for _, link := range metadata.FusionGenres {
		graph.AddVertex(link.Title, models.VertexData{Type: models.Genre, Url: link.Url})
		graph.AddEdge(genreName, link.Title, "fusion genre")
		if link.Url != nil {
			go getConcurrentGenre(link.Title, *link.Url, graph, scraper, layer+1, maxLayers, waitGroup)
		}
	}

	return graph
}
