package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSemordnilaps(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Basic case",
			input:    []string{"desserts", "stressed", "hello", "world"},
			expected: [][]string{{"desserts", "stressed"}},
		},
		{
			name:     "Multiple pairs",
			input:    []string{"dog", "god", "hello", "olleh", "pin", "nip"},
			expected: [][]string{{"dog", "god"}, {"hello", "olleh"}, {"pin", "nip"}},
		},
		{
			name:     "No pairs",
			input:    []string{"hello", "world", "golang"},
			expected: [][]string{},
		},
		{
			name:     "Case sensitive",
			input:    []string{"Hello", "olleH", "hello", "olleh"},
			expected: [][]string{{"hello", "olleh"}, {"Hello", "olleH"}},
		},
		{
			name:     "Duplicate words",
			input:    []string{"level", "level", "radar", "radar"},
			expected: [][]string{},
		},
		{
			name:     "Duplicate words",
			input:    []string{"abc", "cba", "abc", "cba", "bca", "cab"},
			expected: [][]string{{"abc", "cba"}},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "Words with spaces",
			input:    []string{"race a car", "rac a ecar", "hello world"},
			expected: [][]string{{"race a car", "rac a ecar"}},
		},
		{
			name:     "Words with punctuation",
			input:    []string{"race!", "!ecar", "hello!", "!olleh"},
			expected: [][]string{{"race!", "!ecar"}, {"hello!", "!olleh"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Semordnilaps(tt.input)
			if !compareSemordnilaps(result, tt.expected) {
				t.Errorf("Semordnilaps(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// compareSemordnilaps compares two slices of string slices, ignoring order
func compareSemordnilaps(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort each inner slice and the outer slice
	for i := range a {
		sort.Strings(a[i])
	}
	for i := range b {
		sort.Strings(b[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	sort.Slice(b, func(i, j int) bool { return b[i][0] < b[j][0] })

	return reflect.DeepEqual(a, b)
}

func TestRandomSemordnilaps(t *testing.T) {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		input := generateRandomInput(100, 10)
		result := Semordnilaps(input)

		if !compareSemordnilaps(result, SemordnilapCorrect(input)) {
			t.Errorf("Inconsistent results for the same input")
		}
	}
}

func generateRandomInput(maxWords, maxLength int) []string {
	input := make([]string, 0, maxWords)
	chars := "abcdefghijklmnopqrstuvwxyz"

	// Ensure at least one pair
	word := generateRandomWord(maxLength, chars)
	input = append(input, word, reverse(word))

	for len(input) < maxWords {
		word := generateRandomWord(maxLength, chars)
		if word != "" {
			input = append(input, word)
		}
	}

	rand.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})

	return input
}

func generateRandomWord(maxLength int, chars string) string {
	length := rand.Intn(maxLength) + 1 // Ensure non-empty string
	word := make([]byte, length)
	for i := range word {
		word[i] = chars[rand.Intn(len(chars))]
	}
	return string(word)
}

func SemordnilapCorrect(words []string) [][]string {
	wordsSet := make(map[string]bool)
	for _, word := range words {
		wordsSet[word] = true
	}

	pairs := [][]string{}
	for _, word := range words {
		reverse := reverseString(word)
		if wordsSet[reverse] && reverse != word {
			pairs = append(pairs, []string{word, reverse})
			delete(wordsSet, word)
			delete(wordsSet, reverse)
		}
	}

	return pairs
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
