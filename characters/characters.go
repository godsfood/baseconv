package characters

import (
	"github.com/rivo/uniseg"
)

type characters interface {
	Next() bool

	Str() string
}

func NewCharacters(str string) characters {
	if isASCII(str) {
		return newASCII(str)
	}
	return uniseg.NewGraphemes(str)
}
