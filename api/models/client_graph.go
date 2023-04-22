package models

type ClientNodeData struct {
	Id       string     `json:"id"`
	Label    string     `json:"label"`
	Type     VertexType `json:"type"`
	ImageUrl string     `json:"imageUrl,omitempty"`
}

type ClientNode struct {
	Data ClientNodeData `json:"data"`
}

type ClientEdgeData struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label"`
}

type ClientEdge struct {
	Data ClientEdgeData `json:"data"`
}

// Graph DTO formatted for client consumption
type ClientGraph struct {
	NodeCount int          `json:"nodeCount,omitempty"`
	EdgeCount int          `json:"edgeCount,omitempty"`
	Nodes     []ClientNode `json:"nodes,omitempty"`
	Edges     []ClientEdge `json:"edges,omitempty"`
}

// Formats Graph for client consumption
func FormatClientGraph(graph *SyncGraph) ClientGraph {
	var nodes []ClientNode
	var edges []ClientEdge

	graph.Vertices.Range(func(outerKey, outerValue any) bool {
		vertexKey := outerKey.(string)
		vertexValue := outerValue.(*SyncVertex)
		nodes = append(nodes, ClientNode{
			Data: ClientNodeData{
				Id:       vertexKey,
				Label:    vertexKey,
				Type:     vertexValue.Data.Type,
				ImageUrl: vertexValue.Data.ImageUrl,
			},
		})
		vertexValue.Edges.Range(func(innerKey, innerValue any) bool {
			edgeKey := innerKey.(string)
			edgeValue := innerValue.(*Edge)
			edges = append(edges, ClientEdge{
				Data: ClientEdgeData{
					Source: vertexKey,
					Target: edgeKey,
					Label:  edgeValue.Label,
				},
			})
			return true
		})
		return true
	})
	return ClientGraph{
		NodeCount: len(nodes),
		EdgeCount: len(edges),
		Nodes:     nodes,
		Edges:     edges,
	}
}
