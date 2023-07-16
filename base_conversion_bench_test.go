package baseconv

import (
	"testing"
)

func BenchmarkASCIIConverter(bm *testing.B) {
	b, err := NewBaseConversion("0123456789", Hexadecimal)
	if err != nil {
		bm.Fatal(err)
	}

	s := "3735929054"

	bm.ResetTimer()
	for i := 0; i < bm.N; i++ { // 3061 ns/op
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

	s := "3735929౦54"

	bm.ResetTimer()
	for i := 0; i < bm.N; i++ { // 8284 ns/op
		_, err := b.Convert(s)
		if err != nil {
			bm.Fatal(err)
		}
	}
}
