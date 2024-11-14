package main

func GcdOfStrings(str1 string, str2 string) string {
	// If concatenation in both orders isn't equal, no GCD exists
	if str1+str2 != str2+str1 {
		return ""
	}

	// Return substring of length GCD
	return str1[:gcd(len(str1), len(str2))]
}

// Helper function to calculate GCD using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
