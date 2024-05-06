package tests

import (
	"testing"

	"github.com/jfilipedias/algorithms/algorithms"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		list []int
		item int
		want int
	}{
		{[]int{1, 3, 5, 7, 9}, 3, 1},
		{[]int{1, 3, 5, 7, 9}, 13, -1},
	}

	for _, test := range testCases {
		got := algorithms.BinarySearch(test.list, test.item)
		if got != test.want {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}
