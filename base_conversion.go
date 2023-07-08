package baseconv

import (
	"fmt"

	"github.com/rivo/uniseg"
)

type baseConversion struct {
	fromAlphabet alphabet
	toAlphabet   alphabet
}

func NewBaseConversion(from string, to string) (*baseConversion, error) {
	fromAlphabet, err := NewAlphabet(from)
	if err != nil {
		return nil, err
	}
	toAlphabet, err := NewAlphabet(to)
	if err != nil {
		return nil, err
	}

	b := baseConversion{*fromAlphabet, *toAlphabet}

	return &b, nil
}

func (b *baseConversion) Convert(s string) (r string, err error) {
	if !b.fromAlphabet.IsValid(s) {
		err = fmt.Errorf("'%v' is invalid", s)
		return
	}

	r = b.convertIntegralPart(s)
	return
}

func (b *baseConversion) convertIntegralPart(ip string) string {
	fromBase := b.fromAlphabet.Radix()
	toBase := b.toAlphabet.Radix()

	var changeBase func(_values []int) string
	changeBase = func(_values []int) string {
		values := []int{}
		for i, v := range _values {
			if v != 0 {
				values = append(values, _values[i:]...)
				break
			}
		}

		if len(values) == 0 {
			return ""
		}

		remainder := 0
		quotients := []int{}
		for _, value := range values {
			remainder = (remainder * fromBase) + value

			quotients = append(quotients, remainder/toBase)
			remainder = remainder % toBase
		}

		return changeBase(quotients) + b.toAlphabet.characters[remainder]
	}

	values := []int{}
	gr := uniseg.NewGraphemes(ip)
	for gr.Next() {
		character := gr.Str()

		values = append(values, b.fromAlphabet.characterSet[character])
	}

	r := changeBase(values)

	if len(r) == 0 {
		return b.toAlphabet.characters[0]
	}

	return r
}

func (b *baseConversion) Inverse() *baseConversion {
	return &baseConversion{b.toAlphabet, b.fromAlphabet}
}
