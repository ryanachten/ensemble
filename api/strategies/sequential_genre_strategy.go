package strategies

import (
	"ensemble/models"
)

// Recursively builds a genre graph in sequence without concurrent requests
func BuildSequentialGenreGraph(genreName string, genreUrl string, graph models.Graph, scraper models.WikiScraper, maxLayers int) models.Graph {
	graph.AddVertex(genreName, models.VertexData{Type: models.Genre})

	return getSequentialGenre(genreName, genreUrl, graph, scraper, 0, maxLayers)
}

func getSequentialGenre(genreName string, genreUrl string, graph models.Graph, scraper models.WikiScraper, layer int, maxLayers int) models.Graph {
	if layer > maxLayers {
		return graph
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(genreName)
	if vertexExists {
		return graph
	}

	metadata := scraper.GetGenreMetadata(genreUrl)
	graph.UpdateVertexData(genreName, metadata.ImageUrl)

	addSequentialEdges(metadata.StylisticOrigins, genreName, "stylistic origins", graph, scraper, layer, maxLayers)
	addSequentialEdges(metadata.DerivativeForms, genreName, "derivative form", graph, scraper, layer, maxLayers)
	addSequentialEdges(metadata.Subgenres, genreName, "subgenres", graph, scraper, layer, maxLayers)
	addSequentialEdges(metadata.FusionGenres, genreName, "fusion genres", graph, scraper, layer, maxLayers)

	return graph
}

func addSequentialEdges(links []models.Link, genreName, edgeLabel string, graph models.Graph, scraper models.WikiScraper, layer, maxLayers int) {
	for _, link := range links {
		graph.AddVertex(link.Title, models.VertexData{Type: models.Genre, Url: link.Url})
		graph.AddEdge(genreName, link.Title, edgeLabel)
		if link.Url != nil {
			getSequentialGenre(link.Title, *link.Url, graph, scraper, layer+1, maxLayers)
		}
	}
}
