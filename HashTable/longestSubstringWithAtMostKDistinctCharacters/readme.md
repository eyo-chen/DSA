# Problem Description
Given a string s, find the length of the longest substring that contains at most k distinct characters.


Example 1<br>
Input: s = "coffee", k = 2<br>
Output: 4<br>
Explanation: The substring "ffee" has length four, and it has only two distinct characters. No better (longer) solution exists.

Example 2<br>
Input: s = "aaaaafun", k = 1<br>
Output: 5<br>
Explanation: Substring "aaaaa" has at most 1 distinct character (the 'a'), and it is 5 characters in length.

# Problem Explanation

## Brute force solution
The brute force solution is to check all possible substrings and count the number of distinct characters in each substring.<br>
If the number of distinct characters is less than or equal to k, we update the maximum length of the substring.

Let's walk through the process:<br>
1. We initialize the maximum length of the substring to 0.<br>
2. We iterate through the string and for each character, we check all possible substrings that end at this character.<br>
3. For each substring, we count the number of distinct characters using a hash table.<br>
4. If the number of distinct characters is less than or equal to k, we update the maximum length of the substring.<br>
5. We return the maximum length of the substring.

One thing to note that is how to get a substring in Go.<br>
The syntax `s[i:j+1]` gives us the substring from index i to j inclusive.<br>
Note that j is exclusive, so we need to add 1 to j to get the substring from index i to j inclusive.<br>
For example, `s = "coffee"` and `i = 1` and `j = 3`, then the substring is `"off"`

### Complexity Analysis
#### Time Complexity O(n^3)
- where n is the length of the string
- we have two nested loops to check all possible substrings
- for each substring, we use a hash table to count the number of distinct characters, which takes O(n) time

#### Space Complexity O(n)
- we only use a hash table to store the number of distinct characters, which takes O(n) space


## Sliding window solution
The core idea of the sliding window is to widen the window until we have more than k distinct characters.<br>
When we have more than k distinct characters, we need to shrink the window from the left until we have at most k distinct characters.<br>

Let's walk through the process:<br>
1. We initialize the left and right pointers to 0.<br>
2. We add the right pointer character as distinct character to the hash table.<br>
3. Update right pointer
4. Check if the number of distinct characters is more than k
5. If it is, shrink the window from the left by removing the leftmost character from the hash table
6. Update the maximum length of the substring
7. Return the maximum length of the substring

One thing to note is that we first count the distinct characters, then update the right pointer<br>
Try to walk through the process on paper to understand it better.

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of the string
- we only traverse the string for once
- for each character, we update the hash table and check the condition, which takes O(1) time

#### Space Complexity O(n)
- we use a hash table to store the number of distinct characters in the current window, which takes O(n) space
