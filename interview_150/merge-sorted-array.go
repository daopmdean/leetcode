package interview_150

func merge(nums1 []int, m int, nums2 []int, n int) {
	cp := make([]int, m)
	for i := 0; i < m; i++ {
		cp[i] = nums1[i]
	}

	if m == 0 {
		for i, v := range nums2 {
			nums1[i] = v
		}
		return
	}

	if n == 0 {
		return
	}

	p, i, j := 0, 0, 0
	for p < m+n {
		if i >= m {
			nums1[p] = nums2[j]
			j++
		} else if j >= n {
			nums1[p] = cp[i]
			i++
		} else if cp[i] <= nums2[j] {
			nums1[p] = cp[i]
			i++
		} else {
			nums1[p] = nums2[j]
			j++
		}
		p++
	}
}
