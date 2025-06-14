package maxdifference

func MaxDifference(s string) int {
	// Array to store frequency of each lowercase letter (a-z)
	freq := [26]int{}

	// Count frequency of each character
	for _, c := range s {
		freq[c-'a']++
	}

	// Variables to track max odd and min even frequencies
	maxOdd, minEven := -1, len(s)

	// Iterate through frequency array
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 1 {
			maxOdd = max(maxOdd, f)
		} else {
			minEven = min(minEven, f)
		}
	}

	return maxOdd - minEven
}
