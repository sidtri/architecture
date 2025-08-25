package main

import "fmt"

func suggestIitemIndicesForFreeDelivery(prices []int, n int) []int {
	seen := make(map[int]int)

	for i, price := range prices {

		target := n - price

		if idx, ok := seen[target]; ok {
			return []int{idx, i}
		}

		seen[price] = i
	}

	return []int{}
}

func TwoSum () {
	prices := []int{2, 30, 56, 34, 55, 10, 11, 20, 15, 60, 45, 39, 51}
	n := 99

	fmt.Println(suggestIitemIndicesForFreeDelivery(prices, n))
}
