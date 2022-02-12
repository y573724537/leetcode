package main

import "strconv"

func findRelativeRanks(score []int) []string {
	if len(score) == 0 {
		return []string{}
	}

	if len(score) < 2 {
		return []string{"Gold Medal"}
	}

	sorted := shellSort(score)

	mapping := make(map[int]int, len(sorted))
	for i := 0; i < len(sorted); i++ {
		mapping[sorted[i]] = i
	}

	ans := make([]string, len(score))
	for i := 0; i < len(score); i++ {
		idx := mapping[score[i]]
		if idx == 0 {
			ans[i] = "Gold Medal"
		} else if idx == 1 {
			ans[i] = "Silver Medal"
		} else if idx == 2 {
			ans[i] = "Bronze Medal"
		} else {
			ans[i] = strconv.Itoa(idx + 1)
		}
	}

	return ans
}

func shellSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for gap := len(result) >> 1; 0 < gap; gap >>= 1 {
		for i := gap; i < len(result); i++ {
			for j := i; 0 <= j-gap && result[j-gap] < result[j]; j -= gap {
				result[j-gap], result[j] = result[j], result[j-gap]
			}
		}
	}

	return result
}
