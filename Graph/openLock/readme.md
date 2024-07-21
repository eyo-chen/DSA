# Problem Explanation

The problem is we're given a deadend lock and a target lock.<br>
We have to find the shortest path to reach the target lock from '0000'<br>
We're only allowed increment or decrement a digit by 1 for each step.<br>
Also, we can't move to the deadend lock.<br>

Because the problem specifies that the shortest path, we immediately think of using BFS to solve this problem.<br>

The overall idea should be easy<br>
For each exploration, we will generate 8 possible new locks by incrementing or decrementing each digit by 1.<br>
Because for each single digit, we have 2 possible operations, increment or decrement.<br>
For example, suppose we have a lock '1234',<br>
We can generate 8 new locks:<br>
- '0234', '2234'
- '1334', '1134'
- '1244', '1224'
- '1235', '1233'

If the new lock is not in the deadend set and not visited before, we will add it to the queue.<br>
We will continue this process until we find the target lock or the queue is empty.<br>

The hardest part of this problem is how to generate the 8 new locks for each explore.<br>
To be more specific, how to generate '1000' and '9000' from '0000'.<br>
It involves string manipulation.<br>

Small tips:<br>
We can combine `deadendSet` and `seenHash` into one set.<br>
Because we only need to know if the lock is in the deadend or visited before.<br>
But we separate them for better understanding in the code.<br>

# String Manipulation
## C++
For c++, string manipulation is relatively easy.<br>
Because string is mutable in c++.<br>
```c++
in[index] = (in[index] - '0' + 1) % 10 + '0';
```
1. `in[index] - '0'`: This converts the character digit to its integer value. For example, if in[index] is '5', this becomes 5 as an integer.
2. `+ 1`: Increments the integer value.
3. `% 10`: Ensures the result wraps around from 9 to 0.
4. `+ '0'`: It converts the integer result back to its ASCII character representation. For example, if the result is 5, this becomes '5' as a character.

## Go
For Go, string is immutable.<br>
So we can't directly modify a string.<br>

There are two approaches to solve this problem in Go:<br>
1. string concatenation
```go
incrementedString := digit[:index] + strconv.Itoa(int(incrementedDigit)) + digit[index+1:]
```
After incrementing the digit, we concatenate the three parts of the string together.<br>
For example, '12' + '3' + '4' = '1234'<br>
This is less efficient because it creates a new string each time.<br>

2. byte slice
```go
bytes := []byte(digit)
bytes[index] = (bytes[index]-'0'+1)%10 + '0'
```
Since byte slices are mutable, we can directly modify the byte slice.<br>
Each value in the byte slice is a uint8 from 0 to 255, and it represents the ASCII value of the character.<br>
For example, '1' is 49 in ASCII, '2' is 50 in ASCII.<br>
We can first convert '1' to 1 by subtracting '0', then increment it, and finally convert it back to '2' by adding '0'.<br>

### Complexity Analysis For Get In And De Digit In Go

To determine which approach is the most efficient, let's analyze each function in terms of time and space complexity, as well as their practical efficiency:

1. `GenInAndDe`

- Time Complexity: O(n)
  - String slicing (`digit[:index]` and `digit[index+1:]`) is O(n)
  - String concatenation is O(n)
  - Other operations (arithmetic, strconv.Itoa) are O(1)

- Space Complexity: O(n)
  - Creates two new strings, each of length n

2. `GenInAndDe1`:

- Time Complexity: O(n)
  - `[]byte(digit)` conversion is O(n)
  - `make([]byte, len(bytes))` is O(n)
  - copy operations are O(n)
  - Other operations are O(1)

- Space Complexity: O(n)
  - Creates three new byte slices (bytes, inc, dec), each of length n

3. GenInAndDe2:

- Time Complexity: O(n)
  - `[]byte(digit)` conversion is O(n)
  - Other operations are O(1)

- Space Complexity: O(n)
  - Creates three new byte slices (bytes, inBytes, deBytes), each of length n

All three functions have the same time and space complexity in Big O notation. However, in terms of practical efficiency:

1. `genInAndDe1` and `genInAndDe2` are likely to be more efficient than `genInAndDe` because they avoid string concatenation and `strconv.Itoa()`, which can be relatively expensive operations.

2. Between `genInAndDe1` and `genInAndDe2`, `genInAndDe2` is likely to be slightly more efficient because:
   - It avoids the `copy()` operation
   - It uses simple conditional statements instead of modulo arithmetic
   - It directly modifies the byte slices without creating intermediate variables

Therefore, `genInAndDe2` is likely to be the most efficient approach in practice. Here's why:

1. It avoids string concatenation, which can be costly due to the immutability of strings in Go.
2. It doesn't use `strconv.Itoa()`, which involves more operations than simple byte manipulation.
3. It uses direct byte manipulation, which is very fast.
4. Its logic is straightforward and involves fewer arithmetic operations.

Let's be more clear about why string concatenation is costly in Go:<br>
1. String Immutability in Go:
   In Go, ***strings are immutable***. This means that once a string is created, it cannot be changed. ***When you appear to modify a string, you're actually creating a new string.***

2. String Concatenation:
   When you concatenate strings using the `+` operator or `+=`, Go creates a new string each time. This new string contains the combined contents of the original strings.

3. Memory Allocation:
   ***Each new string requires a new memory allocation. The contents of the original strings are copied into this new memory location.*** This is because the original strings are immutable, so they cannot be modified in place.

4. Performance Impact:
   - For small strings or infrequent operations, this isn't a significant issue.
   - However, for larger strings or frequent operations, this can lead to:
     (a) Increased memory usage due to multiple allocations<br>
     (b) Increased CPU time spent on memory allocation and copying<br>
     (c) More pressure on the garbage collector<br>

5. Example:
   Consider this code:
   ```go
   s := "Hello"
   s += ", World"
   s += "!"
   ```
   This creates three distinct strings in memory:
   - "Hello"
   - "Hello, World"
   - "Hello, World!"

6. Contrast with Byte Slices:
   Byte slices (`[]byte`), on the other hand, are mutable. When you modify a byte slice, you're changing the existing memory, not allocating new memory.

7. In the Context of `genInAndDe2`:
   By working with byte slices and only converting to a string at the end, `genInAndDe2` avoids creating multiple intermediate strings. It modifies the bytes in place and only creates a new string once, when returning the result.

8. Performance Optimization:
   For performance-critical code that does a lot of string manipulation, it's often more efficient to work with byte slices and only convert to strings when necessary. 

# Complexity Analysis
## Time Complexity O(1)
- For every single digit, we can have 10 possible values.(0-9)
- There are 4 digits in the lock.
- So, the total number of possible locks is 10^4 = 10000
- We can say the time complexity is O(1)

## Time Complexity O(10^n + d)
- where n is the number of digits
- d is the number of deadends
- The time complexity is O(10^n + d) because we have to visit each lock at most once and each deadend at most once.

## Space Complexity O(1)
- Same as above, we can say the space complexity is O(1)

## Space Complexity O(10^n + d)
- where n is the number of digits
- d is the number of deadends
- The space complexity is O(10^n + d) because we have to store each lock at most once and each deadend at most once.