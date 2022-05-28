package utils

import "strings"

func MultipleContains(text string, contain ...string) bool {
	for _, v := range contain {
		if !(strings.Contains(text, v)) {
			return false
		}
	}

	return true
}
