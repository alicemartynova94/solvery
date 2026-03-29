package lesson_two

import (
	"errors"
	"strings"
	"unicode"
)

func UnpackString(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", nil
	}

	runes := []rune(s)
	result := make([]rune, 0)
	if unicode.IsDigit(runes[0]) {
		return "", errors.New("invalid string: starts with a digit")
	}

	for i, v := range runes {
		if unicode.IsLetter(v) {
			result = append(result, v)
		} else if unicode.IsDigit(v) {
			if unicode.IsDigit(runes[i-1]) {
				return "", errors.New("invalid string format: two digits in a row")
			} else if v == '0' {
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
