package hashmap

// 1. Two Sum
// Given an array of integers `nums“ and an integer `target`, return
// indices of the two numbers such that they add up to `target`.
//
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
//
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	res := []int{0, 1}
	if len(nums) == 2 {
		return res
	}

	m := make(map[int]int)

	for i, n := range nums {
		if val, ok := m[target-n]; ok {
			res[0] = val
			res[1] = i
			return res
		} else {
			m[n] = i
		}
	}

	return res
}
