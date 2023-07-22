package characters

import (
	"testing"

	"github.com/rivo/uniseg"
)

func TestCharacterIterator(t *testing.T) {
	t.Run("ASCII string", func(t *testing.T) {
		str := "zxcvbnm"
		bytes := []byte(str)
		itr := NewCharacters(str)

		if _, ok := itr.(*ascii); !ok {
			t.Fatal("expected ascii iterator")
		}

		length := 0
		for itr.Next() {
			length++

			if itr.Str() != string(bytes[length-1]) {
				t.Fatal()
			}
		}

		if length != 7 {
			t.Fatal("'zxcvbnm' has 7 characters")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		itr := NewCharacters("")

		if _, ok := itr.(*ascii); !ok {
			t.Fatal("expected ascii iterator")
		}

		if itr.Next() {
			t.Fatal()
		}
	})

	t.Run("Non-ASCII string", func(t *testing.T) {
		str := "0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣"
		itr := NewCharacters(str)

		if _, ok := itr.(*uniseg.Graphemes); !ok {
			t.Fatal("expected graphemes iterator")
		}

		length := 0
		for itr.Next() {
			length++
		}

		if length != 10 {
			t.Fatal("'0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣' has 10 characters")
		}
	})
}
