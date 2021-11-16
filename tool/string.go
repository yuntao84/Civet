package tool

import "strings"

func StringHasSuffix(path string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(strings.ToLower(path), strings.ToLower(suffix)) {
			return true
		}
	}
	return false
}
