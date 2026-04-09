package lesson_two

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("string cannot start with a digit")
var ErrTwoDigitsInRow = errors.New("invalid string format: two digits in a row")

func UnpackString(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", nil
	}

	runes := []rune(s)
	result := make([]rune, 0)
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(runes); i++ {
		v := runes[i]

		if v == '\\' && i+1 < len(runes) {
			next := runes[i+1]
			if unicode.IsLetter(next) {
				return "", ErrInvalidString
			}
			if unicode.IsDigit(next) || next == '\\' {
				result = append(result, next)
				i++
				continue
			}
		}

		if unicode.IsLetter(v) {
			result = append(result, v)
			continue
		}

		if unicode.IsDigit(v) {
			prev := runes[i-1]
			resultI := len(result) - 1
			if unicode.IsDigit(prev) && !unicode.IsDigit(result[resultI]) {
				return "", ErrTwoDigitsInRow
			} else if v == '0' {
				result = result[:resultI]
			} else {
				for j := 0; j < int(v-'0')-1; j++ {
					result = append(result, result[resultI])
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
