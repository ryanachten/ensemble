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
