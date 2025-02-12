package search

// complexity O(log n)
func BinarySearch(arr []int, target int) (int, error) {
	min := 0
	max := len(arr) - 1
	for min <= max {
		middle := (min + max) / 2
		if arr[middle] == target {
			return middle, nil
		} else if arr[middle] > target {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return 0, ErrNotFound
}
