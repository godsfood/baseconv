package characters

import (
	"unicode"
)

func isASCII(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

type ascii struct {
	// The original ascii string.
	str string

	// The current ascii character.
	char string

	// The byte offset of the current character relative to the original string.
	offset int
}

// newASCII returns a new ascii character iterator.
func newASCII(str string) *ascii {
	return &ascii{
		str:    str,
		char:   "",
		offset: -1,
	}
}

func (a *ascii) Next() bool {
	if a.offset >= len(a.str)-1 {
		a.char = ""
		return false
	}
	a.offset += 1
	a.char = string(a.str[a.offset])
	return true
}

func (a *ascii) Str() string {
	return a.char
}
