package util

// HasPrefix: A helper like strings.HasPrefix (as look ahead pointer).
//
// returns true if source has prefix as prefix
//         false otherwise
//
func HasPrefix(source []rune, prefix []rune) bool {
	size := len(prefix)
	flag := true

	// since the source is smaller
	if len(source) < size {
		return false
	}

	for ii := 0; ii < size; ii++ {
		flag = flag && (source[ii] == prefix[ii])
	}

	return flag
}

// AnyPrefixes: A helper method that check whether given string ([]rune)
//              have a prefix inside prefixes
//
// return true if any of the prefixes existed in source string
//        false otherwise
//
func AnyPrefixes(source []rune, prefixes [][]rune) bool {
	flag := false

	for ii := 0; ii < len(prefixes); ii++ {
		flag = flag || HasPrefix(source, prefixes[ii])
	}

	return flag
}

// Join; join multiple utf8 string ([]rune) into 1 string
//
func Join(runes [][]rune, ch rune) string {
	var buffer []rune = make([]rune, 0)

	for ii := 0; ii < len(runes); ii++ {
		buffer = append(buffer, runes[ii]...)

		// if we still have next
		if ii < (len(runes) - 1) {
			buffer = append(buffer, []rune{ch}...)
		}
	}

	return string(buffer)
}
