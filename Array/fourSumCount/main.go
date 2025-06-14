package foursumcount

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	res := 0

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				for h := 0; h < n; h++ {
					if nums1[i]+nums2[k]+nums3[j]+nums4[h] == 0 {
						res++
					}
				}
			}
		}
	}

	return res
}

func FourSumCount1(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// Create a hash map to store sums of pairs from nums1 and nums2
	sumCount := make(map[int]int)

	// Step 1: Compute all sums of pairs (i, j) from nums1 and nums2
	for _, a := range nums1 {
		for _, b := range nums2 {
			sum := a + b
			sumCount[sum]++
		}
	}

	// Step 2: Iterate over pairs from nums3 and nums4, and count complementary sums
	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			target := -(c + d) // We need nums1[i] + nums2[j] = -(nums3[k] + nums4[l])
			count += sumCount[target]
		}
	}

	return count
}
