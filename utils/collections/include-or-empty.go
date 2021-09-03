package collections

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
