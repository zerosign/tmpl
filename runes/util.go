package runes

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
func JoinString(runes [][]rune, ch rune) string {
	return string(JoinRunes(runes, ch))
}

// Compare : compare 2 same []rune
//
// - if both []rune nil, then return true
// - if both has different length, then return false
// - if both has same length with same element for each indexes then return true
//
func Compare(lhs []rune, rhs []rune) bool {
	return (lhs == nil && rhs == nil) || ((len(lhs) == len(rhs)) && compareElements(lhs, rhs))
}

// compareElements : compare only elements for 2 []rune
//
// WARN: this function didn't check whether both runes are nil or
//       has the same length or not, if you need to compare safely use
//       runes.Compare.
//
func compareElements(lhs []rune, rhs []rune) bool {
	var result = true

	for ii := 0; ii < len(lhs); ii++ {
		result = result && lhs[ii] == rhs[ii]
		if !result {
			break
		}
	}

	return result
}

func JoinRunes(runes [][]rune, ch rune) []rune {
	var buffer = make([]rune, 0)
	for ii := 0; ii < len(runes); ii++ {
		buffer = append(buffer, runes[ii]...)

		// if we still have next
		if ii < (len(runes) - 1) {
			buffer = append(buffer, []rune{ch}...)
		}
	}

	return buffer
}
