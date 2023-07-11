package baseconv

import (
	"testing"
)

func TestAlphabetValidations(t *testing.T) {
	t.Run("alphabet has duplicate characters", func(t *testing.T) {
		alphabet, err := NewAlphabet("0123453789")

		if alphabet != nil || err == nil {
			t.Fatalf("alphabet '%v' has duplicate characters", alphabet)
		}
	})
}

func TestAlphabetString(t *testing.T) {
	t.Run("base-10", func(t *testing.T) {
		alphabet, err := NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}
		expected := "0,1,2,3,4,5,6,7,8,9"
		if alphabet.String() != expected {
			t.Fatal()
		}
	})
}

func TestAlphabetRadix(t *testing.T) {
	t.Run("base-10", func(t *testing.T) {
		alphabet, err := NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}
		expected := 10
		if alphabet.Radix() != expected {
			t.Fatalf("radix of '%v' is %v", alphabet, expected)
		}
	})

	t.Run("base-2", func(t *testing.T) {
		alphabet, err := NewAlphabet("01")
		if err != nil {
			t.Fatal(err)
		}
		expected := 2
		if alphabet.Radix() != expected {
			t.Fatalf("radix of '%v' is %v", alphabet, expected)
		}
	})

	t.Run("base-16", func(t *testing.T) {
		alphabet, err := NewAlphabet("0123456789abcdef")
		if err != nil {
			t.Fatal(err)
		}
		expected := 16
		if alphabet.Radix() != expected {
			t.Fatalf("radix of '%v' is %v", alphabet, expected)
		}
	})

	t.Run("base-10 (emoji)", func(t *testing.T) {
		alphabet, err := NewAlphabet("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}
		expected := 10
		if alphabet.Radix() != expected {
			t.Fatalf("radix of '%v' is %v", alphabet, expected)
		}
	})
}

func TestAlphabetIsValid(t *testing.T) {
	t.Run("valid base-10 numerals", func(t *testing.T) {
		alphabet, err := NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}

		validNumerals := []string{
			"123",
		}

		for _, n := range validNumerals {
			if !alphabet.IsValid(n) {
				t.Fatalf("'%v' is valid", n)
			}
		}
	})

	t.Run("invalid base-10 numerals", func(t *testing.T) {
		alphabet, err := NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}

		invalidNumerals := []string{
			"",
			"123abc",
		}

		for _, n := range invalidNumerals {
			if alphabet.IsValid(n) {
				t.Fatalf("'%v' is invalid", n)
			}
		}
	})

	t.Run("valid base-10 (emoji) numerals", func(t *testing.T) {
		alphabet, err := NewAlphabet("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}

		validNumerals := []string{
			"2️⃣6️⃣4️⃣8️⃣",
		}

		for _, n := range validNumerals {
			if !alphabet.IsValid(n) {
				t.Fatalf("'%v' is valid", n)
			}
		}
	})

	t.Run("invalid base-10 (emoji) numerals", func(t *testing.T) {
		alphabet, err := NewAlphabet("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}

		invalidNumerals := []string{
			"",
			"2️⃣6️⃣4️⃣8️⃣🔢",
		}

		for _, n := range invalidNumerals {
			if alphabet.IsValid(n) {
				t.Fatalf("'%v' is invalid", n)
			}
		}
	})
}
