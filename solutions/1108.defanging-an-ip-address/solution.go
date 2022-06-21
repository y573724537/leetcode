package solution

func defangIPaddr(address string) string {
	bs := make([]byte, len(address)+6)

	for i, j := 0, 0; i < len(address); i++ {
		if address[i] != '.' {
			bs[j] = address[i]
			j++
			continue
		}

		bs[j] = '['
		j++
		bs[j] = '.'
		j++
		bs[j] = ']'
		j++
	}

	return string(bs)
}
