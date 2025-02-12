package array

import "testing"

var slidingWindowTableTest = []struct {
	in   string
	want int
}{
	{
		in:   "bcbbbcba",
		want: 4,
	},
	{
		in:   "aaaa",
		want: 2,
	},
}

func TestMaximumLengthSubstring(t *testing.T) {
	for _, tt := range slidingWindowTableTest {
		t.Run(tt.in, func(t *testing.T) {
			got := maximumLengthSubstring(tt.in)
			if got != tt.want {
				t.Errorf("expected %d, but got %d", tt.want, got)
			}
		})
	}
}
