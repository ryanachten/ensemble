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
	Data  VertexData       `json:"data,omitempty"`
	Edges map[string]*Edge `json:"edges,omitempty"`
}

type Edge struct {
	Label string `json:"label,omitempty"`
}

// Adjacency list graph
type Graph struct {
	Vertices map[string]*Vertex `json:"vertices,omitempty"`
}

func (graph *Graph) AddVertex(key string, data VertexData) {
	if graph.Vertices == nil {
		graph.Vertices = map[string]*Vertex{}
	}
	graph.Vertices[key] = &Vertex{Data: data, Edges: map[string]*Edge{}}
}

func (graph *Graph) AddEdge(srcKey, destKey, label string) {
	// Ensure src and dest keys exist
	_, srcVertexExists := graph.Vertices[srcKey]
	_, destVertexExists := graph.Vertices[destKey]
	if !srcVertexExists || !destVertexExists {
		return
	}

	graph.Vertices[srcKey].Edges[destKey] = &Edge{Label: label}
}
