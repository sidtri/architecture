package main

import (
	"fmt"
	"sort"
)

func Space () {
	// Sttement: Identify the minimum number of spaces needed for below meetings. so no one has to wait.
	meetings := [][]int{
    {2, 8},
    {3, 4},
    {3, 9},
    {5, 11},
    {8, 20},
    {11, 15},
	}

	starts := make([]int, len(meetings))
	ends := make([]int, len(meetings))

	for i, meet := range meetings {
		starts[i] = meet[0]
		ends[i] = meet[1]
	}


	sort.Ints(starts)
	sort.Ints(ends)

	spaces := 0
	maxSpaces := 0
	i, j := 0, 0

	for i < len(starts) {
		if starts[i] < ends[j] {
			spaces++

			if spaces > maxSpaces {
				maxSpaces = spaces
			}

			i++
		} else {
			// Meeting ended if end is less than start
			spaces--
			j++
		}

	}

	fmt.Printf("max spaces needed %v", maxSpaces)
}
