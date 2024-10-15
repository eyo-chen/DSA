# Problem Explanation

## Brute Force
The idea is to check all the substrings one by one and find the longest one with unique characters.<br>
For example, for the string "abcabcbb"<br>
First start with "a", then check "ab", "abc", "abca", "abcab", "abcabc", "abcabcb", "abcabcbb"<br>
Then start with "b", then check "bc", "bca", "bcab", "bcabc", "bcabcb", "bcabcbb", "bca"
So on...<br>
Whenever we find a duplicate character, we break the inner loop and start with the next character.<br>

For each outer iteration, we use a hash table to store the characters in the current substring and check for duplicates. So, starting with "a" has it's own hash table, then "ab" has it's own hash table and so on...<br>

There's a gotcha in this solution.<br>
At first, I implement the solution like this:
```go
		for k := i + 1; k < len(s); k++ {
			if _, ok := hashTable[s[k]]; ok {
				ans = max(ans, k-i)
				break
			}

			hashTable[s[k]] = true
		}
```
I put the updating of `ans` logic inside when we find a duplicate character.<br>
That means we only check the length of the current substring when we find a duplicate character.<br>
But this is wrong, let's take an example:
```
s = "au"
```
When we are at "a", we don't have any duplicate character, so we add 'a' to the hash table and continue.<br>
When we are at "au", we don't have any duplicate character, so we add 'u' to the hash table and continue.<br>
We finish the nested loop, and return the answer `ans = 1`.<br>
The correct answer should be `ans = 2`.<br>
The reason we have incorrect answer is because we never update `ans`.<br>
Because the input string has no duplicate character, so we never update `ans` in the first place.<br>
Therefore, we need to update `ans` in every iteration of the nested loop.

### Complexity Analysis
#### Time Complexity O(n^2)
- We have two nested loops, each iterating through the string once.

#### Space Complexity O(1)
- We use a hash table to store the characters in the current substring, but the size of the hash table is at most the size of the alphabet (26 for English letters).


## Sliding Window
We use two pointers, right and left, to represent the current substring.<br>
We also use a hash table to store the characters in the current substring and check for duplicates.<br>

The logic is like this:
1. We start with both right and left pointer at the first character.
2. We first try to update the right pointer to the right
3. If the character at the right pointer is already in the hash table, we need to move the left pointer to the right until there's no duplicate character in the current substring.
4. We update the answer with the maximum length of the current substring.
5. Repeat the process until the right pointer reaches the end of the string.

Let's take "abcabcbb" as an example:
1. Start with left = 0, right = 0, hash table = {}
2. left = 0, right = 1, substring = "ab", hash table = {a}, ans = 1
3. left = 0, right = 2, substring = "abc", hash table = {a, b}, ans = 2
4. left = 0, right = 3, substring = "abca", find the duplicate character 'a', update left pointer
   - left = 1, right = 3, substring = "bca", hash table = {b, c, a}, ans = 3
5. left = 1, right = 4, substring = "bcab", find the duplicate character 'b', update left pointer
   - left = 2, right = 4, substring = "cab", hash table = {c, a, b}, ans = 3
6. left = 2, right = 5, substring = "cabc", find the duplicate character 'c', update left pointer
   - left = 3, right = 5, substring = "abc", hash table = {a, b, c}, ans = 3
7. left = 3, right = 6, substring = "abcb", find the duplicate character 'b', update left pointer
   - left = 4, right = 6, substring = "bcb", keep updating left pointer
   - left = 5, right = 6, substring = "cb", hash table = {c, b}, ans = 3
8. left = 5, right = 7, substring = "cbb", find the duplicate character 'b', update left pointer
   - left = 6, right = 7, substring = "bb"
   - left = 7, right = 7, substring = "b"
9. right = 8, end of the string

In the real code implementation, we don't really need to remove the character from the hash table.<br>
Instead, we just mark it as false in the hash table.

### Complexity Analysis
#### Time Complexity O(n)
- We have two pointers, each iterating through the string once.

#### Space Complexity O(1)
- We use a hash table to store the characters in the current substring, but the size of the hash table is at most the size of the alphabet (26 for English letters).

## Optimized Sliding Window
We can further optimize the sliding window solution by using a hash table to store the index of the characters in the current substring.<br>
This way, we don't need to move the left pointer to the right, we can directly jump to the index of the character in the hash table.

The only difference in the code is following:
```go
		if val, ok := hashTable[s[right]]; ok {
			left = max(left, val+1)
		}
```
This logic is saying: "If the character at the right pointer is already in the hash table, we need to move the left pointer to the right until there's no duplicate character in the current substring."<br>
For example, let's say the input string is "abcdba":<br>
When we're at "d", the hash table is {"a": 0, "b": 1, "c": 2, "d": 3}.<br>
where key is the character, and value is the index of the character in the string.<br>
When we hit "b", we find that "b" is in the hash table with value 1.<br>
That means, we just need to move left pointer to `1+1 = 2`<br>
And the hash table should be updated to {"a": 0, "b": 4, "c": 2, "d": 3}.<br>


There's a gotcha in this solution.<br>
We might attempt to update the left pointer like this:
```go
left = val + 1
```
This logic is wrong, let's take an example:
```
s = "abba"
```
1. Start with left = 0, right = 0, hash table = {}
2. left = 0, right = 1, substring = "ab", hash table = {"a": 0, "b": 1}, ans = 1
3. left = 0, right = 2, substring = "abb", find the duplicate character 'b', update left pointer
   - because the last "b" index is at 1, we update left pointer to `1 + 1 = 2`
4. left = 2, right = 3, substring = "ba", hash table = {"a": 0, "b": 2}, we find the duplicate character 'a', update left pointer
   - note that even though the substring does not contain duplicate character, but the hash table has the duplicate character 'a'
   - in this logic, we update the left pointer to `0 + 1 = 1`
   - which is wrong, because left pointer should not move backwards

To fix this, we need to update the left pointer to the maximum of the current left pointer and the index of the character in the hash table + 1.



### Complexity Analysis
#### Time Complexity O(n)
- We have two pointers, each iterating through the string once.

#### Space Complexity O(1)
- We use a hash table to store the characters in the current substring, but the size of the hash table is at most the size of the alphabet (26 for English letters).