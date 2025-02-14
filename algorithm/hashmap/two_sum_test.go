package hashmap

import (
	"reflect"
	"strconv"
	"testing"
)

var twoSumTableTest = []struct {
	in     []int
	target int
	want   []int
}{
	{
		in:     []int{2, 7, 11, 15},
		target: 9,
		want:   []int{0, 1},
	},
	{
		in:     []int{3, 2, 4},
		target: 6,
		want:   []int{1, 2},
	},
	{
		in:     []int{3, 3},
		target: 6,
		want:   []int{0, 1},
	},
	{
		in:     []int{3, 2, 3},
		target: 6,
		want:   []int{0, 2},
	},
}

func TestTwoSumChar(t *testing.T) {
	for _, tt := range twoSumTableTest {
		t.Run(strconv.Itoa(tt.target), func(t *testing.T) {
			got := twoSum(tt.in, tt.target)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
