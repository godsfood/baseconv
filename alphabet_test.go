package baseconv_test

import (
	"fmt"
	"testing"

	"github.com/godsfood/baseconv"
)

func TestValidations(t *testing.T) {
	t.Run("alphabet has duplicate characters", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0123453789")

		if alphabet != nil || err == nil {
			t.Fatal(fmt.Sprintf("alphabet '%v' has duplicate characters", alphabet))
		}
	})
}

func TestRadix(t *testing.T) {
	t.Run("base-10", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}
		expected := 10
		if alphabet.Radix() != expected {
			t.Fatal(fmt.Sprintf("radix of '%v' is %v", alphabet, expected))
		}
	})

	t.Run("base-2", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("01")
		if err != nil {
			t.Fatal(err)
		}
		expected := 2
		if alphabet.Radix() != expected {
			t.Fatal(fmt.Sprintf("radix of '%v' is %v", alphabet, expected))
		}
	})

	t.Run("base-16", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0123456789abcdef")
		if err != nil {
			t.Fatal(err)
		}
		expected := 16
		if alphabet.Radix() != expected {
			t.Fatal(fmt.Sprintf("radix of '%v' is %v", alphabet, expected))
		}
	})

	t.Run("base-10 (emoji)", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}
		expected := 10
		if alphabet.Radix() != expected {
			t.Fatal(fmt.Sprintf("radix of '%v' is %v", alphabet, expected))
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("valid base-10 numerals", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}

		validNumerals := []string{
			"123",
		}

		for _, n := range validNumerals {
			if !alphabet.IsValid(n) {
				t.Fatal(fmt.Sprintf("'%v' is valid", n))
			}
		}
	})

	t.Run("invalid base-10 numerals", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0123456789")
		if err != nil {
			t.Fatal(err)
		}

		invalidNumerals := []string{
			"",
			"123abc",
		}

		for _, n := range invalidNumerals {
			if alphabet.IsValid(n) {
				t.Fatal(fmt.Sprintf("'%v' is invalid", n))
			}
		}
	})

	t.Run("valid base-10 (emoji) numerals", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}

		validNumerals := []string{
			"2️⃣6️⃣4️⃣8️⃣",
		}

		for _, n := range validNumerals {
			if !alphabet.IsValid(n) {
				t.Fatal(fmt.Sprintf("'%v' is valid", n))
			}
		}
	})

	t.Run("invalid base-10 (emoji) numerals", func(t *testing.T) {
		alphabet, err := baseconv.NewAlphabet("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}

		invalidNumerals := []string{
			"",
			"2️⃣6️⃣4️⃣8️⃣🔢",
		}

		for _, n := range invalidNumerals {
			if alphabet.IsValid(n) {
				t.Fatal(fmt.Sprintf("'%v' is invalid", n))
			}
		}
	})
}