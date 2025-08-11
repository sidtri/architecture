package main

// Identify the friends circles
//

// func identifyFriends(data [][]bool) int {
//	for _, d := range data: {
	
// 	}
// }

// d := [[true, false, false, false], [false, true, false, false], [false, false, true, false], [false,false, false, true]]
// identifyFriends(d)
//
import "fmt"


func dfs (matrix [][]bool, visited []bool, i int) {
	for j := 0; j < len(matrix); j++ {
		if matrix[i][j] && !visited[j] {
			visited[j] = true
			dfs(matrix, visited, j)
		}
	}
}

func findFriendCircles (matrix [][]bool) int {
	n := len(matrix)

	wentThrough := make([]bool, n)
	circles := 0


	for i := 0; i < n; i++ {
		if !wentThrough[i] {
			wentThrough[i] = true
			dfs(matrix, wentThrough, i)
			circles ++
		}
	}	

	return circles
}

func Friends () {

	// Friend circles calling
	fmt.Println("Find friend circles")

	dataMatrix := [][]bool{{true, true, false, false}, {true, true, true, false}, {false, true, true, false}, {false, false, false, true}}

	result := findFriendCircles(dataMatrix)
	// total number of groups
	fmt.Println("Total number of groups: ", result)
}
