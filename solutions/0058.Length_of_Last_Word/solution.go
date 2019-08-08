package solution

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for 0 <= i && s[i] == ' ' {
		i--
	}

	length := 0
	for 0 <= i && s[i] != ' ' {
		i--
		length++
	}

	return length
}
