package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var (
		b                 strings.Builder
		isLastCharEscaper bool
	)

	for k, r := range str {
		if k == 0 && unicode.IsDigit(r) {
			return "", ErrInvalidString // начинается с int
		}

		if k > 1 {
			isLastCharEscaper = string(str[k-1]) == `\` && !isLastCharEscaper
		}

		// если это последний символ
		if k+1 == len(str) {
			if !unicode.IsDigit(r) || (unicode.IsDigit(r) && isLastCharEscaper) {
				b.WriteString(string(r))
			}
			continue
		}

		nextSymbol := string(str[k+1])
		nextRune := rune(str[k+1])

		if unicode.IsDigit(r) && !isLastCharEscaper && unicode.IsDigit(nextRune) {
			return "", ErrInvalidString // два int-а подряд
		}

		if !isLastCharEscaper && (unicode.IsDigit(r) || string(r) == `\`) {
			continue // переход к следующему шагу
		}

		if unicode.IsDigit(nextRune) {
			c, _ := strconv.Atoi(nextSymbol)
			b.WriteString(strings.Repeat(string(r), c))
		} else {
			b.WriteString(string(r))
		}
	}

	return b.String(), nil
}
