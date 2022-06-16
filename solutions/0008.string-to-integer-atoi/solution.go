package main

import (
	"fmt"
	"strconv"
)

const MAX_INT32 = 1<<31 - 1
const MIN_INT32 = 1 << 31

func myAtoi(s string) int {
	var result uint
	negative := false

	idx := 0

	// skip beginning space
	for ; idx < len(s); idx++ {
		if s[idx] != ' ' {
			break
		}
	}

	if idx == len(s) {
		return 0
	}

	if s[idx] == '-' {
		negative = true
		idx++
	} else if s[idx] == '+' {
		negative = false
		idx++
	}

	for ; idx < len(s); idx++ {
		if s[idx] != '0' {
			break
		}
	}

	if idx == len(s) {
		return 0
	}

	for ; idx < len(s); idx++ {
		if !isDigit(s[idx]) {
			break
		}

		result = result*10 + asciiToUint(s[idx])

		if !negative && MAX_INT32 <= result {
			return MAX_INT32
		}

		if negative && MIN_INT32 <= result {
			return -MIN_INT32
		}
	}

	if negative {
		return -int(result)
	}

	return int(result)
}

func main() {
	fmt.Printf("%v\n", myAtoi("42"))
	fmt.Printf("%v\n", myAtoi("   -42"))
	fmt.Printf("%v\n", myAtoi("4193 with words"))
	fmt.Printf("%v\n", strconv.Itoa(5))
}

func isDigit(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

func asciiToUint(c byte) uint {
	return uint(c) - 48
}
