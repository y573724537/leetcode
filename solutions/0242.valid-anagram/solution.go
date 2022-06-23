package solution

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		m[t[i]]--
		if m[t[i]] == 0 {
			delete(m, t[i])
		}
	}

	return len(m) == 0
}
