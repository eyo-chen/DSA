//////////////////////////////////////////////////////
// *** Reverse String ***
//////////////////////////////////////////////////////
/*
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

Constraints:
1 <= s.length <= 105
s[i] is a printable ascii character.
*/
/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
/*
Just use two ptrs, and swap

************************************************************
Time: O(n)
Space: O(1)
*/
var reverseString = function (s) {
  let len = s.length;
  let p1 = 0;
  let p2 = len - 1;

  while (p2 > p1) {
    [s[p2], s[p1]] = [s[p1], s[p2]];

    p2--;
    p1++;
  }
};
