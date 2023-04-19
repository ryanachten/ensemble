package models

type ClientNodeData struct {
	Id    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
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
	Nodes []ClientNode `json:"nodes,omitempty"`
	Edges []ClientEdge `json:"edges,omitempty"`
}

// Formats Graph for client consumption
func FormatClientGraph(graph *Graph) ClientGraph {
	var nodes []ClientNode
	var edges []ClientEdge

	for vertexKey, vertexValue := range graph.Vertices {
		nodes = append(nodes, ClientNode{
			Data: ClientNodeData{
				Id:    vertexKey,
				Label: vertexKey,
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
		Nodes: nodes,
		Edges: edges,
	}
}
