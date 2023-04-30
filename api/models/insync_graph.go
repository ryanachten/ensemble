package models

// Adjacency list graph. Does not cater for concurrent read and writes
type InSyncGraph struct {
	Vertices map[string]*Vertex
}

func (graph *InSyncGraph) AddVertex(key string, data VertexData) {
	if graph.Vertices == nil {
		graph.Vertices = map[string]*Vertex{}
	}
	_, vertexExists := graph.Vertices[key]
	if vertexExists {
		return
	}
	graph.Vertices[key] = &Vertex{Data: data, Edges: map[string]*Edge{}}
}

func (graph *InSyncGraph) UpdateVertexData(key string, imageUrl string) {
	graph.Vertices[key].Data.ImageUrl = imageUrl
}

func (graph *InSyncGraph) HasCompleteVertex(key string) bool {
	vertex, exists := graph.Vertices[key]
	if exists {
		return vertex.Data.IsComplete
	}
	return false
}

func (graph *InSyncGraph) AddEdge(srcKey, destKey, label string) {
	// Ensure src and dest keys exist
	_, srcVertexExists := graph.Vertices[srcKey]
	_, destVertexExists := graph.Vertices[destKey]
	if !srcVertexExists || !destVertexExists {
		return
	}

	graph.Vertices[srcKey].Edges[destKey] = &Edge{Label: label}
}

// Formats Graph for client consumption
func (graph *InSyncGraph) ToClientGraph() ClientGraph {
	var nodes []ClientNode
	var edges []ClientEdge

	for vertexKey, vertexValue := range graph.Vertices {
		nodes = append(nodes, ClientNode{
			Data: ClientNodeData{
				Id:       vertexKey,
				Label:    vertexKey,
				Type:     vertexValue.Data.Type,
				ImageUrl: vertexValue.Data.ImageUrl,
			},
		})
		for edgeKey, edgeValue := range vertexValue.Edges {
			edges = append(edges, ClientEdge{
				Data: ClientEdgeData{
					Source: vertexKey,
					Target: edgeKey,
					Label:  edgeValue.Label,
				},
			})
		}
	}
	return ClientGraph{
		NodeCount: len(nodes),
		EdgeCount: len(edges),
		Nodes:     nodes,
		Edges:     edges,
	}
}
