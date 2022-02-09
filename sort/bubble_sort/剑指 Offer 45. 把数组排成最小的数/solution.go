package main

import (
	"bytes"
	"strconv"
)

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	for sentinel := len(strs); 0 < sentinel; {
		tmpSentinel := sentinel
		sentinel = 0
		for i := 0; i+1 < tmpSentinel; i++ {
			if strs[i+1]+strs[i] < strs[i]+strs[i+1] {
				strs[i], strs[i+1] = strs[i+1], strs[i]
				sentinel = i + 1
			}
		}
	}

	var buffer bytes.Buffer
	for i := 0; i < len(strs); i++ {
		buffer.WriteString(strs[i])
	}

	return buffer.String()
}
