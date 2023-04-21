package models

import "sync"

// Adjacency list graph using mutex locks for concurrent read and writes
// While this was performant, it seemed to result in data loss. Use the more reliable `SyncGraph`
type MutexGraph struct {
	sync.RWMutex
	Vertices map[string]*Vertex
}

func (graph *MutexGraph) AddVertex(key string, data VertexData) {
	if graph.Vertices == nil {
		graph.Vertices = map[string]*Vertex{}
	}
	graph.Lock()
	_, vertexExists := graph.Vertices[key]
	graph.Unlock()
	if vertexExists {
		return
	}
	graph.Lock()
	graph.Vertices[key] = &Vertex{Data: data, Edges: map[string]*Edge{}}
	graph.Unlock()
}

func (graph *MutexGraph) UpdateVertexData(key string, imageUrl string) {
	graph.Lock()
	graph.Vertices[key].Data.ImageUrl = imageUrl
	graph.Unlock()
}

func (graph *MutexGraph) AddEdge(srcKey, destKey, label string) {
	// Ensure src and dest keys exist
	graph.Lock()
	_, srcVertexExists := graph.Vertices[srcKey]
	_, destVertexExists := graph.Vertices[destKey]
	graph.Unlock()
	if !srcVertexExists || !destVertexExists {
		return
	}

	graph.Lock()
	graph.Vertices[srcKey].Edges[destKey] = &Edge{Label: label}
	graph.Unlock()
}
