package main

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	m := map[byte]int{}
	for i := 0; i < len(p); i++ {
		m[p[i]]++
		if m[p[i]] == 0 {
			delete(m, p[i])
		}
		m[s[i]]--
		if m[s[i]] == 0 {
			delete(m, s[i])
		}
	}

	ans := []int{}
	if len(m) == 0 {
		ans = append(ans, 0)
	}

	for i := 0; i < len(s)-len(p); i++ {
		m[s[i]]++
		if m[s[i]] == 0 {
			delete(m, s[i])
		}
		m[s[i+len(p)]]--
		if m[s[i+len(p)]] == 0 {
			delete(m, s[i+len(p)])
		}

		if len(m) == 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func main() {
	fmt.Printf("%v\n", findAnagrams("cbaebabacd", "abc"))
	fmt.Printf("%v\n", findAnagrams("abab", "ab"))
}
