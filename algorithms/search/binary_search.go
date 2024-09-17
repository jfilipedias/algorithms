package search

// complexity O(log(n))
func BinarySearch(arr []int, target int) (int, error) {
	min := 0
	max := len(arr) - 1
	for min <= max {
		middle := (min + max) / 2
		guess := arr[middle]
		if guess == target {
			return middle, nil
		} else if guess > target {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return -1, ErrNotFound
}
