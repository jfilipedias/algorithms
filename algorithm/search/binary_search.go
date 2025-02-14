package search

// complexity O(log n)
func binarySearch(nums []int, target, min, max int) int {
	for min <= max {
		middle := (min + max) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return -1
}
