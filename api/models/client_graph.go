package models

type ClientNodeData struct {
	Id       string     `json:"id,omitempty"`
	Label    string     `json:"label,omitempty"`
	Type     VertexType `json:"type,omitempty"`
	ImageUrl string     `json:"imageUrl,omitempty"`
}

type ClientNode struct {
	Data ClientNodeData `json:"data,omitempty"`
}

type ClientEdgeData struct {
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`
	Label  string `json:"label,omitempty"`
}

type ClientEdge struct {
	Data ClientEdgeData `json:"data,omitempty"`
}

// Graph DTO formatted for client consumption
type ClientGraph struct {
	NodeCount int          `json:"nodeCount,omitempty"`
	EdgeCount int          `json:"edgeCount,omitempty"`
	Nodes     []ClientNode `json:"nodes,omitempty"`
	Edges     []ClientEdge `json:"edges,omitempty"`
}

// Formats Graph for client consumption
func FormatClientGraph(graph *Graph) ClientGraph {
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
