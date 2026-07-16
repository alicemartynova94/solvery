package main

func main() {

}

func mergeAlternately(word1 string, word2 string) string {
	result := make([]rune, 0)

	minlen := len(word2)
	if len(word1) < len(word2) {
		minlen = len(word1)
	}

	for i := 0; i < minlen; i++ {
		result = append(result, rune(word1[i]))
		result = append(result, rune(word2[i]))
	}

	if len(word1) > len(word2) {
		for i := len(word1) - len(word2); i < len(word1); i++ {
			result = append(result, rune(word1[i]))
		}
	} else if len(word2) > len(word1) {
		for i := len(word2) - len(word1); i < len(word2); i++ {
			result = append(result, rune(word2[i]))
		}
	}

	return string(result)
}

func mergeAlternately(word1 string, word2 string) string {
	result := make([]rune, 0)

	minlen := len(word2)
	if len(word1) < len(word2) {
		minlen = len(word1)
	}

	k := 0
	for i := 0; i < minlen; i++ {
		result = append(result, rune(word1[i]))
		result = append(result, rune(word2[i]))
		k++
	}

	if len(word1) > len(word2) {
		for i := k; i < len(word1); i++ {
			result = append(result, rune(word1[i]))
		}
	} else if len(word2) > len(word1) {
		for i := k; i < len(word2); i++ {
			result = append(result, rune(word2[i]))
		}
	}

	return string(result)
}
