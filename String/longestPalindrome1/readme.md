# Problem

Given a string s, return the longest palindromic substring in s.<br>

Example 1:<br>
Input: s = "babad"<br>
Output: "bab"<br>
Explanation: "aba" is also a valid answer.<br>

Example 2:<br>
Input: s = "cbbd"<br>
Output: "bb"<br>

## Approach 1: Brute Force
The idea is to check all the possible substrings of the given string, and check if they are palindrome. If they are, we update the longest palindrome substring.

### Complexity Analysis
#### Time Complexity: O(n^3)
- We need to take O(n^2) time to walk through all the substrings.
- For each substring, we need to take O(n) time to check if it is palindrome.

#### Space Complexity: O(1)
- We don't need any extra space.

## Approach 2: Expand Around Center
The idea is like following<br>
If we're given a string "bab", the most straight forward way to check if it is palindrome is to check if the first and last character are the same, then check the second and second last character, and so on.<br>
Basically, we're checking the string from both sides to the center.(inward)<br>

However, we can also ***check the string from the center to both sides (outward).***<br>
For example, we can start from 'a', then check if the character before and after 'a' are the same, so on and so forth.<br>

The idea is to loop through the string, and for each character, we can check the string from the center to both sides (outward).<br>
In this solution, the time complexity is O(n^2), and the space complexity is O(1).<br>

However, there's a edge case that we need to consider, which is the length of the string is even.<br>
If the string is "cbabc", it's obvious that if we start from "a", we can find the longest palindrome substring "cbabc".<br>
However, if the string is "cbaabc", the logic will FAIL to find the longest palindrome substring "cbaabc".<br>
Because when we start from the middle "a", the previous character is "b", and the next character is "a", which is not the same, so the logic will fail.<br>

To solve this problem, we need to consider two cases:
1. The length of the string is odd.
2. The length of the string is even.

How to consider the length of the string is even?<br>
We can start from the middle two characters, and check if the characters before and after are the same, so on and so forth.<br>
For example, if the string is "cbaabc", when we encounter first "a", we can start from "aa" and check if the characters before and after are the same, so on and so forth.<br>
- "aa" -> "a" and "a" are same
- "baab" -> "b" and "b" are same
- "cbaabc" -> "c" and "c" are same

The last thing is that when we find the longest palindrome substring during the process<br>
We don't need to slice the string to get the actual substring, we only need to know the **start index** and **length** of the substring.<br>
If we know the start index and length, we can directly return the substring by using the start index and length in the end.


### Complexity Analysis
#### Time Complexity: O(n^2)
- We need to take O(n) time to walk through the string.
- For each character, we need to take O(n) time to check the string from the center to both sides (outward).

#### Space Complexity: O(1)
- We don't need any extra space.
