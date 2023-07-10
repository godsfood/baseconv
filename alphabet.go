package baseconv

import (
	"fmt"
	"strings"

	"github.com/rivo/uniseg"
)

type alphabet struct {
	characters   []string
	characterSet map[string]int
}

func NewAlphabet(str string) (*alphabet, error) {
	var characters []string

	characterSet := make(map[string]int)

	gr := uniseg.NewGraphemes(str)
	for gr.Next() {
		character := gr.Str()

		characters = append(characters, character)

		characterSet[character] = len(characters) - 1
	}

	if len(characters) != len(characterSet) {
		return nil, fmt.Errorf("Must not have duplicate characters in alphabet")
	}

	a := alphabet{characters, characterSet}

	return &a, nil
}

func (a *alphabet) String() string {
	return strings.Join(a.characters, ",")
}

func (a *alphabet) Radix() int {
	return len(a.characters)
}

func (a *alphabet) IsValid(s string) bool {
	var previousCharacter *string
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		character := gr.Str()

		if _, present := a.characterSet[character]; !present {
			return false
		}

		previousCharacter = &character
	}
	return previousCharacter != nil
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

// Base64 numeral system
const Base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Base64 (URL and Filename safe) numeral system
const Base64url = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
