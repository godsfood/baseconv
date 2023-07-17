package baseconv

import (
	"fmt"
	"math"
	"strings"

	"go.dkinom.dev/baseconv/internal/characters"
	"go.dkinom.dev/baseconv/options"
)

// Provides ability to convert numbers between different
// positional numeral systems
type baseConversion struct {
	fromAlphabet alphabet
	toAlphabet   alphabet

	lengthFactor    float64
	intLengthFactor int

	zeroPadding bool
}

// Initializes baseConversion struct with the given from and to alphabets
func NewBaseConversion(from string, to string, opts ...*options.BaseConversionOptions) (*baseConversion, error) {
	fromAlphabet, err := NewAlphabet(from)
	if err != nil {
		return nil, err
	}
	toAlphabet, err := NewAlphabet(to)
	if err != nil {
		return nil, err
	}

	b := NewBaseConversionAlphabet(*fromAlphabet, *toAlphabet, opts...)

	return b, nil
}

// Initializes baseConversion struct with the given from and to alphabets
func NewBaseConversionAlphabet(from alphabet, to alphabet, opts ...*options.BaseConversionOptions) *baseConversion {
	b := &baseConversion{fromAlphabet: from, toAlphabet: to}

	b.lengthFactor = math.Log(float64(from.Radix())) / math.Log(float64(to.Radix()))

	if b.lengthFactor == math.Trunc(b.lengthFactor) {
		b.intLengthFactor = int(b.lengthFactor)
	} else {
		invLengthFactor := (1 / b.lengthFactor)
		if invLengthFactor == math.Trunc(invLengthFactor) {
			b.intLengthFactor = -int(invLengthFactor)
		}
	}

	opt := options.MergeBaseConversionOptions(opts...)

	b.configure(opt)

	return b
}

func (b *baseConversion) configure(opt *options.BaseConversionOptions) {
	// Zero padding
	b.zeroPadding = false
	if opt.ZeroPadding != nil {
		b.zeroPadding = *opt.ZeroPadding
	}
}

func (b *baseConversion) options() *options.BaseConversionOptions {
	o := options.BaseConversion()

	// Zero padding
	o.SetZeroPadding(b.zeroPadding)

	return o
}

// Returns numeral representation of s in toAlphabet
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

	var changeBase func(_values []int) (string, int)
	changeBase = func(_values []int) (string, int) {
		values := []int{}
		for i, v := range _values {
			if v != 0 {
				values = append(values, _values[i:]...)
				break
			}
		}

		if len(values) == 0 {
			return "", 0
		}

		if b.intLengthFactor == 1 || b.intLengthFactor == -1 {
			sb := strings.Builder{}
			for _, value := range values {
				sb.WriteString(b.toAlphabet.chars[value])
			}

			return sb.String(), len(values)
		} else if b.intLengthFactor == 0 ||
			len(values) == 1 ||
			(b.intLengthFactor < 0 && b.intLengthFactor+len(values) <= 0) { // Horner's method
			remainder := 0
			quotients := []int{}
			for _, value := range values {
				remainder = (remainder * fromBase) + value

				quotients = append(quotients, remainder/toBase)
				remainder = remainder % toBase
			}

			r, rLen := changeBase(quotients)

			return r + b.toAlphabet.chars[remainder], rLen + 1
		} else if b.intLengthFactor > 0 {
			sb, rLen := strings.Builder{}, 0
			for i := 0; i < len(values); i++ {
				_r, _rLen := changeBase(values[i : i+1])
				if i > 0 {
					sb.WriteString(strings.Repeat(b.toAlphabet.zerochar, b.intLengthFactor-_rLen))
					sb.WriteString(_r)
					rLen += b.intLengthFactor
				} else {
					sb.WriteString(_r)
					rLen += _rLen
				}
			}

			return sb.String(), rLen
		} else { // b.intLengthFactor < 0
			sb, rLen := strings.Builder{}, 0
			i := len(values) % (-b.intLengthFactor)
			if i > 0 {
				i += b.intLengthFactor
			}
			for ; i < len(values); i -= b.intLengthFactor {
				var _values []int
				if i > 0 {
					_values = values[i : i-b.intLengthFactor]
					_r, _rLen := changeBase(_values)
					sb.WriteString(strings.Repeat(b.toAlphabet.zerochar, 1-_rLen))
					sb.WriteString(_r)
					rLen += 1
				} else {
					_values = values[:i-b.intLengthFactor]
					_r, _rLen := changeBase(_values)
					sb.WriteString(_r)
					rLen += _rLen
				}
			}

			return sb.String(), rLen
		}
	}

	values := []int{}
	c := characters.NewCharacters(ip)
	for c.Next() {
		char := c.Str()

		values = append(values, b.fromAlphabet.charset[char])
	}
	ipLen := len(values)

	r, rLen := changeBase(values)

	if b.zeroPadding {
		currentLength := rLen
		wantedLength := int(math.Ceil(float64(ipLen) * b.lengthFactor))

		if currentLength < wantedLength {
			var sb strings.Builder
			sb.WriteString(strings.Repeat(b.toAlphabet.zerochar, wantedLength-currentLength))
			sb.WriteString(r)

			r = sb.String()
		}
	} else if rLen == 0 {
		r = b.toAlphabet.zerochar
	}

	return r
}

// Returns a baseConversion which converts numerals
// from toAlphabet to fromAlphabet
func (b *baseConversion) Inverse() *baseConversion {
	return NewBaseConversionAlphabet(b.toAlphabet, b.fromAlphabet, b.options())
}
