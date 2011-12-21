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

// Package ranges contains tools for working with integer ranges.
// 
// An "integer range" allows to give a set of numbers as a string,
// which can be parsed by a call to Parse. The result can be obtained
// as a slice of integers by calling Expand or be tested against with
// Contains.
package ranges

import (
	"fmt"
	"strconv"
	"strings"
)

// An IntRange is a single component of an integer range expression.
type IntRange struct {
	Lo, Hi int
}

// IntRanges is a slice of multiple integer ranges, allowing the
// expression of non-contiguous ranges (for example "1,3-4").
type IntRanges []IntRange

// Contains returns true if ir contains value.
func (ir IntRange) Contains(value int) bool {
	return value >= ir.Lo && value <= ir.Hi
}

// Expand returns a slice of integers that contains all the numbers in ir.
func (ir IntRange) Expand() []int {
	e := make([]int, 0, ir.Hi-ir.Lo+1)
	for i := ir.Lo; i <= ir.Hi; i++ {
		e = append(e, i)
	}
	return e
}

func Parse(r string) ([]int, error) {
	var expanded []int

	for _, item := range strings.Split(r, ",") {
		lohi := strings.Split(item, "-")
		switch len(lohi) {
		case 1:
			v, err := strconv.Atoi(item)
			if err != nil {
				return nil, err
			}
			expanded = append(expanded, v)
		case 2:
			lo, err := strconv.Atoi(lohi[0])
			if err != nil {
				return nil, err
			}
			hi, err := strconv.Atoi(lohi[1])
			if err != nil {
				return nil, err
			}
			for i := lo; i <= hi; i++ {
				expanded = append(expanded, i)
			}
		default:
			return nil, fmt.Errorf("invalid range: %s", item)
		}
	}
	return expanded, nil
}
