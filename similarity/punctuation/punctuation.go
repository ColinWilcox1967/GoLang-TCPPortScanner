package punctuation

import (
	"strings"
	"unicode"
)

// RemovePunctuation removes all extraneous symbols from provided string
func RemovePunctuation(s string) string {
	var str string

	//Phase 1
	str = strings.TrimSpace(s) // remove leading and trailing whitespace

	// Phase 2 : Remove duplicate spaces
	foundPunct := false
	pos := 0
	str2 := ""
	for pos < len(str) {
		if unicode.IsSpace(rune(str[pos])) || unicode.IsPunct(rune(str[pos])) {
			if foundPunct {
				// do nothing
			} else {
				foundPunct = true
				str2 = str2 + string(str[pos])
			}
		} else {
			foundPunct = false
			str2 = str2 + string(str[pos])
		}
		pos++
	}
	return str2
}
