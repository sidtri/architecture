package main

import "fmt"

type Queue struct {
	data []int
}

func (q *Queue) Pop() int {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Len() int {
	return len(q.data)
}

func BFSOnAdjacentMatrix(matrix [][]int, start int, target int) (res []int) {
	length := len(matrix)
	queue := Queue{}
	queue.Push(start)

	seen := make([]bool, length)
	prev := make([]int, length)
	for i := range prev {
		prev[i] = -1
	}
	matchHappen := false
	var currI int

	for queue.Len() > 0 {
		currI = queue.Pop()
		if seen[currI] {
			continue
		}
		seen[currI] = true

		if currI == target {
			matchHappen = true
			break
		}

		for i, v := range matrix[currI] {
			if v == 0 || seen[i] {
				continue
			}

			if prev[i] == -1 {
				prev[i] = currI
			}
			queue.Push(i)
		}
	}

	if matchHappen {
		data := []int{}
		e := target
		for e != -1 {
			data = append([]int{e}, data...)
			e = prev[e]
		}
		return data
	}

	return []int{}
}

func TBFSOnAdjacentMatrix() {
	adjMatrix := [][]int{
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
	}
	start := 0
	target := 3
	// expectedTraversal := []int{0, 1, 2, 3} // BFS order
	// expectedPath := []int{0, 1, 2, 3}      // Exact path from start to target

	result := BFSOnAdjacentMatrix(adjMatrix, start, target)

	fmt.Println("Reslt: ", result)

	adjMatrix = [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	start = 0
	target = 2
	// expectedTraversal := []int{0, 1, 2}
	// expectedPath := []int{0, 1, 2} // Shortest path

	result = BFSOnAdjacentMatrix(adjMatrix, start, target)

	fmt.Println("Reslt: ", result)
}
