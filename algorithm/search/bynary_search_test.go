package search

import "testing"

var binarySearchTableTest = []struct {
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
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range binarySearchTableTest {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.in, tt.target, 0, len(tt.in)-1)
			if got != tt.want {
				t.Errorf("expected %d, but got %d", tt.want, got)
			}
		})
	}
}
