# Problem Explanation

Let me break down why this solution works step by step:

1. **The Key Insight: String Concatenation Check**
```go
if str1+str2 != str2+str1 {
    return ""
}
```
This check is clever because:
- If two strings have a GCD string, they must be made up of the same repeating pattern
- When you concatenate them in either order, they should be identical
- Example:
  ```
  str1 = "ABCABC" (2 times "ABC")
  str2 = "ABC"    (1 time "ABC")

  str1 + str2 = "ABCABCABC"
  str2 + str1 = "ABCABCABC"  // Same!

  But for strings without GCD:
  str1 = "HELLO"
  str2 = "WORLD"

  str1 + str2 = "HELLOWORLD"
  str2 + str1 = "WORLDHELLO"  // Different!
  ```

2. **Finding the GCD Length**
```go
return str1[:gcd(len(str1), len(str2))]
```
This works because:
- If strings have a common pattern, their lengths must be multiples of that pattern's length
- The GCD of their lengths gives us the length of the base pattern
- Example:
  ```
  str1 = "ABCABCABC"  (length 9)
  str2 = "ABCABC"     (length 6)

  gcd(9,6) = 3        // This gives us length of "ABC"
  ```

3. **Complete Example:**
```go
str1 := "ABCABCABC"  // length = 9
str2 := "ABCABC"     // length = 6

// Check 1: ABCABCABCABCABC == ABCABCABCABCABC ✓
// gcd(9,6) = 3
// Return str1[:3] = "ABC"
```

4. **Why It's Guaranteed to Work:**
- If strings pass the concatenation test, they must be built from the same base pattern
- The GCD of their lengths must be the length of that base pattern
- Taking that many characters from either string will give us the base pattern
- This is why we can safely return `str1[:gcd(len(str1), len(str2))]`

This solution is elegant because it uses mathematical properties (GCD) to solve what appears to be a string problem!
