package array

import "testing"

var reverseStringTableTest = []struct {
	in   string
	want string
}{
	{
		in:   "Let's take this contest",
		want: "s'teL ekat siht tsetnoc",
	},
	{
		in:   "Mr Ding",
		want: "rM gniD",
	},
}

func TestReverseWords(t *testing.T) {
	for _, tt := range reverseStringTableTest {
		t.Run(tt.in, func(t *testing.T) {
			got := reverseWords(tt.in)
			if got != tt.want {
				t.Errorf("expected '%s', but got '%s'", tt.want, got)
			}
		})
	}
}
