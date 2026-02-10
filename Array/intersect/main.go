package main

// Work for non-sorted array
func Intersect(nums1 []int, nums2 []int) []int {
	hashTable := map[int]int{}
	for _, n := range nums1 {
		hashTable[n]++
	}

	ans := []int{}
	for _, n := range nums2 {
		if hashTable[n] > 0 {
			ans = append(ans, n)
			hashTable[n]--
		}
	}

	return ans
}

// Work for sorted array
func Intersect1(nums1 []int, nums2 []int) []int {
	p1, p2 := 0, 0
	ans := []int{}

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			ans = append(ans, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}

	return ans
}
