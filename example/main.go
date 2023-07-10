package main

import (
	"fmt"

	"go.dkinom.dev/baseconv"
)

func main() {
	decimalAndBinary()
	decimalAndDecimalEmoji()
	decimalEmojiAndHexadecimal()
}

func decimalAndBinary() {
	// Decimal alphabet - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	// Binary alphabet - 0, 1
	decimalToBinaryConverter, err := baseconv.NewBaseConversion(baseconv.Base10, baseconv.Base2)
	if err != nil {
		panic(err)
	}

	binary, err := decimalToBinaryConverter.Convert("3888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("decimal number 3888 is represented as %v in binary\n", binary) // 111100110000

	decimal, err := decimalToBinaryConverter.Inverse().Convert("111110011111")
	if err != nil {
		panic(err)
	}
	fmt.Printf("binary number 111110011111 is represented as %v in decimal\n", decimal) // 3999
}

func decimalAndDecimalEmoji() {
	// Decimal alphabet - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	// Decimal emoji alphabet - 0️⃣, 1️⃣, 2️⃣, 3️⃣, 4️⃣, 5️⃣, 6️⃣, 7️⃣, 8️⃣, 9️⃣
	decimalToDecimalEmojiConverter, err := baseconv.NewBaseConversion(baseconv.Base10, "0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
	if err != nil {
		panic(err)
	}

	decimalEmoji, err := decimalToDecimalEmojiConverter.Convert("38888")
	if err != nil {
		panic(err)
	}
	fmt.Printf("decimal number 38888 is represented as %v in decimal emoji\n", decimalEmoji) // 3️⃣8️⃣8️⃣8️⃣8️⃣

	decimal, err := decimalToDecimalEmojiConverter.Inverse().Convert("3️⃣9️⃣9️⃣9️⃣9️⃣")
	if err != nil {
		panic(err)
	}
	fmt.Printf("decimal emoji number 3️⃣9️⃣9️⃣9️⃣9️⃣ is represented as %v in decimal\n", decimal) // 39999
}

func decimalEmojiAndHexadecimal() {
	// Decimal emoji alphabet - 0️⃣, 1️⃣, 2️⃣, 3️⃣, 4️⃣, 5️⃣, 6️⃣, 7️⃣, 8️⃣, 9️⃣
	// Hexadecimal alphabet - 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A, B, C, D, E, F
	decimalEmojiToHexadecimalConverter, err := baseconv.NewBaseConversion("0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣", baseconv.Base16)
	if err != nil {
		panic(err)
	}

	hexadecimal, err := decimalEmojiToHexadecimalConverter.Convert("5️⃣1️⃣9️⃣6️⃣6️⃣")
	if err != nil {
		panic(err)
	}
	fmt.Printf("decimal emoji number 5️⃣1️⃣9️⃣6️⃣6️⃣ is represented as %v in hexadecimal\n", hexadecimal) // CAFE

	decimalEmoji, err := decimalEmojiToHexadecimalConverter.Inverse().Convert("DEADC0DE")
	if err != nil {
		panic(err)
	}
	fmt.Printf("hexadecimal number DEADC0DE is represented as %v in decimal emoji\n", decimalEmoji) // 3️⃣7️⃣3️⃣5️⃣9️⃣2️⃣9️⃣0️⃣5️⃣4️⃣
}
