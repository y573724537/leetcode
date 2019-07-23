package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	for j := 0; j < n; j++ {
		for i := j; i < m+j; i++ {
			if nums2[j] < nums1[i] {
				nums1[i], nums2[j] = nums2[j], nums1[i]
			}
		}

		nums1[m+j] = nums2[j]
	}
}
