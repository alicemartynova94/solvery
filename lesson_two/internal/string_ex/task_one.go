package lesson_two

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("string cannot start with a digit")
var ErrTwoDigitsInRow = errors.New("invalid string format: two digits in a row")
var ErrCannotStartWithZero = errors.New("invalid string format: cannot start with 0")

func UnpackString(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", nil
	}

	runes := []rune(s)
	result := make([]rune, 0)
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	for i, v := range runes {

		if unicode.IsLetter(v) {
			result = append(result, v)

		} else if unicode.IsDigit(v) {

			if unicode.IsDigit(runes[i-1]) {
				return "", ErrTwoDigitsInRow
			} else if v == '0' {
				if len(result) == 0 {
					return "", ErrCannotStartWithZero
				}
				result = result[:len(result)-1]
			} else {
				for j := 0; j < int(v-'0')-1; j++ {
					result = append(result, runes[i-1])
				}
			}
		}
	}

	return string(result), nil
}

func PackString(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}

	runes := []rune(s)
	result := make([]rune, 0)
	first := runes[0]
	count := 1

	for i := 1; i < len(runes); i++ {
		v := runes[i]

		if first == v {
			count++
		} else {
			if count == 1 {
				result = append(result, first)

			} else {
				result = append(result, first)
				result = append(result, rune('0'+count))
				count = 1
			}
			first = v
		}
	}
	result = append(result, first)
	if count > 1 {
		result = append(result, rune('0'+count))
	}

	return string(result)
}
