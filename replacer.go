package replacer

import (
	"unicode"
)

func Replace(in, r string) string {
	var result string
	white := false
	for _, c := range in {
		if unicode.IsSpace(c) {
			if !white {
				result = result + r
			}
			white = true
		} else {
			result = result + string(c)
			white = false
		}
	}
	return result
}
