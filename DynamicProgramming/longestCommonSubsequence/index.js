//////////////////////////////////////////////////////
// ***  Longest Common Subsequence ***
//////////////////////////////////////////////////////
/**
 * @param {string} text1
 * @param {string} text2
 * @return {number}
 */
/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.


Example 1:
Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.
*/
/*
I try to solve this problem by myself at first, but didn't work out
The main reason is I can't find the sub problems
Which is the core part of solving DP problems

For example, we're given text1 = "abcde", text2 = "ace" 
It can be saw from begin to end, or end to begin
Here choose to see the string from end to begin

"e", "e"
=> both character are the same, so find the first one common subsequence
=> the subproblem would be what's the longest common subsequence of "abcd" and "ac"
=> We know we can just add one to the answer once the subproblem is solved
=> After finding this common subsequence, we guarantee this common letter or subsequence is in our final answer, so we could just say "hey, i just extract this letter on both string, and solve the rest subsequence(aka. subproblem)"

"d", "c"
=> now the characters are not the same, how to find the subproblem?
=> Again, one thing to guarantee is that, there's no way both of them are in the final answer, it get to be one of them or none of them, but just no way is both of them
=> How could we know is which of them, or none of them?
=> respectively subtract the letter
=> we can either ask what's the longest common subsequence of "abcd" and "a" and what's the longest common subsequence of "abc" and "ac"
=> we just get the maximum value of those two answers

so on and so forth

So the pattern would be
If both character are the same, answer = 1 + (text1[i - 1], text2[k - 1])
else answer = Math.max((text1[i], text2[k - 1]), (text1[i - 1], text2[k]))

Once findign the problem, we could solve it by using DP table or recursion

************************************************************
DP table
n1 = the length of text1, n2 = the length of text2
Time compelxity: O(n1 * n2)
Space comelxity: O(n1 * n2)

Recursion with memoization
n1 = the length of text1, n2 = the length of text2
Time compelxity: O(n1 * n2)
=> We only do n1 * n2 because of memoization

Space comelxity: O(Max(n1, n2))
=> the deepest height of recursive tree 
*/
// DP table
function longestCommonSubsequence(text1, text2) {
  const table = [];

  /*
    extra one for empty string
  
    "" and "adc", the longest common subsequence just 0 
    */
  for (let i = 0; i < text1.length + 1; i++) {
    const row = new Array(text2.length + 1).fill(0);
    table.push(row);
  }

  // the top row and left column just initialize with 0, because it's emtpy string
  for (let i = 1; i < text1.length + 1; i++) {
    for (let k = 1; k < text2.length + 1; k++) {
      // both charater are the same, answer = 1 + (text1[i - 1], text2[k - 1])
      if (text1[i - 1] === text2[k - 1]) table[i][k] = table[i - 1][k - 1] + 1;
      // Math.max((text1[i], text2[k - 1]), (text1[i - 1], text2[k]))
      else table[i][k] = Math.max(table[i - 1][k], table[i][k - 1]);
    }
  }

  return table[text1.length][text2.length];
}

// recursion (scan the string from the beginning)
function longestCommonSubsequence1(text1, text2) {
  return recursiveHelper(text1, text2, 0, 0, {});

  function recursiveHelper(text1, text2, p1, p2, memo) {
    if (p1 >= text1.length || p2 >= text2.length) return 0;

    const key = `${p1},${p2}`;
    if (memo[key] !== undefined) return memo[key];

    if (text1[p1] === text2[p2])
      memo[key] = 1 + recursiveHelper(text1, text2, p1 + 1, p2 + 1, memo);
    else
      memo[key] = Math.max(
        recursiveHelper(text1, text2, p1 + 1, p2, memo),
        recursiveHelper(text1, text2, p1, p2 + 1, memo)
      );

    return memo[key];
  }
}

// recursion (scan the string from the end)
function longestCommonSubsequence2(text1, text2) {
  return recursiveHelper(text1, text2, text1.length - 1, text2.length - 1, {});

  function recursiveHelper(text1, text2, p1, p2, memo) {
    if (p1 < 0 || p2 < 0) return 0;

    const key = `${p1},${p2}`;
    if (memo[key] !== undefined) return memo[key];

    if (text1[p1] === text2[p2])
      memo[key] = 1 + recursiveHelper(text1, text2, p1 - 1, p2 - 1, memo);
    else
      memo[key] = Math.max(
        recursiveHelper(text1, text2, p1 - 1, p2, memo),
        recursiveHelper(text1, text2, p1, p2 - 1, memo)
      );

    return memo[key];
  }
}

// Optimize DP table solution
// Time: O(n1 * n2) / Space: O(Min(n1, n2))
function longestCommonSubsequence3(text1, text2) {
  // find the shorter text, so the guarantee the space complexity will upper bound to O(Min(n1, n2))
  let short = text1,
    long = text2;
  if (text1.length > text2.length) {
    short = text2;
    long = text1;
  }

  let firstRow = new Array(short.length + 1).fill(0);
  let secondRow = new Array(short.length + 1).fill(0);

  for (let i = 1; i < long.length + 1; i++) {
    for (let k = 1; k < short.length + 1; k++) {
      if (short[k - 1] === long[i - 1]) secondRow[k] = firstRow[k - 1] + 1;
      else secondRow[k] = Math.max(firstRow[k], secondRow[k - 1]);
    }

    // swap
    [firstRow, secondRow] = [secondRow, firstRow];
  }

  return firstRow[short.length];
}

// Variant -> return the string of array
/*
  n1 = the length of text1, n2 = the length of text2
  Time compelxity: O(n1 * n2 * Min(n1, n2))
  => [...array] does the O(n) work
  => The longest common subsequence would be Min(n1, n2) -> O(n) work for time and space
  => So the worst case have to mutiply Min(n1, n2)

  Space comelxity: O(n1 * n2 * Min(n1, n2))
  => Basically same idea as time complexity
  */
function longestCommonSubsequence4(text1, text2) {
  const table = [];

  for (let i = 0; i < text1.length + 1; i++) {
    const row = [];
    for (let k = 0; k < text2.length + 1; k++) {
      row.push([]);
    }
    table.push(row);
  }

  for (let i = 1; i < text1.length + 1; i++) {
    for (let k = 1; k < text2.length + 1; k++) {
      // the reason to minus one is beacause having empty string as first character
      if (text1[i - 1] === text2[k - 1]) {
        table[i][k].push(...table[i - 1][k - 1], text1[i - 1]);
      } else {
        if (table[i - 1][k].length >= table[i][k - 1].length)
          table[i][k].push(...table[i - 1][k]);
        else table[i][k].push(...table[i][k - 1]);
      }
    }
  }

  return table[text1.length][text2.length];
}
/*
  Optimize Solution
  n1 = the length of text1, n2 = the length of text2
  Time compelxity: O(n1 * n2)
  => Instead of storing array for each cell, we only store number as same as original problem
  => So now the work inside nested for-loop is only O(1)
  => In order to find the path, we use while-loop to backtrack
  
  Space comelxity: O(n1 * n2)
  => Basically same idea as time complexity
  */
function longestCommonSubsequence5(text1, text2) {
  const table = [];
  const res = [];

  for (let i = 0; i < text1.length + 1; i++) {
    const row = new Array(text2.length + 1).fill(0);
    table.push(row);
  }

  for (let i = 1; i < text1.length + 1; i++) {
    for (let k = 1; k < text2.length + 1; k++) {
      if (text1[i - 1] === text2[k - 1]) table[i][k] = table[i - 1][k - 1] + 1;
      else table[i][k] = Math.max(table[i - 1][k], table[i][k - 1]);
    }
  }

  // from the bottom-right cell to backtrack the answer
  let p1 = text1.length,
    p2 = text2.length;

  while (p1 > 0 && p2 > 0) {
    if (
      table[p1][p2] > table[p1 - 1][p2] &&
      table[p1][p2] > table[p1][p2 - 1]
    ) {
      // again, note that have to minus one because of empty string
      res.unshift(text1[p1 - 1]);
      p1--;
      p2--;
    } else if (table[p1 - 1][p2] > table[p1][p2 - 1]) p1--;
    else p2--;
  }

  return res;
}

// console.log(
//   longestCommonSubsequence5('wefcwefweqxsacwe', 'cwwecwecwecwedwvew')
// );
// console.log(
//   longestCommonSubsequence6('wefcwefweqxsacwe', 'cwwecwecwecwedwvew')
// );
