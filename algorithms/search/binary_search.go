package search

// complexity O(log n)
func binarySearch(arr []int, target int) int {
	min := 0
	max := len(arr) - 1
	for min <= max {
		middle := (min + max) / 2
		if arr[middle] == target {
			return middle
		} else if arr[middle] > target {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return -1
}
