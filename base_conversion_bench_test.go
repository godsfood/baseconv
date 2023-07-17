package baseconv

import (
	"testing"
)

func BenchmarkASCIIConverter(bm *testing.B) {
	b, err := NewBaseConversion("0123456789", Hexadecimal)
	if err != nil {
		bm.Fatal(err)
	}

	s := "33204"

	bm.ResetTimer()
	for i := 0; i < bm.N; i++ { // 1366 ns/op
		_, err := b.Convert(s)
		if err != nil {
			bm.Fatal(err)
		}
	}
}

func BenchmarkNonASCIIConverter(bm *testing.B) {
	b, err := NewBaseConversion("౦123456789", Hexadecimal)
	if err != nil {
		bm.Fatal(err)
	}

	s := "332౦4"

	bm.ResetTimer()
	for i := 0; i < bm.N; i++ { // 4104 ns/op
		_, err := b.Convert(s)
		if err != nil {
			bm.Fatal(err)
		}
	}
}

func BenchmarkIntegerLengthFactorConverter(bm *testing.B) {
	bm.Run("Base64 -> Base8", func(bm *testing.B) {
		b, err := NewBaseConversion(Base64, Base8)
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 1617 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})

	bm.Run("Base10 -> Base10 (reverted)", func(bm *testing.B) {
		b, err := NewBaseConversion(Base10, "9876543210")
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 811.8 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})

	bm.Run("Base6 -> Base36", func(bm *testing.B) {
		b, err := NewBaseConversion(Base6, Base36)
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 1081 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})
}

func BenchmarkNonIntegerLengthFactorConverter(bm *testing.B) {
	bm.Run("Base64 -> Base6", func(bm *testing.B) {
		b, err := NewBaseConversion(Base64, Base6)
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 2564 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})

	bm.Run("Base64 -> Base10", func(bm *testing.B) {
		b, err := NewBaseConversion(Base64, Base10)
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 2145 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})

	bm.Run("Base5 -> Base36", func(bm *testing.B) {
		b, err := NewBaseConversion("01234", Base36)
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 1239 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})

	bm.Run("Base7 -> Base36", func(bm *testing.B) {
		b, err := NewBaseConversion("0123456", Base36)
		if err != nil {
			bm.Fatal(err)
		}

		s := "33204"

		bm.ResetTimer()
		for i := 0; i < bm.N; i++ { // 1242 ns/op
			_, err := b.Convert(s)
			if err != nil {
				bm.Fatal(err)
			}
		}
	})
}
