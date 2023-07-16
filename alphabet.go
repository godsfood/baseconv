package baseconv

import (
	"fmt"
	"strings"

	"go.dkinom.dev/baseconv/characters"
)

type alphabet struct {
	chars    []string
	charset  map[string]int
	zerochar string
}

func NewAlphabet(str string) (*alphabet, error) {
	var chars []string

	charset := make(map[string]int)

	c := characters.NewCharacters(str)
	for c.Next() {
		char := c.Str()

		chars = append(chars, char)

		charset[char] = len(chars) - 1
	}

	if len(chars) != len(charset) {
		return nil, fmt.Errorf("Must not have duplicate characters in alphabet")
	}

	a := &alphabet{
		chars:    chars,
		charset:  charset,
		zerochar: chars[0],
	}

	return a, nil
}

func (a *alphabet) String() string {
	return strings.Join(a.chars, ",")
}

// Radix of this alphabet
func (a *alphabet) Radix() int {
	return len(a.chars)
}

// Checks if s is a valid numeral representation
func (a *alphabet) IsValid(s string) bool {
	var previousChar *string
	c := characters.NewCharacters(s)
	for c.Next() {
		char := c.Str()

		if _, present := a.charset[char]; !present {
			return false
		}

		previousChar = &char
	}
	return previousChar != nil
}

// Binary numeral system
const Binary = "01"

// Same as Binary numeral system
const Base2 = Binary

// Ternary numeral system
const Ternary = "012"

// Same as Ternary numeral system
const Base3 = Ternary

// Quaternary numeral system
const Quaternary = "0123"

// Same as Quaternary numeral system
const Base4 = Quaternary

// Quinary numeral system
const Quinary = "01234"

// Same as Quinary numeral system
const Base5 = Quinary

// Senary numeral system
const Senary = "012345"

// Same as Senary numeral system
const Base6 = Senary

// Octal numeral system
const Octal = "01234567"

// Same as Octal numeral system
const Base8 = Octal

// Decimal numeral system
const Decimal = "0123456789"

// Same as Decimal numeral system
const Base10 = Decimal

// Duodecimal numeral system
const Duodecimal = "0123456789AB"

// Same as Duodecimal numeral system
const Base12 = Duodecimal

// Hexadecimal numeral system
const Hexadecimal = "0123456789ABCDEF"

// Same as Hexadecimal numeral system
const Base16 = Hexadecimal

// Base32 numeral system
const Base32 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

// Base32 (Extended Hex) numeral system
const Base32hex = "0123456789ABCDEFGHIJKLMNOPQRSTUV"

// Base36 numeral system
const Base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Base58 numeral system
const Base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base62 numeral system
const Base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Base64 numeral system
const Base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Base64 (URL and Filename safe) numeral system
const Base64url = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
