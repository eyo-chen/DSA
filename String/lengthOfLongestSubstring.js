//////////////////////////////////////////////////////
// *** Longest Substring Without Repeating Characters ***
//////////////////////////////////////////////////////
/*
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 
Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
/**
 * @param {string} s
 * @return {number}
 */
/*
I came up with the solution by myself
Actually, this problem is relatively easy after practicing a lot sliding window prolbems

hashTable
It represents if any character ever appears before

The main idea is keep moving right pointer
If we hit the duplicate character
It's time to move the left pointer
We'll use while-loop to keep moving left pointer
we'll keep setting the hashTable[leftChar] to fasle
until hashTable[curChar] is false

************************************************************
Time complexity: O(n)
Space complexity: O(n)
*/
var lengthOfLongestSubstring = function (s) {
  const hashTable = [];
  let right = 0;
  let left = 0;
  let max = 0;

  while (right < s.length) {
    // get the current character
    const curChar = s[right].charCodeAt(0);

    // the current character appears before
    if (hashTable[curChar]) {
      // moving left pointer
      while (hashTable[curChar]) {
        // keep setting left pointer of character to false
        hashTable[s[left].charCodeAt(0)] = false;
        left++;
      }
    }

    // set curChar always to true
    hashTable[curChar] = true;
    max = Math.max(max, right - left + 1);
    right++;
  }

  return max;
};

/*
This is another solution from discuss
It solves the problem without inner while-loop

hashTable
=> keep tracking the last found position for each character

the main idea is this line of code
left = Math.max(left, hashTable[curChar] + 1);
When we hitting duplicate character
We know it's time to move the left pointer
We have two choices in this solution
1. move the left pointer to last found position of duplicate character, and plus 1
For example, "abcdeb"
When left = 0, right = 5
We have duplicate character "b"
And hashTable["b"] = 1
It means the last found position of "b" is index 1
In this case, we move left pointer to hashTable["b"] + 1
which is index 2

2. we don't move the left pointer
For example, "abba"
When left = 0, right = 2
We have duplicate character "b"
And hashTable["b"] = 1
It means the last found position of "b" is index 1
In this case, we move left pointer to hashTable["b"] + 1
which is index 2

When left = 2, right = 3
We have duplicate character "a"
And hashTable["a"] = 0
It means the last found position of "b" is index 0
But can we now move left pointer to 0 + 1??
NO
left pointer should never move backward
So what does hashTable["a"] < left means?
It means "hey, yes, we found "a" before, but now left pointer is over this last found position, which means the current sliding window is not includes "a", so we don't need to move left pointer "

************************************************************
Time complexity: O(n)
Space complexity: O(n)
*/
var lengthOfLongestSubstring = function (s) {
  const hashTable = {};
  let right = 0;
  let left = 0;
  let max = 0;

  while (right < s.length) {
    const curChar = s[right];

    if (hashTable[curChar] !== undefined) {
      // great trick to solve the problem
      left = Math.max(left, hashTable[curChar] + 1);
    }

    hashTable[curChar] = right;
    max = Math.max(right - left + 1, max);
    right++;
  }

  return max;
};
