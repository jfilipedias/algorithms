package hashmap

type runeInfo struct {
	pos   int
	count int
}

// 387. First Unique Character in a String
// Given a string `s`, find the first non-repeating character in it and
// return its index. If it does not exist, return `-1`.
func firstUniqChar(s string) int {
	runes := []rune(s)
	m := make(map[rune]runeInfo)

	for i, rn := range runes {
		if v, ok := m[rn]; ok {
			v.count += 1
			m[rn] = v
		} else {
			m[rn] = runeInfo{pos: i, count: 1}
		}
	}

	min := -1
	for _, v := range m {
		if v.count == 1 && (min < 0 || v.pos < min) {
			min = v.pos
		}
	}

	return min
}
