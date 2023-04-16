//////////////////////////////////////////////////////
// *** Word Break ***
//////////////////////////////////////////////////////
/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
*/
/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
/*
Unlike previous question asking us to return the "BEST" answer
This question only asks us "CAN" answer, 
which means it dones't care the best possible to construct the string
it only cares can we do this?

So we can earily return inside the for-loop
Because once we find the very first possibilities to construct the string
We can just return the answer
*/
/*
Top-Down
Non-Optimize solution

************************************************************
s = length of string, w = length of wordDict(array)

Time compelxity: O(w ^ s * s)
=> The recursive tree will have s deep, and for each node will have w branching factor at worst
=> Also, for each node, we will do O(n) of s work because of s.slice(word.length)

Space comelxity: O(s ^ 2)
=> Because the maximum call stack will have s deep
=> And for each node, we will create O(n) space because of s.slice(word.length)
*/
function wordBreak(s, wordDict) {
  if (s === '') return true;

  for (const word of wordDict) {
    // indexOf returns 0 meaning it's preffix of string
    // "string".indexOf("str") = 0, "string".indexOf("stri") === 0
    if (s.indexOf(word) === 0) {
      const remainingStr = s.slice(word.length);

      if (wordBreak(remainingStr, wordDict)) {
        return true;
      }
    }
  }

  return false;
}

/*
  Top-Down
  Optimize solution with memoization
  
  ************************************************************
  s = length of string, w = length of wordDict(array)
  
  Time compelxity: O(s * w * s)
  => At maximum work of this solution, we will only have s amount of word because we use "memo" object
     So we won't do any duplicate works, and for each work, we iterate w times
  => Same thing as above, the additional s is for s.slice(word.length);
  
  Space comelxity: O(s ^ 2)
  => Same thing as above
  */
function wordBreak1(s, wordDict, memo = {}) {
  if (memo[s] !== undefined) return memo[s];
  if (s === '') return true;

  for (const word of wordDict) {
    if (s.indexOf(word) === 0) {
      const remainingStr = s.slice(word.length);

      if (wordBreak(remainingStr, wordDict, memo)) {
        memo[s] = true;
        return true;
      }
    }
  }

  memo[s] = false;
  return false;
}

/*
  Bottom-Up (Tabulation)
  
  Main Idea
  1. Create DP table, and the length is set to (string.length + 1) as usual
     => For each element of the DP table, it represents the each character of given string
     => And the very first element is prepresents the empty string, that's why we initialize to true
        Because we can always construct empty string
     => It's very important to understand what does the DP table represent
  
  2. Iterate through each element of DP table
  
  3. If the element of DP table is true, it means we can keep exploring from this position because we have made sure this position can be achieved
     For example, s = "hello", wordDict = ["he", "llo"]
     The DP table would look like this
  
     "" h  e  l  l  o
     T  F  F  F  F  F
     __ __ __ __ __ __ 
     0  1  2  3  4  5
  
     for i = 0, table[i] = true
     s.slice(0, "he".length + 0) = "he"
     So we can set table[0 + 2] = true
  
     "" h  e  l  l  o
     T  F  T  F  F  F
     __ __ __ __ __ __ 
     0  1  2  3  4  5
  
     It means we have found one way to reach the index 2
     In other words, we have found "he" can be constructed to "hello"
  
     for i = 1, table[i] = false
     It means can't use "h" to construct, so there's no reason to try any subString from "e"
     There's no "h", "ello"... in the wordDict
  
     for i = 2, table[2] = true
     It means we can and we have used "he", so we can keep exploring any possibilites after "he" which is "llo"
     To check if there's any way to construct "llo" from wordDict
     This is the core idea of this algorithm
  
     Later we found that 
     s.slice(2, "llo".length + 0) = "llo"
     So we can set table[2 + 3] = true
  
  4. The core idea of this algorithm is
     "" h  e  l  l  o
     T  F  T  F  F  F
     __ __ __ __ __ __ 
     0  1  2  3  4  5
  
     If index is 2, we know we can loop through wordDict
     And what we wanna ask is 
     Can we find any element in the wordDict to construct being equal to the subString after the position I'm currently on
     The subString after the position I'm currently on is "llo"
  
  4. In this question, it's also very important to deeply undestand "slice" method
     If it takes two arguements
     The first index is "start", it's inclusive
     The second index is "end", it's exclusive
     For example, s = "hello"
     s.slice(0, 2) = "he"
     s.slice(2, 4) = "ll" 
  
  
  ************************************************************
  s = length of string, w = length of wordDict(array)
  
  Time compelxity: O(s * w * s)
  => The outer loop takes s amount of work
  => The inner loop takes w amount of work
  => For each work, we also do s amount of work because of slice method
  
  Space comelxity: O(s)
  => DP table
  */
function wordBreak2(s, wordDict) {
  const table = new Array(s.length + 1).fill(false);
  table[0] = true;

  for (let i = 0; i < s.length + 1; i++) {
    // table[i] represent sub problem
    if (table[i] === true) {
      for (const word of wordDict) {
        if (s.slice(i, word.length + i) === word) table[i + word.length] = true;
      }
    }
  }

  return table[s.length];
}

// function wordBreak4(s, wordDict) {
//   const table = new Array(s.length + 1).fill(false);
//   table[0] = true;

//   for (let i = 0; i < table.length; i++) {
//     const curSubProblem = table[i];

//     if (curSubProblem === true) {
//       for (const word of wordDict) {
//         if (word.length + i <= s.length) {
//           const subString = s.slice(i, word.length + i);

//           if (subString === word) {
//             table[word.length + i] = true;
//           }
//         }
//       }
//     }
//   }

//   return table[s.length];
// }

/*
  The is also a Bottom-Up approach
  But has different mindset
  
  For example, s = "hello", wordDict = ["he", "llo"]
  
  1. Create DP table, and it also represents each character of given string
     "" h  e  l  l  o
     T  F  F  F  F  F
     __ __ __ __ __ __ 
     0  1  2  3  4  5
  
     For outer for-loop, we still loop through each cell of DP table
     For inner for-loop, we're finding any possibilites is equal to the word in wordDict
  
     Base case, empty string
  
     For example, 
     for i is 2, the question is "Can I find any substring of "he" to match the word in wordDict?"
     k = 0, subString = "he" (s.slice(0, 2))
     k = 1, subString = "e"  (s.slice(1, 2))
     k = 3, subString = ""   (s.slice(2, 2))
  
     We know that "he" is in the wordDict
     It means we have found one subString in the wordDict in the index 2, and can keep constructig later subString, like "llo"
     So set it to true
  
     "" h  e  l  l  o
     T  F  T  F  F  F
     __ __ __ __ __ __ 
     0  1  2  3  4  5
  
     But for i is 3, the question is "Can I find any substring of "hel" to match the word in wordDict?"
     k = 0, subString = "hel" (s.slice(0, 3))
     k = 1, subString = "el"  (s.slice(1, 3))
     k = 2, subString = "l"   (s.slice(2, 3))
     k = 3, subString = ""   (s.slice(3, 3))
     We can't find any subString in the wordDict
     It means in the index 3, we can't keep construcing later subString like "lo"
  
  
     for i is 4, the question is "Can I find any substring of "hell" to match the word in wordDict?"
     k = 0, subString = "hell" (s.slice(0, 4))
     k = 1, subString = "ell"  (s.slice(1, 4))
     k = 2, subString = "ll"   (s.slice(2, 4))
     k = 3, subString = "l"   (s.slice(3, 4))
     k = 4, subString = ""   (s.slice(4, 4))
     We can't find any subString in the wordDict
     It means in the index 4, we can't keep construcing later subString like "o"
  
     for i is 5, the question is "Can I find any substring of "hello" to match the word in wordDict?"
     k = 0, subString = "hello" (s.slice(0, 5))
     k = 1, subString = "ello"  (s.slice(1, 5))
     k = 2, subString = "llo"   (s.slice(2, 5)) -> earily break
     k = 3, subString = "lo"   (s.slice(3, 4))
     k = 4, subString = "o"   (s.slice(4, 4))
     k = 5, subString = ""   (s.slice(4, 4))
     Once finding one possibility to match the word in wordDict, we can earily break
  
  The core idea of this approach is quite different from previous one
  For each cell, the question is "Can I find any subString from this index to match the word in wordDict?"
  Once finding 
  
  
  
  ************************************************************
  s = length of string, w = length of wordDict(array)
  
  Time compelxity: O(s * w)
  
  Space comelxity: O(s)
  */
function wordBreak3(s, wordDict) {
  const table = new Array(s.length + 1).fill(false);
  table[0] = true;

  for (let i = 1; i < s.length + 1; i++) {
    for (let k = 0; k < i; k++) {
      if (table[k] === true && wordDict.includes(s.slice(k, i))) {
        table[i] = true;
        break;
      }
    }
  }

  return table[s.length];
}

// console.log(wordBreak3('hello', ['he', 'llo']));lfi
