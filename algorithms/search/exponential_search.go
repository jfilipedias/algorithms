package search

func exponentialSearch(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}

	length := len(nums)
	i := 1

	for i < length && nums[i] < target {
		i *= 2
	}

	if i >= length {
		i = length - 1
	}

	if nums[i] == target {
		return i
	}

	return binarySearch(nums[i/2:i], target)
}
