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

	go graph.watchQueue() // start the watching process

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
		updatedVertex.Data.IsComplete = true
		graph.Vertices.Store(key, updatedVertex)
	}
}

func (graph *SyncGraph) HasCompleteVertex(key string) bool {
	vertex, vertexExists := graph.Vertices.Load(key)
	if vertexExists {
		vertexData := vertex.(SyncVertex).Data
		return vertexData.IsComplete
	}
	return false
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

func (graph *SyncGraph) Wait() {
	graph.Actions.Wait()
}

// Formats SyncGraph for client consumption
func (graph *SyncGraph) ToClientGraph() ClientGraph {
	var nodes []ClientNode
	var edges []ClientEdge

	graph.Vertices.Range(func(outerKey, outerValue any) bool {
		vertexKey := outerKey.(string)
		vertexValue := outerValue.(SyncVertex)
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
			edgeValue := innerValue.(Edge)
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

// Watches for actions added to the queue.
func (graph *SyncGraph) watchQueue() {
	for {
		action := <-graph.Queue
		action()
		graph.Actions.Done()
	}
}
