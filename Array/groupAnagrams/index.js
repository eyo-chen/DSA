//////////////////////////////////////////////////////
// *** Group Anagrams ***
//////////////////////////////////////////////////////
/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]
 

Constraints:
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/
/**
 * @param {string[]} strs
 * @return {string[][]}
 */
/*
This problem is fairly easy

This solution using sorting and hash table
The idea is loop through each string
Then sort the string
Using hashTable to store the original string
Two string are Anagrams if they are same after sorting

************************************************************
n = the legnth of array, m = the length of average string in the array
Time compelxity: O(n * m * log(m))
=> Sort each string -> m * log(m)
=> There are n string in the array

Space comelxity: O(n)
*/
var groupAnagrams = function (strs) {
  const hashTable = {};

  for (const str of strs) {
    const sortedStr = str
      .split('')
      .sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0))
      .join('');

    if (hashTable[sortedStr]) {
      hashTable[sortedStr].push(str);
    } else {
      hashTable[sortedStr] = [str];
    }
  }

  return Object.values(hashTable);
};

/*
This is another solution 
https://www.youtube.com/watch?v=vzdNOK2oB2E
Using hashTable

Loop through each string in the array
Create a new array represent char from "a" ~ "z"
Loop through each character of string
and to caculate the frequency for each character
Using array as key, and original string as value

Two string are Anagrams if their char array is the same
For example, "bat", "tab"
The char array should only have b1, a1, t1, and all other chars are 0
So they are Anagrams

Note that
It's important to know char.charCodeAt(0) - 'a'.charCodeAt(0)
char represent each character in the string
And we minus 'a'.charCodeAt(0) to get the corresponding number
Because we want index from 0 ~ 25
0: a
1: b
2: c
3: d
....
Look at this table
https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
a -> 97

If char is b
then 
"b".charCodeAt(0) - 'a'.charCodeAt(0)
=> 98 - 97 = 1
So b is at the correct position

************************************************************
n = the legnth of array, m = the length of average string in the array
Time compelxity: O(n * m)
=> We need to count each character O(m)
=> There are n string in the array O(n)

Space comelxity: O(n)
*/
var groupAnagrams = function (strs) {
  const hashTable = {};

  // O(n)
  for (const str of strs) {
    const charArr = new Array(26).fill(0);

    // O(m)
    for (const char of str) {
      charArr[char.charCodeAt(0) - 'a'.charCodeAt(0)]++;
    }

    if (hashTable[charArr]) {
      hashTable[charArr].push(str);
    } else {
      hashTable[charArr] = [str];
    }
  }

  return Object.values(hashTable);
};
