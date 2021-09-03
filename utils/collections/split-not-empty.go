package collections

import "strings"

func SplitNotEmpty(src string) []string {
	dest := []string{}

	for _, v := range strings.Split(src, ",") {
		if len(v) > 0 {
			dest = append(dest, v)
		}
	}

	return dest
}
