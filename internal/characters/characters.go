package characters

import (
	"github.com/rivo/uniseg"
)

type characters interface {
	Next() bool

	Str() string
}

// NewCharacters returns a new character iterator.
func NewCharacters(str string) characters {
	if isASCII(str) {
		return newASCII(str)
	}
	return uniseg.NewGraphemes(str)
}
