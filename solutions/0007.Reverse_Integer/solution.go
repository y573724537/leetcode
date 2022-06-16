package solution

const MAX_INT32 = 1<<31 - 1
const MIN_INT32 = 1 << 31

func reverse(x int) int {
	negative := false

	if x < 0 {
		negative = true
		x = -x
	}

	reverse := 0

	for x != 0 {
		mod := x % 10

		if negative {
			if (MIN_INT32-mod)/10 < reverse {
				return 0
			}
		} else {
			if (MAX_INT32-mod)/10 < reverse {
				return 0
			}
		}
		reverse = reverse*10 + mod
		x /= 10
	}

	if negative {
		return -reverse
	}

	return reverse
}
