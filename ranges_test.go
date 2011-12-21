/*-
 * Copyright (c) 2011
 *	Benny Siegert <bsiegert@gmail.com>
 *
 * Provided that these terms and disclaimer and all copyright notices
 * are retained or reproduced in an accompanying document, permission
 * is granted to deal in this work without restriction, including un-
 * limited rights to use, publicly perform, distribute, sell, modify,
 * merge, give away, or sublicence.
 *
 * This work is provided "AS IS" and WITHOUT WARRANTY of any kind, to
 * the utmost extent permitted by applicable law, neither express nor
 * implied; without malicious intent or gross negligence. In no event
 * may a licensor, author or contributor be held liable for indirect,
 * direct, other damage, loss, or other issues arising in any way out
 * of dealing in the work, even if advised of the possibility of such
 * damage or existence of a defect, except proven that it results out
 * of said person's immediate fault when using the work as intended.
 */

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
