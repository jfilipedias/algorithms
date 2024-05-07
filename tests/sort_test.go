package tests

import (
	"reflect"
	"testing"

	"github.com/jfilipedias/algorithms/algorithms/sort"
)

type sortTest struct {
	input    []int
	expected []int
}

var sortTableTests = []sortTest{
	// unsorted slice
	{
		input:    []int{5, 3, 1, 4, 8, 2, 6, 10, 7, 9},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	// sorted slice
	{
		input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	// reversed slice
	{
		input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	// sorted slice with negative values
	{
		input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	// unsorted slice with negative values
	{
		input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	// duplicated values
	{
		input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
		expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
	},
	// single-entry slice
	{
		input:    []int{1},
		expected: []int{1},
	},
	// empty slice
	{
		input:    []int{},
		expected: []int{},
	},
}

func TestSelectionSort(t *testing.T) {
	for _, test := range sortTableTests {
		got := sort.SelectionSort(test.input)
		isSorted := reflect.DeepEqual(got, test.expected)

		if !isSorted {
			t.Errorf("expected %v but got %v", test.expected, got)
		}
	}
}
