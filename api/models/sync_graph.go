package models

import (
	"sync"
)

// Vertex with edges stored using a sync map for concurrent read and writes
// While this was performant, it seemed to result in data loss. Use the more reliable `SyncGraph`
type SyncVertex struct {
	Data  VertexData
	Edges sync.Map // sync map[string]*Edge
}

// Adjacency list graph stored using a sync map for concurrent read and writes
type SyncGraph struct {
	Vertices sync.Map       // sync map[string]*Vertex
	Queue    chan func()    // stores a list of actions to execute sequentially
	Actions  sync.WaitGroup // represents tasks waiting in the queue
}

func NewSyncGraph() *SyncGraph {
	graph := &SyncGraph{
		Queue: make(chan func()),
	}

	go graph.WatchQueue() // start the watching process

	return graph
}

func (graph *SyncGraph) AddVertex(key string, data VertexData) {
	graph.Actions.Add(1)
	graph.Queue <- func() {
		_, vertexExists := graph.Vertices.Load(key)
		if vertexExists {
			return
		}
		graph.Vertices.Store(key, SyncVertex{Data: data})
	}
}

func (graph *SyncGraph) UpdateVertexData(key string, imageUrl string) {
	graph.Actions.Add(1)
	graph.Queue <- func() {
		vertex, vertexExists := graph.Vertices.Load(key)
		if !vertexExists {
			return
		}
		updatedVertex := vertex.(SyncVertex) // sync map stores values as `any`, so we need to cast them back to vertices
		updatedVertex.Data.ImageUrl = imageUrl
		graph.Vertices.Store(key, updatedVertex)
	}
}

func (graph *SyncGraph) AddEdge(srcKey, destKey, label string) {
	graph.Actions.Add(1)
	graph.Queue <- func() {
		// Ensure src and dest keys exist
		srcVertex, srcVertexExists := graph.Vertices.Load(srcKey)
		_, destVertexExists := graph.Vertices.Load(destKey)
		if !srcVertexExists || !destVertexExists {
			return
		}
		updatedVertex := srcVertex.(SyncVertex)
		updatedVertex.Edges.Store(destKey, Edge{Label: label})
		graph.Vertices.Store(srcKey, updatedVertex)
	}

}

// Watches for actions added to the queue.
func (graph *SyncGraph) WatchQueue() {
	for {
		action := <-graph.Queue
		action()
		graph.Actions.Done()
	}
}
