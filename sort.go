package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func BubbleSort() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArr := bubbleSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
