package search

import "testing"

var searchTableTest = []struct {
	name   string
	in     []int
	target int
	out    int
	err    error
}{
	{
		name:   "existing value",
		in:     []int{1, 3, 5, 7, 9},
		target: 3,
		out:    1,
		err:    nil,
	},
	{
		name:   "non-existing value",
		in:     []int{1, 3, 5, 7, 9},
		target: 13,
		out:    0,
		err:    ErrNotFound,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range searchTableTest {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.in, tt.target)
			if got != tt.out || err != tt.err {
				t.Errorf("expected %d and %v, but got %d and %v", tt.out, tt.err, got, err)
			}
		})
	}
}
