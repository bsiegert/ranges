package ranges

import (
	"fmt"
	"strconv"
	"strings"
)

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
