package services

import (
	clients "ensemble/clients"
	models "ensemble/models"
	"fmt"
	"net/url"
	"sync"
)

func SearchSyncBandGraph(bandName string, degreesOfSeparation int) (*models.SyncGraph, error) {
	searchBandName := bandName + " (band)" // suffix 'band' to ensure band results appear at the top of Wikipedia search results. TODO: check if this is already present in string
	searchResults, err := clients.GetSearchResults(searchBandName)
	if err != nil {
		return nil, err
	}

	encodedTitle := url.QueryEscape(searchResults[0].Title)
	requestUrl := fmt.Sprintf("https://en.wikipedia.org/w/index.php?title=%s", encodedTitle)

	var graph models.SyncGraph
	graph.AddVertex(bandName, models.VertexData{Type: models.Band})

	maxLayers := degreesOfSeparation
	if maxLayers > MAX_LAYERS {
		maxLayers = MAX_LAYERS
	}
	var waitGroup sync.WaitGroup
	getSyncBandGraph(bandName, requestUrl, &graph, 0, maxLayers, &waitGroup)
	waitGroup.Wait()
	return &graph, err
}

func getSyncBandGraph(bandName string, bandUrl string, graph *models.SyncGraph, layer int, maxLayers int, waitGroup *sync.WaitGroup) *models.SyncGraph {
	if layer > maxLayers {
		return graph
	}
	// Add an entry to the wait group and defer removing entry until function completes
	waitGroup.Add(1)
	defer waitGroup.Done()

	metadata := ScrapeBandMetadata(bandUrl)
	graph.UpdateVertexData(bandName, metadata.ImageUrl)

	for _, member := range metadata.Members {
		graph.AddVertex(member.Title, models.VertexData{Type: models.Artist, Url: member.Url})
		graph.AddEdge(bandName, member.Title, "member")
		if member.Url != nil {
			go getSyncArtistGraph(member.Title, *member.Url, graph, layer+1, maxLayers, waitGroup)
		}
	}
	for _, pastMember := range metadata.PastMembers {
		graph.AddVertex(pastMember.Title, models.VertexData{Type: models.Artist, Url: pastMember.Url})
		graph.AddEdge(bandName, pastMember.Title, "past member")
		if pastMember.Url != nil {
			go getSyncArtistGraph(pastMember.Title, *pastMember.Url, graph, layer+1, maxLayers, waitGroup)
		}
	}
	return graph
}

func getSyncArtistGraph(artistName, artistUrl string, graph *models.SyncGraph, layer int, maxLayers int, waitGroup *sync.WaitGroup) {
	if layer > maxLayers {
		return
	}
	// Add an entry to the wait group and defer removing entry until function completes
	waitGroup.Add(1)
	defer waitGroup.Done()

	metadata := ScrapeArtistMetadata(artistUrl)
	graph.UpdateVertexData(artistName, metadata.ImageUrl)

	for _, currentBand := range metadata.MemberOf {
		graph.AddVertex(currentBand.Title, models.VertexData{Type: models.Band, Url: currentBand.Url})
		graph.AddEdge(artistName, currentBand.Title, "member of")
		if currentBand.Url != nil {
			go getSyncBandGraph(currentBand.Title, *currentBand.Url, graph, layer+1, maxLayers, waitGroup)
		}
	}
	for _, formerBand := range metadata.FormerlyOf {
		graph.AddVertex(formerBand.Title, models.VertexData{Type: models.Band, Url: formerBand.Url})
		graph.AddEdge(artistName, formerBand.Title, "formerly of")
		if formerBand.Url != nil {
			go getSyncBandGraph(formerBand.Title, *formerBand.Url, graph, layer+1, maxLayers, waitGroup)
		}
	}
}