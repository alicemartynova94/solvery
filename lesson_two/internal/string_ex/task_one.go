package lesson_two

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("string cannot start with a digit")

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
				return "", errors.New("invalid string format: two digits in a row")
			} else if v == '0' {
				if len(result) == 0 {
					return "", errors.New("invalid string format: 0 at start")
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

func PackString(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return "", nil
	}

	runes := []rune(s)
	result := make([]rune, 0)

	runesCount := make(map[rune]int)
	runesCount[runes[0]] = 1

	for i, v := range runes {
		if _, ok := runesCount[v]; !ok {
			runesCount[v] = 1
			val, _ := map[runes[i-1]]
			for j := 0; j < val; j++ {
				result = append(result, runes[i-1])
			}
		}else{
			runesCount[v]++
		}
	}

	return string(result), nil
}
