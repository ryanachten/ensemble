package models

import "sync"

// Adjacency list graph using mutex locks for concurrent read and writes
type MutexGraph struct {
	sync.RWMutex
	Vertices map[string]*Vertex
	Queue    chan func()    // stores a list of actions to execute sequentially
	Actions  sync.WaitGroup // represents tasks waiting in the queue
}

func NewMutexGraph() *MutexGraph {
	graph := &MutexGraph{
		Vertices: make(map[string]*Vertex),
		Queue:    make(chan func()),
	}

	go graph.WatchQueue() // start the watching process

	return graph
}

func (graph *MutexGraph) AddVertex(key string, data VertexData) {
	graph.Actions.Add(1)
	graph.Queue <- func() {
		_, vertexExists := graph.Vertices[key]
		if vertexExists {
			return
		}
		graph.Vertices[key] = &Vertex{Data: data, Edges: map[string]*Edge{}}
	}
}

func (graph *MutexGraph) UpdateVertexData(key string, imageUrl string) {
	graph.Actions.Add(1)
	graph.Queue <- func() {
		graph.Vertices[key].Data.ImageUrl = imageUrl
	}
}

func (graph *MutexGraph) AddEdge(srcKey, destKey, label string) {
	graph.Actions.Add(1)
	graph.Queue <- func() {
		// Ensure src and dest keys exist
		_, srcVertexExists := graph.Vertices[srcKey]
		_, destVertexExists := graph.Vertices[destKey]
		if !srcVertexExists || !destVertexExists {
			return
		}
		graph.Vertices[srcKey].Edges[destKey] = &Edge{Label: label}
	}
}

// Watches for actions added to the queue.
// When an action is added, the graph is locked, the action is executed, and then the graph is unlocked
func (graph *MutexGraph) WatchQueue() {
	for {
		action := <-graph.Queue
		graph.Lock()
		action()
		graph.Actions.Done()
		graph.Unlock()
	}
}
