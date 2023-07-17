package baseconv

import (
	"fmt"
	"testing"

	"go.dkinom.dev/baseconv/options"
)

func TestValidations(t *testing.T) {
	t.Run("alphabet has duplicate characters 1", func(t *testing.T) {
		alphabet := "0123453789"
		b, err := NewBaseConversion(Hexadecimal, alphabet)

		if b != nil || err == nil {
			t.Fatalf("alphabet '%v' has duplicate characters", alphabet)
		}
	})

	t.Run("alphabet has duplicate characters 2", func(t *testing.T) {
		alphabet := "0123453789"
		b, err := NewBaseConversion(alphabet, Binary)

		if b != nil || err == nil {
			t.Fatalf("alphabet '%v' has duplicate characters", alphabet)
		}
	})
}

func TestConverter(t *testing.T) {
	t.Run("invalid input", func(t *testing.T) {
		b, err := NewBaseConversion(Decimal, Binary)
		if err != nil {
			t.Fatal(err)
		}

		s := "123abc"
		r, err := b.Convert(s)

		if r != "" || err == nil {
			t.Fatalf("'%v' is invalid", s)
		}
	})

	{
		b, err := NewBaseConversion(Decimal, Binary)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToBinary = map[string]interface{}{
			"instance": b,
			"testCases": []map[string]string{
				{"from": "0", "to": "0"},
				{"from": "1", "to": "1"},
				{"from": "2", "to": "10"},
				{"from": "3", "to": "11"},
				{"from": "4", "to": "100"},
				{"from": "5", "to": "101"},
				{"from": "6", "to": "110"},
				{"from": "7", "to": "111"},
				{"from": "8", "to": "1000"},
				{"from": "9", "to": "1001"},
				{"from": "10", "to": "1010"},
			},
		}
		testConverter(t, decimalToBinary)
	}
	{
		b, err := NewBaseConversion(
			Decimal, Binary,
			nil, options.BaseConversion().SetZeroPadding(true), nil,
		)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToBinary = map[string]interface{}{
			"instance": b,
			"testCases": []map[string]string{
				{"from": "0", "to": "0000"},
				{"from": "1", "to": "0001"},
				{"from": "2", "to": "0010"},
				{"from": "3", "to": "0011"},
				{"from": "4", "to": "0100"},
				{"from": "5", "to": "0101"},
				{"from": "6", "to": "0110"},
				{"from": "7", "to": "0111"},
				{"from": "8", "to": "1000"},
				{"from": "9", "to": "1001"},
				{"from": "10", "to": "0001010"},
			},
		}
		testConverter(t, decimalToBinary)
	}
	{
		b, err := NewBaseConversion(
			Decimal, Binary,
			options.BaseConversion().SetZeroPadding(true),
		)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToBinary = map[string]interface{}{
			"instance": b.Inverse(),
			"testCases": []map[string]string{
				{"to": "00", "from": "0000"},
				{"to": "01", "from": "0001"},
				{"to": "02", "from": "0010"},
				{"to": "03", "from": "0011"},
				{"to": "04", "from": "0100"},
				{"to": "05", "from": "0101"},
				{"to": "06", "from": "0110"},
				{"to": "07", "from": "0111"},
				{"to": "08", "from": "1000"},
				{"to": "09", "from": "1001"},
				{"to": "010", "from": "0001010"},
			},
		}
		testConverter(t, decimalToBinary)
	}
	{
		b, err := NewBaseConversion(Decimal, "0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣")
		if err != nil {
			t.Fatal(err)
		}
		var decimalToDecimalEmoji = map[string]interface{}{
			"instance": b,
			"testCases": []map[string]string{
				{"from": "0", "to": "0️⃣"},
				{"from": "1", "to": "1️⃣"},
				{"from": "2", "to": "2️⃣"},
				{"from": "3", "to": "3️⃣"},
				{"from": "4", "to": "4️⃣"},
				{"from": "5", "to": "5️⃣"},
				{"from": "6", "to": "6️⃣"},
				{"from": "7", "to": "7️⃣"},
				{"from": "8", "to": "8️⃣"},
				{"from": "9", "to": "9️⃣"},
				{"from": "10", "to": "1️⃣0️⃣"},
			},
		}
		testConverter(t, decimalToDecimalEmoji)
	}
	{
		b, err := NewBaseConversion(
			Decimal, "0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣",
			nil, options.BaseConversion().SetZeroPadding(true), nil,
		)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToDecimalEmoji = map[string]interface{}{
			"instance": b,
			"testCases": []map[string]string{
				{"from": "0", "to": "0️⃣"},
				{"from": "1", "to": "1️⃣"},
				{"from": "2", "to": "2️⃣"},
				{"from": "3", "to": "3️⃣"},
				{"from": "4", "to": "4️⃣"},
				{"from": "5", "to": "5️⃣"},
				{"from": "6", "to": "6️⃣"},
				{"from": "7", "to": "7️⃣"},
				{"from": "8", "to": "8️⃣"},
				{"from": "9", "to": "9️⃣"},
				{"from": "10", "to": "1️⃣0️⃣"},
			},
		}
		testConverter(t, decimalToDecimalEmoji)
	}
	{
		b, err := NewBaseConversion(
			Hexadecimal, "0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣",
			options.BaseConversion().SetZeroPadding(true),
		)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToDecimalEmoji = map[string]interface{}{
			"instance": b.Inverse(),
			"testCases": []map[string]string{
				{"to": "0CAFE", "from": "5️⃣1️⃣9️⃣6️⃣6️⃣"},
				{"to": "0DEADC0DE", "from": "3️⃣7️⃣3️⃣5️⃣9️⃣2️⃣9️⃣0️⃣5️⃣4️⃣"},
			},
		}
		testConverter(t, decimalToDecimalEmoji)
	}
	{
		b, err := NewBaseConversion(
			Hexadecimal, Binary,
			options.BaseConversion().SetZeroPadding(false),
		)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToBinary = map[string]interface{}{
			"instance": b,
			"testCases": []map[string]string{
				{"from": "80", "to": "10000000"},
				{"from": "91", "to": "10010001"},
				{"from": "A2", "to": "10100010"},
				{"from": "B3", "to": "10110011"},
				{"from": "C4", "to": "11000100"},
				{"from": "D5", "to": "11010101"},
				{"from": "E6", "to": "11100110"},
				{"from": "F7", "to": "11110111"},
				{"from": "08", "to": "1000"},
				{"from": "19", "to": "11001"},
				{"from": "2A", "to": "101010"},
				{"from": "3B", "to": "111011"},
				{"from": "4C", "to": "1001100"},
				{"from": "5D", "to": "1011101"},
				{"from": "6E", "to": "1101110"},
				{"from": "7F", "to": "1111111"},
			},
		}
		testConverter(t, decimalToBinary)
	}
	{
		b, err := NewBaseConversion(
			Hexadecimal, Binary,
			options.BaseConversion().SetZeroPadding(false),
		)
		if err != nil {
			t.Fatal(err)
		}
		var decimalToBinary = map[string]interface{}{
			"instance": b.Inverse(),
			"testCases": []map[string]string{
				{"to": "80", "from": "10000000"},
				{"to": "91", "from": "10010001"},
				{"to": "A2", "from": "10100010"},
				{"to": "B3", "from": "10110011"},
				{"to": "C4", "from": "11000100"},
				{"to": "D5", "from": "11010101"},
				{"to": "E6", "from": "11100110"},
				{"to": "F7", "from": "11110111"},
				{"to": "8", "from": "00001000"},
				{"to": "19", "from": "00011001"},
				{"to": "2A", "from": "00101010"},
				{"to": "3B", "from": "00111011"},
				{"to": "4C", "from": "01001100"},
				{"to": "5D", "from": "01011101"},
				{"to": "6E", "from": "01101110"},
				{"to": "7F", "from": "01111111"},
			},
		}
		testConverter(t, decimalToBinary)
	}
}

func testConverter(t *testing.T, converter map[string]interface{}) {
	t.Run("Convert", func(t *testing.T) {
		for _, testCase := range converter["testCases"].([]map[string]string) {
			t.Run(fmt.Sprintf("%v -> %v", testCase["from"], testCase["to"]), func(t *testing.T) {
				actual, err := (converter["instance"]).(*baseConversion).Convert(testCase["from"])
				if err != nil {
					t.Fatal(err)
				}
				expected := testCase["to"]
				if actual != expected {
					t.Fatal()
				}
			})
		}
	})
}
