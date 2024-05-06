package search

type SearchTest struct {
	data          []int
	target        int
	expected      int
	expectedError error
}

var SearchTableTest = []SearchTest{
	{[]int{1, 3, 5, 7, 9}, 3, 1, nil},
	{[]int{1, 3, 5, 7, 9}, 13, -1, ErrNotFound},
}
