package strategies

import (
	"ensemble/models"
)

// Recursively builds a artist graph in sequence without concurrent requests
func BuildSequentialArtistGraph(artistName string, artistUrl string, graph models.Graph, scraper models.WikiScraper, maxLayers int) models.Graph {
	graph.AddVertex(artistName, models.VertexData{Type: models.Artist})

	return getSequentialArtist(artistName, artistUrl, graph, scraper, 0, maxLayers)
}

func getSequentialArtist(artistName, artistUrl string, graph models.Graph, scraper models.WikiScraper, layer int, maxLayers int) models.Graph {
	if layer > maxLayers {
		return graph
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(artistName)
	if vertexExists {
		return graph
	}

	metadata := scraper.GetArtistMetadata(artistUrl)
	graph.UpdateVertexData(artistName, metadata.ImageUrl)

	for _, currentBand := range metadata.MemberOf {
		graph.AddVertex(currentBand.Title, models.VertexData{Type: models.Band, Url: currentBand.Url})
		graph.AddEdge(artistName, currentBand.Title, "member of")
		if currentBand.Url != nil {
			getSequentialBand(currentBand.Title, *currentBand.Url, graph, scraper, layer+1, maxLayers)
		}
	}
	for _, formerBand := range metadata.FormerlyOf {
		graph.AddVertex(formerBand.Title, models.VertexData{Type: models.Band, Url: formerBand.Url})
		graph.AddEdge(artistName, formerBand.Title, "formerly of")
		if formerBand.Url != nil {
			getSequentialBand(formerBand.Title, *formerBand.Url, graph, scraper, layer+1, maxLayers)
		}
	}
	return graph
}
