package utils

func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
