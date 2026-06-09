package graphs

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     string
	Weight int
}

type Graph struct {
	Edges map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]Edge),
	}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Weight: weight})
}

type Item struct {
	Value    string
	Priority int
	Index    int
}

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
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (g *Graph) Dijkstra(start string) map[string]int {
	dist := make(map[string]int)
	for node := range g.Edges {
		dist[node] = math.MaxInt32
	}

	for _, edges := range g.Edges {
		for _, edge := range edges {
			if _, ok := dist[edge.To]; !ok {
				dist[edge.To] = math.MaxInt32
			}
		}
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

		for _, edge := range g.Edges[u] {
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

func DijkstraSetup() {
	g := NewGraph()

	data := map[string][]string{
		"0": {"1", "4"},
		"1": {"0", "2", "3", "4"},
		"2": {"1", "3"},
		"3": {"1", "2", "4"},
		"4": {"0", "1", "3"},
	}

	for k, neighbors := range data {
		for _, v := range neighbors {
			g.AddEdge(k, v, 1)
		}
	}
}
