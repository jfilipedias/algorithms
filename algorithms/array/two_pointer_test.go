package array

import "testing"

var reverseStringTableTest = []struct {
	in  string
	out string
}{
	{
		in:  "Let's take this contest",
		out: "s'teL ekat siht tsetnoc",
	},
	{
		in:  "Mr Ding",
		out: "rM gniD",
	},
}

func TestReverseWords(t *testing.T) {
	for _, tt := range reverseStringTableTest {
		t.Run(tt.in, func(t *testing.T) {
			got := reverseWords(tt.in)
			if got != tt.out {
				t.Errorf("expected '%s', but got '%s'", tt.out, got)
			}
		})
	}
}
