package main

import (
	"fmt"
)

var result = []string{}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return result
	}

	backTracking("", 0, 0, n)

	return result
}

func backTracking(cur string, l int, r int, sum int) {
	if len(cur) == sum*2 {
		result = append(result, cur)
		return
	}

	if l < sum {
		backTracking(cur+"(", l+1, r, sum)
	}

	if r < l {
		backTracking(cur+")", l, r+1, sum)
	}
}

func main() {
	fmt.Printf("%v\n", generateParenthesis(4))
}
