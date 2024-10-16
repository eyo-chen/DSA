# Problem Explanation

## Brute Force
The idea is simple, just check all the substrings of s2, if any of the substring is the permutation of s1, return true.<br>
For example, if s1 is "ab" and s2 is "eidbaooo"<br>
We start at "e", then check "ei",<br>
Update to "i", then check "id",<br>
Update to "d", then check "db",<br>
So on...<br>
We only need to check the substring of length of s1<br>
In this example, because s1 is length 2, we only need to check the substring of length 2 for each character in s2.<br>

For each substring, how do we check if it's a permutation of s1?<br>
We first use a hash table to store the frequency of each character in s1.<br>
Then we use another hash table to store the frequency of each character in the current substring of s2.<br>
If the two hash tables are the same, then the substring is a permutation of s1.<br>
This comparison only need O(26) = O(1) time, because the length of the hash table is 26 (for each character in alphabet).

### Complexity Analysis
#### Time Complexity O(n * m)
- n is the length of s1
- m is the length of s2
- the outer loop runs m times, and inner loop runs n times, so it's O(n * m)

#### Space Complexity O(1)
- We are using a hash table to store the frequency of each character in s1, which is constant space.

## Sliding Window
The idea is to use a sliding window to check if the substring of s2 is a permutation of s1.<br>
For example, if s1 is "ab" and s2 is "eidbaooo"<br>
The window size is 2, and we slide it from left to right.<br>
```
r   l
e i d b a o o o o
```
Move both right and left pointer to right to check the next window.
```
  r   l
e i d b a o o o o
```

Now, we know that we can use a sliding window to find the substring of s2<br>
How do we check if the substring is a permutation of s1?<br>
First, we use a hash table to store the frequency of each character in s1.<br>
Then we use another hash table to store the frequency of each character in the current window of s2.<br>
If the two hash tables are the same, then the substring is a permutation of s1.<br>
When we move the right pointer to the next character, we increment the frequency for this character in the hash table.<br>
When we move the left pointer to the next character, we decrement the frequency for this character in the hash table.<br>

Let's summarize the steps:<br>
1. Create frequency hash table for s1
2. Initialize two pointers, right and left, both at 0
3. Create frequency hash table for s2
4. While right pointer is less than the length of s2
  - Add the character of right pointer to freq2
  - If the window size is equal to s1's length
    - Check if freq2 is equal to freq1
    - If yes, return true
    - If no, continue
  - If the window size is not equal to s1's length
    - That means we haven't get a full window, move right pointer to right
  - If the window size is equal to s1's length
    - That means we have a full window, and the current window is not a permutation of s1
    - Update window by moving left pointer and right pointer to right
    - decrement the frequency of the character of left pointer in freq2

Let's walk through the example:<br>
Suppose s1 is "abc" and s2 is "abbca"<br>
We start with freq1 = {a: 1, b: 1, c: 1}, right and left pointer at 0<br>
- First iteration
  - right pointer at 0, left pointer at 0
  - add the character of right pointer to freq2, freq2 = {a: 1}
  - check if the window size is equal to s1's length
  - No, move right pointer to 1
- Second iteration
  - right pointer at 1, left pointer at 0
  - add the character of right pointer to freq2, freq2 = {a: 1, b: 1}
  - check if the window size is equal to s1's length
  - No, move right pointer to 2
- Third iteration
  - right pointer at 2, left pointer at 0
  - add the character of right pointer to freq2, freq2 = {a: 1, b: 2}
  - check if the window size is equal to s1's length
  - Yes, now window is "abb", check if freq2 is equal to freq1
  - No, we need to move both right and left pointer to right
  - Move left pointer to 1, and right pointer to 3
  - decrease the frequency of the character of left pointer in freq2, freq2 = {b: 2}
- Fourth iteration
  - right pointer at 3, left pointer at 1
  - add the character of right pointer to freq2, freq2 = {b: 2, c: 1}
  - check if the window size is equal to s1's length
  - Yes, now window is "bbc", check if freq2 is equal to freq1
  - No, we need to move both right and left pointer to right
  - Move left pointer to 2, and right pointer to 4
  - decrease the frequency of the character of left pointer in freq2, freq2 = {b: 1, c: 1}
- Fifth iteration
  - right pointer at 4, left pointer at 2
  - add the character of right pointer to freq2, freq2 = {b: 1, c: 1, a: 1}
  - check if the window size is equal to s1's length
  - Yes, now window is "bca", check if freq2 is equal to freq1
  - Yes, return true

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of s2
- we loop through s2 once, and each operation is O(1), so it's O(n)

#### Space Complexity O(1)
- we use two hash tables to store the frequency of each character in s1 and s2, which is constant space.


## Optimized Sliding Window
The idea is similar to the above sliding window. However, we don't need to create a hash table for s2, and we don't need to compare two hash tables.<br>
Instead, we use a variable to count the number of characters in s2 that are in s1.<br>
For example, if s1 is "abc" and s2 is "abbca"<br>
When the window is "abb", we find that there are only two characters in s2 that are in s1, so count = 2<br>
When the window is "bbc", we find that there are two characters in s2 that are in s1, so count = 2<br>
When the window is "bca", we find that there are three characters in s2 that are in s1, so count = 3<br>
When count equals to the length of s1, we find a permutation of s1 in s2, return true<br>
Therefore, for each window, we have to find the number of characters in s2 that are in s1<br>

How can we maintain this variable which is the number of characters in s2 that are in s1?<br>
Let's split two parts to explain:
1. When do we increase the count?
2. When do we decrease the count?

When do we increase the count?<br>
When right pointer encounter a character,<br>
It first checks if this character is in s1<br>
If yes, that means the current character is in s1<br>
However, we can't directly increase the count<br>
We also need to check the frequency<br>
What does that mean?<br>
For example, s1 is "ab", and s2 is "aaaab"<br>
When right pointer is at 1, "a" is in s1, but can we increase the count?<br>
No, because s1 only has one "a"<br>
If we do increase the count, then that means the count will be 2<br>
But this 2 means that there are two "a" in the current window<br>
However, s1 only has one "a"<br>
The correct way to handle this is to<br>
First, decrease the frequency of the character in s1<br>
If the frequency is still equal to or greater than 0, that means the character is in s1, and the frequency is still within the limit<br>
So we increase the count<br>

When do we decrease the count?<br>
This is very similar to the increase count<br>
When left pointer encounter a character,<br>
It first checks if this character is in s1<br>
If yes, that means the current character is in s1<br>
Also, we first increase the frequency of the character in s1<br>
Then, we check if the frequency is greater than 0<br>
If yes, that means the character is in s1, and the frequency is still within the limit<br>
So we decrease the count<br>

### Complexity Analysis
#### Time Complexity O(n)
- where n is the length of s2
- we loop through s2 once, and each operation is O(1), so it's O(n)

#### Space Complexity O(1)
- we use two hash tables to store the frequency of each character in s1 and s2, which is constant space.



