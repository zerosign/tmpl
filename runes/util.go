package runes

//
// A helper like strings.HasPrefix (as look ahead pointer).
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

	for ii := 0; ii < size; ii += 1 {
		flag = flag && (source[ii] == prefix[ii])
	}

	return flag
}

func HasAllPrefixes(source []rune, prefixes [][]rune) bool {
	flag := true

	for ii := 0; ii < len(prefixes); ii += 1 {
		flag = flag && HasPrefix(source, prefixes[ii])
	}

	return flag
}

func Join(runes [][]rune, ch rune) string {
	var buffer []rune = make([]rune, 0)

	for ii := 0; ii < len(runes); ii += 1 {
		buffer = append(buffer, runes[ii]...)

		// if we still have next
		if ii < (len(runes) - 1) {
			buffer = append(buffer, []rune{' ', ch}...)
		}
	}

	return string(buffer)
}
