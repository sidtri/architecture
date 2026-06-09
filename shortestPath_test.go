package main

import (
	"reflect"
	"testing"
)

func TestShortestPathDijkstra(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[int][]Edge
		start    int
		expected map[int]int
	}{
		{
			name: "Simple Graph",
			graph: map[int][]Edge{
				0: {{To: 1, Weight: 4}, {To: 2, Weight: 1}},
				1: {{To: 3, Weight: 1}},
				2: {{To: 1, Weight: 2}, {To: 3, Weight: 5}},
				3: {},
			},
			start: 0,
			expected: map[int]int{
				0: 0,
				1: 3, // 0 -> 2 -> 1
				2: 1,
				3: 4, // 0 -> 2 -> 1 -> 3
			},
		},
		{
			name: "Disconnected Graph",
			graph: map[int][]Edge{
				0: {{To: 1, Weight: 2}},
				1: {},
				2: {{To: 3, Weight: 1}},
				3: {},
			},
			start: 0,
			expected: map[int]int{
				0: 0,
				1: 2,
				2: 2147483647, // MaxInt32
				3: 2147483647, // MaxInt32
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShortestPathDijkstra(tt.graph, tt.start)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ShortestPathDijkstra() = %v, want %v", got, tt.expected)
			}
		})
	}
}
