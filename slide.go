package main

import (
	"fmt"

	"architecture.com/m/utils"
)

func allocateSpace(processes []int, size int) int {
	var space int

	var start int
	var end int
	start = 0
	end = 0

	plength := len(processes)

	for start < plength {
		sum := utils.Sum(processes[start:(end + 1)])

		if sum == size {
			space += 1
			start += 1
			end = start
			continue
		}

		if sum < size {
			if end < plength-1 {
				end += 1
			} else {
				// end the loop
				start = end + 1
			}
			continue
		}

		if sum > size {
			start += 1
			end = start
		}
	}

	return space
}

func Slide() {
	// Driver Code
	processes := []int{1, 2, 3, 4, 5, 6, 7, 1, 23, 21, 3, 1, 2, 1, 1, 1, 1, 1, 12, 2, 3, 2, 3, 2, 2}
	fmt.Printf("Total allocate space %d", allocateSpace(processes, 1))
}
