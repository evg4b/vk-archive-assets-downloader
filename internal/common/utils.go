package common

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

func IncludeOrEmpty(a string, list []string) bool {
	if list == nil || len(list) == 0 {
		return true
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
