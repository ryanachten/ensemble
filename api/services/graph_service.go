package services

import (
	clients "ensemble/clients"
	models "ensemble/models"
	"fmt"
	"net/url"
)

const MAX_LAYERS = 10 // Hard limit to prevent hypothetical endless recursive searching

func SearchBandGraph(bandName string, degreesOfSeparation int) (*models.Graph, error) {
	searchBandName := bandName + " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results. TODO: check if this is already present in string
	searchResults, err := clients.GetSearchResults(searchBandName)
	if err != nil {
		return nil, err
	}

	encodedTitle := url.QueryEscape(searchResults[0].Title)
	requestUrl := fmt.Sprintf("https://en.wikipedia.org/w/index.php?title=%s", encodedTitle)

	var graph models.Graph
	graph.AddVertex(bandName, models.VertexData{Type: models.Band})

	maxLayers := degreesOfSeparation
	if maxLayers > MAX_LAYERS {
		maxLayers = MAX_LAYERS
	}

	return getBandGraph(bandName, requestUrl, &graph, 0, maxLayers), err
}

func getBandGraph(bandName string, bandUrl string, graph *models.Graph, layer int, maxLayers int) *models.Graph {
	if layer > maxLayers {
		return graph
	}

	metadata := ScrapeBandMetadata(bandUrl)
	graph.UpdateVertexData(bandName, metadata.ImageUrl)

	for _, member := range metadata.Members {
		graph.AddVertex(member.Title, models.VertexData{Type: models.Artist, Url: member.Url})
		graph.AddEdge(bandName, member.Title, "member")
		if member.Url != nil {
			getArtistGraph(member.Title, *member.Url, graph, layer+1, maxLayers)
		}
	}
	for _, pastMember := range metadata.PastMembers {
		graph.AddVertex(pastMember.Title, models.VertexData{Type: models.Artist, Url: pastMember.Url})
		graph.AddEdge(bandName, pastMember.Title, "past member")
		if pastMember.Url != nil {
			getArtistGraph(pastMember.Title, *pastMember.Url, graph, layer+1, maxLayers)
		}
	}
	return graph
}

func getArtistGraph(artistName, artistUrl string, graph *models.Graph, layer int, maxLayers int) {
	if layer > maxLayers {
		return
	}

	metadata := ScrapeArtistMetadata(artistUrl)
	graph.UpdateVertexData(artistName, metadata.ImageUrl)

	for _, currentBand := range metadata.MemberOf {
		graph.AddVertex(currentBand.Title, models.VertexData{Type: models.Band, Url: currentBand.Url})
		graph.AddEdge(artistName, currentBand.Title, "member of")
		if currentBand.Url != nil {
			getBandGraph(currentBand.Title, *currentBand.Url, graph, layer+1, maxLayers)
		}
	}
	for _, formerBand := range metadata.FormerlyOf {
		graph.AddVertex(formerBand.Title, models.VertexData{Type: models.Band, Url: formerBand.Url})
		graph.AddEdge(artistName, formerBand.Title, "formerly of")
		if formerBand.Url != nil {
			getBandGraph(formerBand.Title, *formerBand.Url, graph, layer+1, maxLayers)
		}
	}
}
