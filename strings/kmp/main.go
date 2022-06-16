package main

import (
	"fmt"
)

func getLPS(pattern string) []int {
	lps := make([]int, len(pattern))

	for j, i := 0, 1; i < len(lps); {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j == 0 {
				lps[i] = 0
				i++
			} else {
				j = lps[j-1]
			}
		}
	}
	return lps
}

func strStr(haystack string, needle string) int {
	n := len(needle)

	if n == 0 {
		return 0
	}

	if n == 1 {
		c := needle[0]
		for i := 0; i < len(haystack); i++ {
			if haystack[i] == c {
				return i
			}
		}
		return -1
	}

	if n == len(haystack) {
		if haystack == needle {
			return 0
		}
		return -1
	}
	if len(haystack) < n {
		return -1
	}

	lps := getLPS(needle)
	i, j := 0, 0

	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}
	}

	if j == len(needle) {
		return i - j
	}

	return -1
}

func main() {
	fmt.Printf("%v\n", strStr("abxabcabcaby", "abcaby"))
	fmt.Printf("%v\n", "9"[0]-48)
	fmt.Printf("%v\n", 1<<7-1)
}
