package models

import "strings"

type GraphStrategy string

const (
	Sequential GraphStrategy = "sequential"
	SyncMap    GraphStrategy = "syncmap"
	Mutex      GraphStrategy = "mutex"
)

var (
	strategyMap = map[string]GraphStrategy{
		"sequential": Sequential,
		"syncmap":    SyncMap,
		"mutex":      Mutex,
	}
)

const DEFAULT_GRAPH_STRATEGY = Mutex

func ParseStrategyString(str string) GraphStrategy {
	strategy, ok := strategyMap[strings.ToLower(str)]
	if !ok {
		return DEFAULT_GRAPH_STRATEGY
	}
	return strategy
}
