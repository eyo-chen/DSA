# Problem Explanation
The idea to solve this problem is kind of easy<br>
We can use two pointers to check the string from both sides<br>
If the character is different, we can remove either one and check if the rest of the string is palindrome<br>

For example, s = "abca"<br>
We first at "a" and "a", it's a palindrome, then we move the pointer to "b" and "c"<br>
We found "b" and "c" are different,<br>
We remove "c", and check if the rest of the string is palindrome. Yes, the rest of string is "c", it's a palindrome<br>
We remove "b", and check if the rest of the string is palindrome. Yes, the rest of string is "b", it's a palindrome<br>
So the answer is true<br>
When either one of them returns true, it means we can make the string a palindrome by removing at most one character<br>

Let's walk through the process step by step<br>
s = "bececabbacecb"<br>
- s = "bececabbacecb"
  - left = 0, right = 12
  - 'b' = 'b', move on
- s = "ececabbacec"
  - left = 1, right = 11
  - 'e' != 'c', remove either one and check if the rest of the string is palindrome
  - remove 'e'
    - s = "cecabbacec"
    - simply validate if this string is palindrome
    - Yes, it's a palindrome, so we return true
  - remove 'c'
    - s = "ececabbace"
    - simply validate if this string is palindrome
    - No, it's not a palindrome, so we return false
  - after removing either one, we find that the rest of the string is a palindrome after removing 'e', so we return true

# Complexity Analysis
## Time Complexity O(n)
- We only iterate the string once

## Space Complexity O(1)
- We don't use any extra space, so the space complexity is O(1).
