package main

import (
	"fmt"
	"testing"

	"math/rand"
)

// This is for testing the NearestRepeatedEntries function
func TestNearestRepeatedDistance(t *testing.T) {
	testCases := generateTestCases(1000)

	for i, tc := range testCases {
		result := NearestRepeatedEntries(tc.input)
		fmt.Println("result", result)
		fmt.Println("expected", tc.expected)
		if result != tc.expected {
			t.Errorf("Test case %d failed. Expected %d, got %d", i, tc.expected, result)
			fmt.Printf("Input: %v\n", tc.input)
		}
	}
}

type testCase struct {
	input    []string
	expected int
}

func generateTestCases(count int) []testCase {
	testCases := make([]testCase, count)
	words := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon"}

	for i := 0; i < count; i++ {
		length := rand.Intn(20) + 1 // Random length between 1 and 20
		input := make([]string, length)
		for j := 0; j < length; j++ {
			input[j] = words[rand.Intn(len(words))]
		}
		expected := calculateExpected(input)
		testCases[i] = testCase{input, expected}
	}

	return testCases
}

func calculateExpected(words []string) int {
	lastSeen := make(map[string]int)
	minDistance := -1

	for i, word := range words {
		if lastIndex, exists := lastSeen[word]; exists {
			distance := i - lastIndex
			if minDistance == -1 || distance < minDistance {
				minDistance = distance
			}
		}
		lastSeen[word] = i
	}

	return minDistance
}
