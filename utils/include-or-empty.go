package utils

import "strings"

func IncludeOrEmpty(a string, list []string) bool {
	if len(list) == 0 {
		return true
	}

	for _, b := range list {
		if strings.EqualFold(a, b) {
			return true
		}
	}

	return false
}
