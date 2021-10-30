package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if (len(nums1)+len(nums2))%2 != 0 {
		return findKthElement(nums1, nums2, (len(nums1)+len(nums2))/2)
	}

	if (len(nums1)+len(nums2))/2 == 0 {
		return 0
	}

	return (findKthElement(nums1, nums2, (len(nums1)+len(nums2))/2) + findKthElement(nums1, nums2, (len(nums1)+len(nums2))/2-1)) / 2
}

func findKthElement(nums1 []int, nums2 []int, targetIdx int) float64 {
	var (
		i      = 0
		j      = 0
		curIdx = 0
		curV   = 0
	)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			curV = nums1[i]
			i++
		} else {
			curV = nums2[j]
			j++
		}

		if curIdx == targetIdx {
			return float64(curV)
		}
		curIdx++
	}

	if i == len(nums1) {
		return float64(nums2[targetIdx-len(nums1)])
	}

	return float64(nums1[targetIdx-len(nums2)])
}
