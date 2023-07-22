package characters

import (
	"testing"
)

func TestIsASCII(t *testing.T) {
	t.Run("ASCII string", func(t *testing.T) {
		if !isASCII("asdfgh") {
			t.Fatal("'asdfgh' contains only ascii characters")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		if !isASCII("") {
			t.Fatal()
		}
	})

	t.Run("Non-ASCII string", func(t *testing.T) {
		if isASCII("jklπoi") {
			t.Fatal()
		}
	})
}

func TestASCIICharacterIterator(t *testing.T) {
	t.Run("ASCII string", func(t *testing.T) {
		str := "qwerty"
		bytes := []byte(str)
		itr := newASCII(str)

		length := 0
		for itr.Next() {
			length++

			if itr.Str() != string(bytes[length-1]) {
				t.Fatal()
			}
		}

		if length != 6 {
			t.Fatal("'qwerty' has 6 characters")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		itr := newASCII("")

		if itr.Next() {
			t.Fatal()
		}
	})
}
