package solution

func reverseVowels(s string) string {
	reverse := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		for i < len(s) && !isVowel(s[i]) {
			i++
		}

		for 0 <= j && !isVowel(s[j]) {
			j--
		}

		if i < j {
			reverse[i], reverse[j] = s[j], s[i]
			i++
			j--
		}
	}

	return string(reverse)
}

func isVowel(b byte) bool {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	_, ok := vowels[b]

	return ok
}

func isVowel(b byte) bool {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	_, ok := vowels[b]

	return ok
}
