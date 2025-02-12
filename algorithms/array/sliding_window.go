package array

func maximumLengthSubstring(s string) int {
	runes := []rune(s)
	l, r := 0, 0
	max := 1
	counter := make(map[rune]int)

	for r < len(runes) {
		if _, ok := counter[runes[r]]; !ok {
			counter[runes[r]] = 1
		} else {
			counter[runes[r]] += 1
		}

		for counter[runes[r]] >= 3 {
			counter[runes[l]] -= 1
			l++
		}

		if r-l+1 > max {
			max = r - l + 1
		}

		r++
	}

	return max
}
