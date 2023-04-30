package strategies

import (
	"ensemble/models"
)

// Recursively builds a band graph in sequence without concurrent requests
func BuildSequentialBandGraph(bandName string, bandUrl string, graph models.Graph, scraper models.WikiScraper, maxLayers int) models.Graph {
	graph.AddVertex(bandName, models.VertexData{Type: models.Band})

	return getSequentialBand(bandName, bandUrl, graph, scraper, 0, maxLayers)
}

func getSequentialBand(bandName string, bandUrl string, graph models.Graph, scraper models.WikiScraper, layer int, maxLayers int) models.Graph {
	if layer > maxLayers {
		return graph
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(bandName)
	if vertexExists {
		return graph
	}

	metadata := scraper.GetBandMetadata(bandUrl)
	graph.UpdateVertexData(bandName, metadata.ImageUrl)

	for _, member := range metadata.Members {
		graph.AddVertex(member.Title, models.VertexData{Type: models.Artist, Url: member.Url})
		graph.AddEdge(bandName, member.Title, "member")
		if member.Url != nil {
			getSequentialArtist(member.Title, *member.Url, graph, scraper, layer+1, maxLayers)
		}
	}
	for _, pastMember := range metadata.PastMembers {
		graph.AddVertex(pastMember.Title, models.VertexData{Type: models.Artist, Url: pastMember.Url})
		graph.AddEdge(bandName, pastMember.Title, "past member")
		if pastMember.Url != nil {
			getSequentialArtist(pastMember.Title, *pastMember.Url, graph, scraper, layer+1, maxLayers)
		}
	}
	for _, genre := range metadata.Genres {
		graph.AddVertex(genre.Title, models.VertexData{Type: models.Genre, Url: genre.Url})
		graph.AddEdge(bandName, genre.Title, "genre")
	}
	return graph
}

func getSequentialArtist(artistName, artistUrl string, graph models.Graph, scraper models.WikiScraper, layer int, maxLayers int) {
	if layer > maxLayers {
		return
	}

	// If the vertex exists and complete, we don't need to revisit it
	vertexExists := graph.HasCompleteVertex(artistName)
	if vertexExists {
		return
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
}
