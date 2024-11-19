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
