package solution

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	length := len(s)
	if numRows < length {
		length = numRows
	}

	containers := make([]string, length)
	isDown := false

	for i, j := 0, 0; i < len(s); i++ {
		if j == 0 || j == length-1 {
			isDown = !isDown
		}

		containers[j] += string(s[i])

		if isDown {
			j++
		} else {
			j--
		}
	}

	convert := ""
	for _, s := range containers {
		convert += s
	}

	return convert
}
