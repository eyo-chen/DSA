package main

func StrStr(haystack string, needle string) int {
	left, right := 0, len(needle)

	for right <= len(haystack) {
		if haystack[left:right] == needle {
			return left
		}

		right++
		left++
	}

	return -1
}
