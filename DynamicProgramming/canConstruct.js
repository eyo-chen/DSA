//////////////////////////////////////////////////////
// *** Ransom Note ***
//////////////////////////////////////////////////////
/*
Given two stings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:
Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:
Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
function canConstruct(ransomNote, magazine) {
  const arr = Array(26).fill(0);

  for (let i = 0; i < magazine.length; i++) {
    arr[magazine.charCodeAt(i) - 97]++;
  }

  for (let i = 0; i < ransomNote.length; i++) {
    arr[ransomNote.charCodeAt(i) - 97]--;

    if (arr[ransomNote.charCodeAt(i) - 97] < 0) return false;
  }

  return true;
}

// console.log(canConstruct('aa', 'aab'));
