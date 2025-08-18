package main

import (
	"container/heap"
	"fmt"
)


type Driver struct {
	X, Y int
	D2 int // squared distance
}

type MaxHeap []Driver

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].D2 > h[j].D2 } // '>' makes it max-heap
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Driver)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func kClosestHeap(points [][]int, k int) []Driver {
	h := &MaxHeap{}
	heap.Init(h)

	for _, p := range points {
		x, y := p[0], p[1]

    d2 := (x*x) + (y*y)

		driver := Driver{X: x,Y: y, D2: d2}

		if h.Len() < k {
			heap.Push(h, driver)
		} else if d2 < (*h)[0].D2 {
			heap.Pop(h)
			heap.Push(h, driver)
		}
	}

	res := make([]Driver, h.Len())

	for i := 0; h.Len() > 0; i++ {
		res[i] = heap.Pop(h).(Driver)
	}

	return res
}


func Distance() {
	points := [][]int{
			{1, 3},   // d2=10
			{-2, 2},  // d2=8
			{5, 8},   // d2=89
			{0, 1},   // d2=1
	}
  k := 2

	closest := kClosestHeap(points, k)
	fmt.Println("Closest Drivers:")

	for _, d := range closest {
		fmt.Printf("(%d, %d) dist2=%d\n", d.X, d.Y, d.D2)
	}
}

