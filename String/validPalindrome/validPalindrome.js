//////////////////////////////////////////////////////
// *** Valid Palindrome II ***
//////////////////////////////////////////////////////
/*
Given a string s, return true if the s can be palindrome after deleting at most one character from it.

Example 1:
Input: s = "aba"
Output: true

Example 2:
Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.

Example 3:
Input: s = "abc"
Output: false
 
Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
*/
/**
 * @param {string} s
 * @return {boolean}
 */
/*
I fialed this problem at first
But the answer should be very easy

1. We just use to ptrs to check the first and last character
2. If there's a mismatch pair, then we need to remove either one
3. So we just remove either one, if one of them is true, we can return true
For eaxmple,
"bececabbacecb"
"ececabbacec"
Now we have a mis match pair
We first choose remove "c"
"cecabbac"
"ecabba"
Another mis match, now we have to return false

Back to original mis match pair
Now we choose remvoe "e"
"ecabbace"
"cabbac"
"abba"
"bb"
""
done


************************************************************
Time complexity: O(N)O(N).

The main while loop we use can iterate up to N / 2 times, since each iteration represents a pair of characters. On any given iteration, we may find a mismatch and call checkPalindrome twice. checkPalindrome can also iterate up to N / 2 times, in the worst case where the first and last character of s do not match.

Because we are only allowed up to one deletion, the algorithm only considers one mismatch. This means that checkPalindrome will never be called more than twice.

As such, we have a time complexity of O(N)O(N).

Space complexity: O(1)O(1).

The only extra space used is by the two pointers i and j, which can be considered constant relative to the input size.
*/
var validPalindrome = function (s) {
  let p1 = 0;
  let p2 = s.length - 1;

  while (p2 > p1) {
    if (s[p2] !== s[p1]) {
      return checkPalindrome(s, p1 + 1, p2) || checkPalindrome(s, p1, p2 - 1);
    }

    p1++;
    p2--;
  }

  return true;
};

function checkPalindrome(s, p1, p2) {
  while (p2 > p1) {
    if (s[p1] !== s[p2]) {
      return false;
    }

    p1++;
    p2--;
  }

  return true;
}
