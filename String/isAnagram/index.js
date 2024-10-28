//////////////////////////////////////////////////////
// *** Valid Anagram ***
//////////////////////////////////////////////////////
/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
 
Constraints:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
*/
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
/*
Using hashTable to store each character of the string
Then loop through another string to minue the count
If char is not in the hashTable, or it's -1 (over minue)
then return false
*/
var isAnagram = function (s, t) {
  if (s.length !== t.length) {
    return false;
  }

  const hashTable = s.split('').reduce((acc, cur) => {
    acc[cur] = acc[cur] ? acc[cur] + 1 : 1;
    return acc;
  }, {});

  for (const char of t) {
    if (hashTable[char] === undefined || --hashTable[char] === -1) {
      return false;
    }
  }

  return true;
};

/*
Store the char into array by converting them to code number
*/
var isAnagram = function (s, t) {
  if (s.length !== t.length) {
    return false;
  }

  const charHashTable = new Array(26).fill(0);

  for (let i = 0; i < s.length; i++) {
    charHashTable[s[i].charCodeAt(0) - 'a'.charCodeAt(0)]++;
    charHashTable[t[i].charCodeAt(0) - 'a'.charCodeAt(0)]--;
  }

  // they are Anagram if there are all 0
  for (const num of charHashTable) {
    if (num !== 0) {
      return false;
    }
  }

  return true;
};
