package models

type VertexType string

const (
	Band   VertexType = "band"
	Artist VertexType = "artist"
)

type VertexData struct {
	Type     VertexType
	Url      *string
	ImageUrl string
}

type Vertex struct {
	Data  VertexData
	Edges map[string]*Edge
}

type Edge struct {
	Label string
}

// Adjacency list graph. Does not cater for concurrent read and writes
type Graph struct {
	Vertices map[string]*Vertex
}

func (graph *Graph) AddVertex(key string, data VertexData) {
	if graph.Vertices == nil {
		graph.Vertices = map[string]*Vertex{}
	}
	_, vertexExists := graph.Vertices[key]
	if vertexExists {
		return
	}
	graph.Vertices[key] = &Vertex{Data: data, Edges: map[string]*Edge{}}
}

func (graph *Graph) UpdateVertexData(key string, imageUrl string) {
	graph.Vertices[key].Data.ImageUrl = imageUrl
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
