package hashmap

import "testing"

var firstUniqueTableTest = []struct {
	in   string
	want int
}{
	{
		in:   "loveAlgorithms",
		want: 2,
	},
	{
		in:   "bcbbbcba",
		want: 7,
	},
	{
		in:   "aabb",
		want: -1,
	},
}

func TestFirstUniqChar(t *testing.T) {
	for _, tt := range firstUniqueTableTest {
		t.Run(tt.in, func(t *testing.T) {
			got := firstUniqChar(tt.in)
			if got != tt.want {
				t.Errorf("expected %d, but got %d", tt.want, got)
			}
		})
	}
}
