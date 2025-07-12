package main

import "slices"

func ReverseVowels(s string) string {
	bs := []byte(s)
	vowels := []byte{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		for left < right && !slices.Contains(vowels, bs[left]) {
			left++
		}

		for left < right && !slices.Contains(vowels, bs[right]) {
			right--
		}

		bs[left], bs[right] = bs[right], bs[left]
	}

	return string(bs)
}

func ReverseVowels1(s string) string {
	ans := []byte(s)
	vowels := map[byte]bool{'a': true, 'A': true, 'e': true, 'E': true, 'i': true, 'I': true, 'o': true, 'O': true, 'u': true, 'U': true}
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !vowels[s[left]] {
			left++
		}
		for left < right && !vowels[s[right]] {
			right--
		}

		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return string(ans)
}
