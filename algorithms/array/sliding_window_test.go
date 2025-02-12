package array

import "testing"

var slidingWindowTableTest = []struct {
	in  string
	out int
}{
	{
		in:  "bcbbbcba",
		out: 4,
	},
	{
		in:  "aaaa",
		out: 2,
	},
}

func TestMaximumLengthSubstring(t *testing.T) {
	for _, tt := range slidingWindowTableTest {
		t.Run(tt.in, func(t *testing.T) {
			got := maximumLengthSubstring(tt.in)
			if got != tt.out {
				t.Errorf("expected %d, but got %d", tt.out, got)
			}
		})
	}
}
