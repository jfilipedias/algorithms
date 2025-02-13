package hashmap

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
