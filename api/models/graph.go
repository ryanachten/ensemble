package models

type VertexType string

const (
	Band   VertexType = "band"
	Artist VertexType = "artist"
	Genre  VertexType = "genre"
)

type VertexData struct {
	Type       VertexType
	Url        *string
	ImageUrl   string
	IsComplete bool // Has all the necessary data for the vertex been defined. TODO: this speaks to an issue with the algorithm
}

type Vertex struct {
	Data  VertexData
	Edges map[string]*Edge
}

type Edge struct {
	Label string
}

type Graph interface {
	AddVertex(key string, data VertexData)
	UpdateVertexData(key string, imageUrl string)
	HasCompleteVertex(key string) bool
	AddEdge(srcKey, destKey, label string)
	ToClientGraph() ClientGraph
}

type ConcurrentGraph interface {
	Graph
	Wait() // Waits for graph actions to complete
}
