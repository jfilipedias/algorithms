package search

func BinarySearch(arr []int, target int) (int, error) {
	low := 0
	high := len(arr) - 1
	for low <= high {
		middle := (low + high) / 2
		guess := arr[middle]
		if guess == target {
			return middle, nil
		} else if guess > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1, ErrNotFound
}
