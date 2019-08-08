package solution

import math

func reverse(x int) int {
    if x == math.MinInt32 {
        return 0
    }

    sign := 1
	reverse := 0

	if x < 0 {
		sign = -1
		x = -x
	}

	for x != 0 {
        if reverse*10.0 + x%10 > math.MaxInt32 {
            return 0
        }

		reverse = reverse*10 + x%10
		x /= 10
	}

	return reverse * sign
}
