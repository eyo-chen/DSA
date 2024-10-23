# Problem Explanation

Why check `i*i <= n` for the loop condition?<br>
You're asking specifically about why we can stop at `i*i <= n` rather than continuing all the way to `i <= n`. Let's dive deeper into this with a practical example.<br>

Consider the number 100 again. We'll walk through its factors to understand why checking up to `i*i <= n` is sufficient.

The factors of 100 are: 1, 2, 4, 5, 10, 20, 25, 50, 100

Let's pair these factors:

1. 1 x 100
2. 2 x 50
3. 4 x 25
4. 5 x 20
5. 10 x 10

Notice something interesting? The pairs are symmetric around the square root of 100 (which is 10). Once we reach 10, we've actually checked all possible unique factor pairs.

Let's modify our `isPrime` function to illustrate this:

```go
func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("Found factor pair: %d x %d = %d\n", i, n/i, n)
			return false
		}
	}

	return true
}

func main() {
	n := 100
	fmt.Printf("Is %d prime? %v\n", n, isPrime(n))
}
```

When we run this for n = 100, we get:

```
Found factor pair: 2 x 50 = 100
Is 100 prime? false
```

The function stops at the first factor it finds (2), but let's modify it to continue and show all factor pairs up to `i*i <= n`:

```go
func showFactorPairs(n int) {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("Factor pair: %d x %d = %d\n", i, n/i, n)
		}
	}
}

func main() {
	n := 100
	showFactorPairs(n)
}
```

Output:
```
Factor pair: 2 x 50 = 100
Factor pair: 4 x 25 = 100
Factor pair: 5 x 20 = 100
Factor pair: 10 x 10 = 100
```

This output shows us that:

1. We've found all factor pairs without going beyond `i*i <= n`.
2. The last pair (10 x 10) is where `i*i == n`.
3. If we continued beyond this point, we'd just find the same pairs in reverse order.

Why is this important for primality testing?

- If a number is not prime, it must have at least one factor pair where one of the factors is less than or equal to its square root.
- If we haven't found any factors up to the square root, we won't find any beyond it either.

This is why checking up to `i*i <= n` is sufficient:

- It covers all possible unique factor pairs.
- It significantly reduces the number of checks needed, especially for large numbers.
- For prime numbers, it allows us to conclude primality much faster.

For example, to check if 997 is prime:
- Using `i <= n`, we'd need 995 iterations.
- Using `i*i <= n`, we only need 31 iterations (since 32*32 > 997).

This optimization makes the primality test much more efficient, especially for large numbers, without sacrificing accuracy.

## Optimized solution
The idea is loop through 2 to n, and if the number(i) is prime (!hashTable[i])<br>
We want to mark all the i + i + i + i.... to not prime<br>
For example, n = 20<br>
```
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F F F F F F F F  F  F  F  F  F  F  F  F  F  F  
```

We don't need to care 0 and 1, just start at 2 because it's first prime number<br>
i = 2<br>
!isNotPrime[i] = true, which means it's prime number, mark all i + i + i... to true<br>
Because all of them is not prime<br>
```
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F T F T F T F T  F  T  F  T  F  T  F  T  F  T
```

i = 3<br>
!isNotPrime[i] = true, which means it's prime number, mark all i + i + i... to true<br>
Because all of them is not prime<br>
```
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
F F F F T F T F T T T  F  T  F  T  T  T  F  T  F  T
```

i = 4, !isNotPrime[i] = false, which means it's not the prime number, skip it<br>
so on and so forth....<br>
