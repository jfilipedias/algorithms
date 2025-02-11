package sort

import (
	"reflect"
	"testing"
)

var quickSortTableTests = []struct {
	name string
	in   []int
	out  []int
}{
	{
		name: "unsorted slice",
		in:   []int{5, 3, 1, 4, 8, 2, 6, 10, 7, 9},
		out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name: "sorted slice",
		in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name: "reversed slice",
		in:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		out:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name: "sorted slice with negative values slice",
		in:   []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		out:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name: "unsorted slice with negative values slice",
		in:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		out:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name: "duplicated slice",
		in:   []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
		out:  []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
	},
	{
		name: "single-entry slice",
		in:   []int{1},
		out:  []int{1},
	},
	{
		name: "empty slice",
		in:   []int{},
		out:  []int{},
	},
}

func TestQuickSort(t *testing.T) {
	for _, tt := range quickSortTableTests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSort(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("out %v but got %v", tt.out, got)
			}
		})
	}
}
