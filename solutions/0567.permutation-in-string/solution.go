package solution

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	m := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		if m[s1[i]] == 0 {
			delete(m, s1[i])
		}

		m[s2[i]]--
		if m[s2[i]] == 0 {
			delete(m, s2[i])
		}
	}

	if len(m) == 0 {
		return true
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		m[s2[i]]++
		if m[s2[i]] == 0 {
			delete(m, s2[i])
		}

		m[s2[i+len(s1)]]--
		if m[s2[i+len(s1)]] == 0 {
			delete(m, s2[i+len(s1)])
		}

		if len(m) == 0 {
			return true
		}
	}

	return false
}
