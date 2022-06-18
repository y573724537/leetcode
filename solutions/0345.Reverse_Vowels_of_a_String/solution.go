package solution

func reverseVowels(s string) string {
	reverse := []byte(s)

	i, j := 0, len(s)-1
	for i < j {
		if isVowel(reverse[i]) && isVowel(reverse[j]) {
			reverse[i], reverse[j] = reverse[j], reverse[i]
			i++
			j--
			continue
		}

		if isVowel(reverse[i]) {
			j--
			continue
		}

		if isVowel(reverse[j]) {
			i++
			continue
		}

		i++
		j--
	}

	return string(reverse)
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'A' ||
		c == 'e' || c == 'E' ||
		c == 'i' || c == 'I' ||
		c == 'o' || c == 'O' ||
		c == 'u' || c == 'U' {
		return true
	}

	return false
}
