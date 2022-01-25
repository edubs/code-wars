package highlow

import (
	"testing"
)

type test struct {
	list string
	want string
}

func TestHighLow(t *testing.T) {
	var tests = []test{
		{"1 2 3 4 5", "5 1"},
		{"1 1", "1 1"},
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
		{"42", "42 42"},
		{"-1 -1", "-1 -1"},
	}

	for _, tt := range tests {
		testname := tt.list

		t.Run(testname, func(t *testing.T) {
			ans := HighAndLow(tt.list)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
