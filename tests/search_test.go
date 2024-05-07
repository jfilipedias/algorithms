package tests

import (
	"testing"

	"github.com/jfilipedias/algorithms/algorithms/search"
)

type searchTest struct {
	input         []int
	target        int
	expected      int
	expectedError error
}

var searchTableTest = []searchTest{
	// existing value
	{
		input:         []int{1, 3, 5, 7, 9},
		target:        3,
		expected:      1,
		expectedError: nil,
	},
	// non-existing value
	{
		input:         []int{1, 3, 5, 7, 9},
		target:        13,
		expected:      -1,
		expectedError: search.ErrNotFound,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTableTest {
		got, err := search.BinarySearch(test.input, test.target)

		if test.expectedError != nil && err != test.expectedError {
			t.Errorf("expected %v, but got %v", test.expectedError, err)
		}

		if got != test.expected {
			t.Errorf("expected %d but got %d", test.expected, got)
		}
	}
}
