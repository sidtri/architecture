package main

import "fmt"

var nodeDelays = make(map[int]int)

func findNodeDelays(index int) int {
	if index == -1 {
		return 0
	}

	parents := []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	delays := []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}

	parentNodeDelay := 0

	if delay, ok := nodeDelays[parents[index]]; ok {
		parentNodeDelay = delay
	} else {
		parentNodeDelay = findNodeDelays(parents[index])
	}

	nodeDelays[index] = delays[index] + parentNodeDelay

	return nodeDelays[index]
}

func totalTime(rootID int, parents []int, delays []int) int {
	totalTime := 0

	for nodeIndex := range parents {
		delay := findNodeDelays(nodeIndex)

		if delay > totalTime {
			totalTime = delay
		}
	}

	return totalTime
}

// Imagine a binary tree where each node parent is denoted in parents list and delay it takes in delays. Identify the max time it takes to traverse.
func MaxDistanceBetweenNodes() {
	mainServerId := 0
	parents := []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	delays := []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println(totalTime(mainServerId, parents, delays))
}
