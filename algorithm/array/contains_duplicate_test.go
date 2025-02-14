package array

import (
	"strconv"
	"testing"
)

var containsDuplicateTableTest = []struct {
	in     []int
	target int
	want   bool
}{
	{
		in:     []int{1, 2, 3, 1},
		target: 3,
		want:   true,
	},
	{
		in:     []int{1, 0, 1, 1},
		target: 1,
		want:   true,
	},
	{
		in:     []int{1, 2, 3, 1, 2, 3},
		target: 2,
		want:   false,
	},
}

func TestContainsNearbyDuplicate(t *testing.T) {
	for i, tt := range containsDuplicateTableTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := containsNearbyDuplicate(tt.in, tt.target)
			if got != tt.want {
				t.Errorf("want %t but got %t", tt.want, got)
			}
		})
	}
}
