package strings

import "unicode/utf8"

// IndexNotByte is similar with strings.IndexByte but returns
// the index of the first instance of character except c in s.
// or -1 if s only contains c.
func IndexNotByte(s string, c byte) int {
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] != c {
			return i
		}
	}
	return -1
}

// LastIndexByte is similar with strings.IndexByte but returns
// the index of the last instance of character except c in s,
// or -1 if s only contains c.
func LastIndexNotByte(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != c {
			return i
		}
	}
	return -1
}

type asciiSet [8]uint32

func (as *asciiSet) notContains(c byte) bool {
	return (as[c>>5] & (1 << uint(c&31))) == 0
}

func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c>>5] |= 1 << uint(c&31)
	}
	return as, true
}

// IndexNotAny is similar with strings.IndexAny but returns
// the index of the first instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
func IndexNotAny(s, chars string) int {
	if len(chars) > 0 {
		if len(s) > 8 {
			if as, isASCII := makeASCIISet(chars); isASCII {
				for i := 0; i < len(s); i++ {
					if as.notContains(s[i]) {
						return i
					}
				}
				return -1
			}
		}

	LabelFirstLoop:
		for i, c := range s {
			for j, m := range chars {
				if c != m && j == len(chars)-1 {
					return i
				} else if c != m {
					continue
				} else {
					continue LabelFirstLoop
				}
			}
		}
	}
	return -1
}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexNotAny(s, chars string) int {
	if len(chars) > 0 {
		if len(s) > 8 {
			if as, isASCII := makeASCIISet(chars); isASCII {
				for i := len(s) - 1; i >= 0; i-- {
					if as.notContains(s[i]) {
						return i
					}
				}
				return -1
			}
		}
	LabelFirstLoop:
		for i := len(s); i > 0; {
			r, size := utf8.DecodeLastRuneInString(s[:i])
			i -= size
			for j, m := range chars {
				if r != m && j == len(chars)-1 {
					return i
				} else if r != m {
					continue
				} else {
					continue LabelFirstLoop
				}
			}
		}
	}
	return -1
}
