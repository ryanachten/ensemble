package models

import "sync"

// Adjacency list graph stored using a sync map for concurrent read and writes
type SyncGraph struct {
	Vertices *sync.Map
}

func (graph *SyncGraph) AddVertex(key string, data VertexData) {
	if graph.Vertices == nil {
		graph.Vertices = &sync.Map{}
	}
	_, vertexExists := graph.Vertices.Load(key)
	if vertexExists {
		return
	}
	graph.Vertices.Store(key, &Vertex{Data: data, Edges: map[string]*Edge{}})
}

func (graph *SyncGraph) UpdateVertexData(key string, imageUrl string) {
	vertex, vertexExists := graph.Vertices.Load(key)
	if !vertexExists {
		return
	}
	updatedVertex := vertex.(*Vertex) // sync map stores values as `any`, so we need to cast them back to vertices
	updatedVertex.Data.ImageUrl = imageUrl
	graph.Vertices.Store(key, updatedVertex)
}

func (graph *SyncGraph) AddEdge(srcKey, destKey, label string) {
	// Ensure src and dest keys exist
	srcVertex, srcVertexExists := graph.Vertices.Load(srcKey)
	_, destVertexExists := graph.Vertices.Load(destKey)
	if !srcVertexExists || !destVertexExists {
		return
	}
	updatedVertex := srcVertex.(*Vertex)
	updatedVertex.Edges[destKey] = &Edge{Label: label}
	graph.Vertices.Store(srcKey, updatedVertex)
}
