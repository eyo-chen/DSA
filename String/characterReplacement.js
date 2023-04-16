//////////////////////////////////////////////////////
// *** Longest Repeating Character Replacement ***
//////////////////////////////////////////////////////
/*
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
 
Constraints:
1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length
*/
/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
/*
I didn't come up with the solution by myself
But this problem is kind of easy after understanding the core key

The core key is how do we identify the current string is valid string?
For example, s = "AABABBA", k = 1
If the current string(sliding window) is "AABA", is it the valid string?
YES.
Because the (length of string) - (largest frequency) <= k
4 - 3 <= k
Let's see the detail
If we're given the string "AABA", do we change "A" or "B"?
Definitely "B", why?
Because the frequency of "B" is less than "A", it doesn't make senese to change "A"
In other words, in the string "AABA", there's one character need to be changed
one character = (length of string) - (largest frequency)
1 = 4 - 3(frequency of "A")
1 <= k
It's valid string

If we're given the string "AABAB", we still need to change "B"
But, it's invalid string
Because in the string "AABAB", there are two characters need to be changed
two characters = (length of string) - (largest frequency)
2 = 5 - 3(frequency of "A")
2 > k
It's invalid string

So the core key is that 
"We need to keep checking if current string is the valid string"
"Just keep checking when we moving right pointer and left pointer"
"We use this condition(length of string) - (largest frequency) <= k to keep checking"

Note that we use array as frequency, it can also use object as well

I code this solution after watching this video https://www.youtube.com/watch?v=gqXU1UyA8pk

************************************************************
Time: O(26 * n)
=> validStr function will have 26 computation at most
=> But we can still say O(n)

Space: O(1)
=> frequency array will only contain 26 character at most
*/
var characterReplacement = function (s, k) {
  const frequency = new Array(26).fill(0);
  let right = 0;
  let left = 0;
  let max = 0;

  while (right < s.length) {
    // get the ASCII of current string
    const curStr = s[right].charCodeAt(0) - 'A'.charCodeAt(0);

    // add the frequency
    frequency[curStr]++;

    // check if current string is valid
    if (validStr(right - left + 1, frequency, k)) {
      // update the max
      max = Math.max(max, right - left + 1);
    }
    // current string is invalid
    else {
      // keep moving left pointer until current string is valid
      while (left <= right && !validStr(right - left + 1, frequency, k)) {
        // minus the frequency
        frequency[s[left].charCodeAt(0) - 'A'.charCodeAt(0)]--;
        left++;
      }
    }

    /*
    always update right pointer
    It seems doesn't need to update right pointer when current string is invalid
    For example, s = "AABBABA", k = 1
    If current string is "AABB", 
    curChar is "B", and we add the frequency
    But this string is invalid
    we need to update left pointer
    After moving left pointer, current string is "ABB"
    If we didn't update right pointer in this case
    The next iteration,
    curChar is still "B", and we will add the frequency again
    Which is incorrect

    So no matter is valid or invalid
    We always update right pointer
    */
    right++;
  }

  return max;
};

function validStr(len, frequency, k) {
  const maxFrequency = Math.max(...frequency);
  return len - maxFrequency <= k;
}
