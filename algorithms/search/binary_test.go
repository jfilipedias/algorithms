package search

import "testing"

func TestBinarySearch(t *testing.T) {

	for _, test := range SearchTableTest {
		got, err := BinarySearch(test.data, test.target)

		if test.expectedError != nil && err != test.expectedError {
			t.Errorf("expected %v, but got %v", test.expectedError, err)
		}

		if got != test.expected {
			t.Errorf("expected %d but go %d", test.expected, got)
		}
	}
}
