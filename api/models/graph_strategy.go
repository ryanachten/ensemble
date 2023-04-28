package models

import "strings"

type GraphStrategy string

const (
	InSync  GraphStrategy = "insync"
	SyncMap GraphStrategy = "sync"
	Mutex   GraphStrategy = "mutex"
)

var (
	strategyMap = map[string]GraphStrategy{
		"insync": InSync,
		"sync":   SyncMap,
		"mutex":  Mutex,
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
