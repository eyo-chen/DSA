//////////////////////////////////////////////////////
// *** Permutation in String ***
//////////////////////////////////////////////////////
/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

Example 1:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
 
Constraints:
1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.
*/
/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
/*
I came up with the solution and wrote the code by myself
After practice couple sliding window problems before

hashTable
=> keep tracking the frequency of each character of s1

match
=> the current match between s1 and s2

The main idea of this problem is similar to other sliding window problem
1. create a hashTable redoring the frequency of each character of s1
2. Iterate thorugh s2 string by moving right pointer
3. If curChar is in the hashTable (hashTable[curChar] !== undefined)
   => decrease the frequency of curChar (hashTable[curChar]--;)
   => If the frequency of curChar is still eqaul or greater than 0, add the match by 1
      -> It's very important to keep tracking the frequency, so that we won't keep adding the match variable
      -> For example, s1 = "hello", s2 = "ooolleoooleh"
      -> Without frequency, we only know that "o" in the s1, and we'll keep adding match when traversing the first three "o", "ooo"
      -> That's not correct
      -> We only have one "o" in s1, so that "ooo" should only match one time
      -> Another example, if sliding window is "oooll"
      -> current match should be 3 
      -> "o" match one time, and "l" match two times
4. If match is eqaul to the length of s1, it means we find the permutation of s1
   => return true
5. If the length of sliding window is equal to s1, it's time to also update left pointer
   => When moving left pointer, it means we remove one character out of the sliding window
   => But we can't just simply remove it, we have to keep tracking the frequency
   => Let's called hashTable[s2[left]] as leftChar
   => If leftChar in the hashTable, we need to increase it's frequency
   => Why?
   => Because if leftChar in the hashTable, it means we definitely decrease it's frequency before when moving right poitner
   => We need to maintain the consistency of frquency, so if we decrease before, now we have to increase it
   => If hashTable[leftChar] is greater than 0, now we have to decrease the match
   => For example, s1 = "hello", s2 = "ooollaaaa"
   => When sliding window is "oooll", and the current match is 3, one "o" and two "l"
   => the next sliding window is "oolla" because we remove the first "o"
   => But can we just simply decrease the match to 2?
   => NO. "oolla" is still match 3, one "o" and two "l"
   => Let's consider the frequency
   => Before entering the while-loop,
   => hashTable: {h:1, e:1, l:2, o:1} -> (character: frequency)
   => When sliding window is "oooll", hashTable: {h:1, e:1, l:0, o:-2}
   => We touch "o" three times, so it decrease three times, so does "l"
   => next sliding window is "oolla", hashTable: {h:1, e:1, l:0, o:-1}, we remove one "o"
   => next sliding window is "ollaa", hashTable: {h:1, e:1, l:0, o:0}, we remove one "o" again
   => next sliding window is "llaaa", hashTable: {h:1, e:1, l:0, o:1}, we remove one "o" again
   => NOW, the frequency of "o" is greater than 1, so that it's time to decrease the match
   => It's correct!!!
   => s1 = "hello" and sliding window "llaaa" are exactly have 2 match, which is two "l"
   => Let's the whole logic
6. Always update right pointer no matter what

************************************************************
n = the legnth of array
Time compelxity: O(n)
Space comelxity: O(1)
=> Note that the size of hashTable won't over 26
=> Because we only store 26 character at most
*/
var checkInclusion = function (s1, s2) {
  const hashTable = {};
  let right = 0;
  let left = 0;
  let match = 0;

  // create hashTable
  for (const char of s1) {
    hashTable[char] = hashTable[char] ? hashTable[char] + 1 : 1;
  }

  while (right < s2.length) {
    const curChar = s2[right];

    // check if curChar is in the hashTable
    if (hashTable[curChar] !== undefined) {
      // always decrease it's frequency
      hashTable[curChar]--;

      // can only count match if it's frquency is greater and equal 0
      if (hashTable[curChar] >= 0) {
        match++;
      }
    }

    // find all the match
    if (match === s1.length) {
      return true;
    }

    // the length of sliding window is equal to s1, it's time to move left pointer too
    if (right - left + 1 === s1.length) {
      // check if leftChar is in the hashTable
      if (hashTable[s2[left]] !== undefined) {
        // always increase it's frequency
        hashTable[s2[left]]++;

        // can only decrease match if it's frequency is greater than 0
        if (hashTable[s2[left]] > 0) {
          match--;
        }
      }
      left++;
    }

    right++;
  }

  return false;
};

/*
This is another approach from discuss section

1. create the frequency array of s2
2. Iterate through s2 by moving right pointer
3. increase the frequency2 array
4. If the length of sliding window is equal to s1
   => Check if two frequency is the same
   => return true if it's the same
   => move the left pointer, also decrease the frequency array2

Basically, the main idea of this problem is using two frequency array
And we keep checking these two frequency array is the same
If it's the same, return true
Also, when moving right pointer, it means we add new character, increase the frequency
when moving left pointer, it means we remove an old character, decrease the frequency

See the discuss section
https://leetcode.com/problems/permutation-in-string/discuss/1762469/C%2B%2B-oror-SLIDING-WINDOW-OPTIMIZED-oror-Well-Explained

************************************************************
n = the legnth of array
Time compelxity: O(26 * n)
=> Note that we may invoke compareFrequency function for about n times
=> And compareFrequency function only does O(26) times at most
=> So we can stil say it's O(n) times
Space comelxity: O(1)
=> Note that the size of two array won't over 26
=> Because we only store 26 character at most
*/
var checkInclusion = function (s1, s2) {
  const frequency1 = new Array(26).fill(0);
  const frequency2 = new Array(26).fill(0);
  let right = 0;
  let left = 0;

  for (const char of s1) {
    const charCode = char.charCodeAt(0) - 'a'.charCodeAt(0);
    frequency1[charCode]++;
  }

  while (right < s2.length) {
    const curChar = s2[right].charCodeAt(0) - 'a'.charCodeAt(0);
    frequency2[curChar]++;

    if (right - left + 1 === s1.length) {
      if (compareFrequency(frequency1, frequency2)) {
        return true;
      }

      const leftChar = s2[left].charCodeAt(0) - 'a'.charCodeAt(0);
      frequency2[leftChar]--;
      left++;
    }

    right++;
  }

  return false;
};

function compareFrequency(arr1, arr2) {
  for (let i = 0; i < arr1.length; i++) {
    if (arr1[i] !== arr2[i]) {
      return false;
    }
  }

  return true;
}