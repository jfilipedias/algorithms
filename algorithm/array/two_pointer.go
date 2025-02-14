package array

// 557. Reverse Words in a String III
// Given a string `s`, reverse the order of characters in each word
// within a sentence while still preserving whitespace and initial
// word order.
func reverseWords(s string) string {
	runes := []rune(s)
	length := len(runes)
	l, r := 0, 0

	for r < length {
		if r+1 < length && runes[r+1] != ' ' {
			r += 1
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
