package solution

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := make([]byte, 0)

	for i, j := 0, 1; i < len(strs[0]); i++ {
		for ; j < len(strs); j++ {
			if len(strs[j]) <= i {
				break
			}

			if strs[0][i] != strs[j][i] {
				break
			}
		}

		if j == len(strs) {
			prefix = append(prefix, strs[0][i])
			j = 0
		} else {
			break
		}
	}

	return string(prefix)
}
