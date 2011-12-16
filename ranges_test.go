package ranges

import (
	"testing"
)

var parseTests = []struct {
	in  string
	out []int
}{
	{"1", []int{1}},
	{"1-3", []int{1, 2, 3}},
	{"6,8", []int{6, 8}},
	{"23,42-44,46-47,55", []int{23, 42, 43, 44, 46, 47, 55}},
}

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParse(t *testing.T) {
	for _, pt := range parseTests {
		actual, err := Parse(pt.in)
		if err != nil {
			t.Errorf("Parse(%q) returned error: %v", pt.in, err)
		} else if !equalIntSlice(actual, pt.out) {
			t.Errorf("Parse(%q): got %v, want %v", pt.in, actual, pt.out)
		}
	}
}
