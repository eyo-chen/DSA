package main

// Brute force solution
func CountPrimes(n int) int {
	ans := 0

	for i := 1; i < n; i++ {
		if isPrime(i) {
			ans++
		}
	}

	return ans
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	// Only need to check up to sqrt(n)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func CountPrimes2(n int) int {
	hashTable := make([]bool, n)
	ans := 0

	for i := 2; i < n; i++ {
		// Find the prime number
		if !hashTable[i] {
			ans++

			// Mark all the i + i + i + i.... to not prime
			for j := i + i; j < n; j += i {
				hashTable[j] = true
			}
		}
	}

	return ans
}
