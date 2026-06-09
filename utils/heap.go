package utils

import (
	"fmt"
	"math"
)

type Heap struct {
	nodes []int

	len int
}

func (h *Heap) HeapifyUp(idx int) bool {
	if idx <= 0 {
		return true
	}

	current := h.nodes[idx]
	pidx := parent(idx)
	parent := h.nodes[pidx]

	if current < parent {
		h.nodes[idx], h.nodes[pidx] = h.nodes[pidx], h.nodes[idx]
		h.HeapifyUp(pidx)
	}

	return true
}

func (h *Heap) HeapifyDown(idx int) bool {
	if idx >= h.len {
		return true
	}
	left := leftChild(idx)
	right := rightChild(idx)

	if left >= h.len || right >= h.len {
		return true
	}

	current := h.nodes[idx]
	if h.len <= right {
		return true
	}

	leftChild := h.nodes[left]
	rightChild := h.nodes[right]

	if current > leftChild {
		h.nodes[current], h.nodes[leftChild] = h.nodes[leftChild], h.nodes[current]
		h.HeapifyDown(left)
	} else if current > rightChild {
		h.nodes[current], h.nodes[rightChild] = h.nodes[rightChild], h.nodes[current]
		h.HeapifyDown(right)
	}

	return true
}

func (h *Heap) Add(node int) bool {
	h.nodes = append(h.nodes, node)
	h.len = h.len + 1

	h.HeapifyUp(h.len - 1)

	return true
}

func (h *Heap) Pop() int {
	node := h.nodes[0]
	h.len = h.len - 1

	h.nodes[0] = h.nodes[h.len-1]
	h.nodes = h.nodes[0:(h.len - 1)]

	h.HeapifyDown(0)

	return node
}

func (h *Heap) Print() {
	fmt.Println(h.nodes)
}

func parent(idx int) int {
	return int(math.Floor((float64(idx) - 1.0) / 2.0))
}

func leftChild(idx int) int {
	return (2 * idx) + 1
}

func rightChild(idx int) int {
	return (2 * idx) + 2
}

func HeapCall() {
	h := Heap{nodes: make([]int, 0, 10), len: 0}
	h.Add(3)
	h.Add(1)
	h.Add(2)
	h.Add(4)
	h.Pop()

	h.Print()
}
