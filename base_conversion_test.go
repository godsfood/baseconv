package baseconv

import (
	"fmt"
	"testing"
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
}

func testConverter(t *testing.T, converter map[string]interface{}) {
	t.Run("function", func(t *testing.T) {
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

	t.Run("inverse function", func(t *testing.T) {
		for _, testCase := range converter["testCases"].([]map[string]string) {
			t.Run(fmt.Sprintf("%v -> %v", testCase["to"], testCase["from"]), func(t *testing.T) {
				actual, err := (converter["instance"]).(*baseConversion).Inverse().Convert(testCase["to"])
				if err != nil {
					t.Fatal(err)
				}
				expected := testCase["from"]
				if actual != expected {
					t.Fatal()
				}
			})
		}
	})
}
