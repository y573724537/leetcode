package solution

func lengthOfLongestSubstring(s string) int {
	dict := map[byte]int{}
	max, preLen, curLen := 0, 0, 0

	for i := 0; i < len(s); i++ {
		pos, ok := dict[s[i]]

		if !ok {
			curLen = preLen + 1
		} else {
			conLen := preLen + 1
			posLen := i - pos

			if conLen < posLen {
				curLen = conLen
			} else {
				curLen = posLen
			}
		}

		if max < curLen {
			max = curLen
		}

		dict[s[i]] = i
		preLen = curLen
	}

	return max
}
