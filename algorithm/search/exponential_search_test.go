package search

import "testing"

var exponentailSearchTableTest = []struct {
	name   string
	in     []int
	target int
	want   int
}{
	{
		name:   "existing value",
		in:     []int{1, 3, 5, 7, 9},
		target: 3,
		want:   1,
	},
	{
		name:   "non-existing value",
		in:     []int{1, 3, 5, 7, 9},
		target: 13,
		want:   -1,
	},
	{
		name:   "big slice",
		in:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		target: 32,
		want:   31,
	},
}

func TestExponentialSearch(t *testing.T) {
	for _, tt := range exponentailSearchTableTest {
		t.Run(tt.name, func(t *testing.T) {
			got := exponentialSearch(tt.in, tt.target)
			if got != tt.want {
				t.Errorf("expected %d, but got %d", tt.want, got)
			}
		})
	}

}
