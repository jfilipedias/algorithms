package search

func BinarySearch(list []int, target int) (int, error) {
	low := 0
	high := len(list) - 1

	for low <= high {
		middle := (low + high) / 2
		guess := list[middle]

		if guess == target {
			return middle, nil
		}

		if guess > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1, ErrNotFound
}
