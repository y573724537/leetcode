package main

func myPow(x float64, n int) float64 {
	if 0 <= n {
		return power(x, n)
	}

	return 1.0 / power(x, -n)
}

func power(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	val := power(x, n/2)

	if n%2 == 0 {
		return val * val
	}

	return val * val * x
}
