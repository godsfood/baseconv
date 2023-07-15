# baseconv - Base conversion

[![Base conversion](https://raw.githubusercontent.com/godsfood/baseconv/master/.img/baseconv.png)](https://github.com/godsfood/baseconv)

[![Version](https://img.shields.io/github/v/tag/godsfood/baseconv)](https://pkg.go.dev/go.dkinom.dev/baseconv)
[![Build Status](https://github.com/godsfood/baseconv/actions/workflows/.github/workflows/go.yml/badge.svg?branch=master)](https://github.com/godsfood/baseconv/actions)
[![Coverage Status](https://coveralls.io/repos/github/godsfood/baseconv/badge.svg)](https://coveralls.io/github/godsfood/baseconv)
[![License](https://img.shields.io/badge/license-MIT-green)](https://github.com/godsfood/baseconv/blob/master/LICENSE)

A Go module for converting between different bases, e.g., decimal ↔ binary, octal ↔ hexadecimal

## Installing

```bash
go get go.dkinom.dev/baseconv
```

## Usage

See `example/main.go`

#### Hexadecimal ↔ Base58
```go
hexadecimalToBase58Converter, _ := baseconv.NewBaseConversion("0123456789abcdef", baseconv.Base58)

base58Value, _ := hexadecimalToBase58Converter.Convert("415a59758fb933b6049b050a556dd4d916b7b483f6966615")
// base58Value == "6xZA4Qt9vH7rePWeT5WLaVUZNjB6u6rGc"

hexadecimal, _ := hexadecimalToBase58Converter.Inverse().Convert("GjWGF6jERR9ymrC1bHcGmsJYkLMDoaySr")
// hexadecimal == "ac93c8d619c76f823f184110759b278f246cc7cc3cadcac3"
```

#### Decimal (emoji) ↔ Hexadecimal
```go
decimalEmojiToHexadecimalConverter, _ := baseconv.NewBaseConversion(
  "0️⃣1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣7️⃣8️⃣9️⃣", baseconv.Base16,
  options.BaseConversion().
    SetZeroPadding(true),
)

hexadecimal, _ := decimalEmojiToHexadecimalConverter.Convert("5️⃣1️⃣9️⃣6️⃣6️⃣")
// hexadecimal == "0CAFE"

decimalEmoji, _ := decimalEmojiToHexadecimalConverter.Inverse().Convert("DEADC0DE")
// decimalEmoji == "3️⃣7️⃣3️⃣5️⃣9️⃣2️⃣9️⃣0️⃣5️⃣4️⃣"
```

### Exported alphabets
- `Base2` - `01`
- `Base3` - `012`
- `Base4` - `0123`
- `Base5` - `01234`
- `Base6` - `012345`
- `Base8` - `01234567`
- `Base10` - `0123456789`
- `Base12` - `0123456789AB`
- `Base16` - `0123456789ABCDEF`
- `Base32` - `ABCDEFGHIJKLMNOPQRSTUVWXYZ234567`
- `Base32hex` - `0123456789ABCDEFGHIJKLMNOPQRSTUV`
- `Base36` - `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- `Base58` - `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`
- `Base62` - `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
- `Base64` - `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
- `Base64url` - `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_`

## Syntax

### *`NewBaseConversion(from string, to string, opts ...*options.BaseConversionOptions) (*baseConversion, error)`*

*`from`* - String of numeral symbols representing the digits of `from` numeral system.

*`to`* - String of numeral symbols representing the digits of `to` numeral system.

See [documentation](https://pkg.go.dev/go.dkinom.dev/baseconv#section-documentation) for more

## License

[MIT](https://github.com/godsfood/baseconv/blob/master/LICENSE)
