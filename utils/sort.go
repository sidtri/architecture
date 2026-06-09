package utils

func qs(data []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(data, lo, hi)

	qs(data, lo, pivotIdx-1)
	qs(data, pivotIdx+1, hi)
}

func partition(data []int, lo, hi int) int {
	pivot := data[hi]
	idx := lo - 1

	for j := lo; j < hi; j++ {
		if data[j] <= pivot {
			idx++
			data[idx], data[j] = data[j], data[idx]
		}
	}
	idx++
	data[idx], data[hi] = data[hi], data[idx]
	return idx
}

func QuickSort(data []int) {
	if len(data) < 2 {
		return
	}
	qs(data, 0, len(data)-1)
}
