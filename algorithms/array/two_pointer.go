package array

// complexity O(n)
func reverseWords(s string) string {
	runes := []rune(s)
	length := len(runes)
	l, r := 0, 0

	for r < length {
		if r+1 < length && runes[r+1] != ' ' {
			r++
		} else {
			for i, j := l, r; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			r += 2
			l = r
		}
	}

	return string(runes)
}
