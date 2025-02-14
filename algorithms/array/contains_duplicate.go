package array

// Given an integer array `nums“ and an integer `k“, return `true“ if there are two distinct
// indices `i“ and `j“ in the array such that `nums[i] == nums[j]` and `abs(i - j) <= k`.
func containsNearbyDuplicate(nums []int, k int) bool {
	l, r := 0, 0
	m := make(map[int]int)

	for r < len(nums) {
		if _, ok := m[nums[r]]; !ok {
			m[nums[r]] = 1
		} else {
			m[nums[r]] += 1
		}

		for m[nums[r]] >= 2 {
			if (l >= r && l-r <= k) || (l < r && r-l <= k) {
				return true
			} else {
				m[nums[l]] -= 1
				l += 1
			}
		}

		r += 1
	}

	return false
}
