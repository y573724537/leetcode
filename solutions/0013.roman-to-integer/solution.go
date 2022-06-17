package solution

func romanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		if 0 < i && subtraction(s[i-1], s[i]) {
			result -= (val(s[i-1]) * 2)
		}
		result += val(s[i])
	}

	return result
}

func subtraction(prev, cur byte) bool {
	if prev == 'I' && (cur == 'V' || cur == 'X') {
		return true
	}

	if prev == 'X' && (cur == 'L' || cur == 'C') {
		return true
	}

	if prev == 'C' && (cur == 'D' || cur == 'M') {
		return true
	}

	return false
}

func val(cur byte) int {
	if cur == 'I' {
		return 1
	}

	if cur == 'V' {
		return 5
	}

	if cur == 'X' {
		return 10
	}

	if cur == 'L' {
		return 50
	}

	if cur == 'C' {
		return 100
	}

	if cur == 'D' {
		return 500
	}

	if cur == 'M' {
		return 1000
	}

	return 0
}
