package algorithms

func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		middle := (low + high) / 2
		guess := list[middle]

		if guess == item {
			return middle
		}

		if guess > item {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}
