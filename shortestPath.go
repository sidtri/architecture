package main

import (
	"container/heap"
	"math"
)

// Edge represents a directed edge with a weight.
type Edge struct {
	To     int
	Weight int
}

// Item is an item in the priority queue.
type Item struct {
	Value    int // The node ID
	Priority int // The distance from source
	Index    int // The index of the item in the heap
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// ShortestPathDijkstra finds the shortest path from start to all other nodes in the graph.
// The graph is represented as an adjacency list: nodeID -> list of edges.
func ShortestPathDijkstra(graph map[int][]Edge, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Value: start, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.Value
		d := item.Priority

		if d > dist[u] {
			continue
		}

		for _, edge := range graph[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{Value: v, Priority: dist[v]})
			}
		}
	}

	return dist
}
