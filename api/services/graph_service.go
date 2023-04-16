package services

import (
	clients "ensemble/clients"
	models "ensemble/models"
)

func GetBandGraph(bandName string) *models.Graph {
	searchResults, err := clients.GetSearchResults(bandName)
	if err != nil {
		return nil
	}
	metadata := ScrapeBandMetadata(searchResults[0].Title)

	var graph models.Graph
	graph.AddVertex(bandName, models.VertexData{Type: models.Band})

	for _, member := range metadata.Members {
		graph.AddVertex(member.Title, models.VertexData{Type: models.Artist, Url: member.Url})
		graph.AddEdge(bandName, member.Title, "member")
		if member.Url != nil {
			GetArtistGraph(member.Title, *member.Url, &graph)
		}
	}
	for _, pastMember := range metadata.PastMembers {
		graph.AddVertex(pastMember.Title, models.VertexData{Type: models.Artist, Url: pastMember.Url})
		graph.AddEdge(bandName, pastMember.Title, "past member")
		// TODO: can these be parallelized or something?
		if pastMember.Url != nil {
			GetArtistGraph(pastMember.Title, *pastMember.Url, &graph)
		}
	}
	return &graph
}

func GetArtistGraph(artistName, artistUrl string, graph *models.Graph) {
	metadata := ScrapeArtistMetadata(artistUrl)

	for _, member := range metadata.MemberOf {
		graph.AddVertex(member.Title, models.VertexData{Type: models.Band, Url: member.Url})
		graph.AddEdge(artistName, member.Title, "member of")
	}
	for _, pastMember := range metadata.FormerlyOf {
		graph.AddVertex(pastMember.Title, models.VertexData{Type: models.Band, Url: pastMember.Url})
		graph.AddEdge(artistName, pastMember.Title, "formerly of")
	}
}
