package models

type VertexType string

const (
	Band   VertexType = "Band"
	Artist VertexType = "Artist"
)

type VertexData struct {
	Type VertexType `json:"type,omitempty"`
	Url  *string    `json:"url,omitempty"`
}

type Vertex struct {
	Val   VertexData       `json:"val,omitempty"`
	Edges map[string]*Edge `json:"edges,omitempty"`
}

type Edge struct {
	Label  string  `json:"label,omitempty"`
	Vertex *Vertex `json:"vertex,omitempty"` // TODO: pretty sure we just need to store the destination key here, not the entire vertex since this leads to a lot of redundant data
}

// Adjacency list graph
type Graph struct {
	Vertices map[string]*Vertex `json:"vertices,omitempty"`
}

func (graph *Graph) AddVertex(key string, val VertexData) {
	if graph.Vertices == nil {
		graph.Vertices = map[string]*Vertex{}
	}
	graph.Vertices[key] = &Vertex{Val: val, Edges: map[string]*Edge{}}
}

func (graph *Graph) AddEdge(srcKey, destKey, label string) {
	// Ensure src and dest keys exist
	_, srcVertexExists := graph.Vertices[srcKey]
	_, destVertexExists := graph.Vertices[destKey]
	if !srcVertexExists || !destVertexExists {
		return
	}

	graph.Vertices[srcKey].Edges[destKey] = &Edge{Label: label, Vertex: graph.Vertices[destKey]}
}
