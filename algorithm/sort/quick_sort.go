package sort

import "slices"

// complexity O(n*log(n))
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	lte := make([]int, 0)
	gt := make([]int, 0)
	for _, i := range arr[1:] {
		if i <= pivot {
			lte = append(lte, i)
		} else {
			gt = append(gt, i)
		}
	}
	return slices.Concat(QuickSort(lte), []int{pivot}, QuickSort(gt))
}
