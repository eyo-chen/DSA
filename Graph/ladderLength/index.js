//////////////////////////////////////////////////////
// *** Word Ladder ***
//////////////////////////////////////////////////////
/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.

Example 1:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

Example 2:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
 
Constraints:
1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
All the words in wordList are unique.
Accepted
744,860
Submissions
*/
/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
/*
This problem is not that hard
Just using BFS and queue
Every word in wordList is a vertices
And only one difference between vertices is an edge
See the code directly

************************************************************
w = the longest length in the wordList
l = the length of wordList
Time: O(w * l ^ 2)
=> BFS -> |V| + |W|
=> modal as l + l
=> 2l
=> for (let i = 0; i < wordList.length; i++) + countDifference(word, w)
=> w * l
=> 2l * w * l
=> w * l ^ 2
Space: O(l)
=> seen object, and queue
*/
var ladderLength = function (beginWord, endWord, wordList) {
  // if endWord is not in the wordList, just return 0
  if (wordList.find(word => word === endWord) === undefined) return 0;

  const seen = { beginWord: true };
  const queue = [beginWord];

  // keep chaing change variable alogn with the process
  // cuz it also counts beginWord, so we start at 1
  let change = 1;

  while (queue.length > 0) {
    const len = queue.length;

    // loop all the word is only one difference
    for (let i = 0; i < len; i++) {
      const word = queue.shift();

      // find the endWord, return change
      if (word === endWord) return change;

      // loop through wordList
      for (let i = 0; i < wordList.length; i++) {
        const w = wordList[i];

        // if the word has been seen, just skip it
        if (seen[w]) continue;

        // check if two words only have one difference
        if (countDifference(word, w) === 1) {
          // mark it as seen
          seen[w] = true;

          // add to the queue for next iteration
          queue.push(w);
        }
      }
    }

    change++;
  }

  return 0;
};

function countDifference(str1, str2) {
  if (str1.length !== str2.length) return 2;

  let diff = 0;
  for (let i = 0; i < str1.length; i++) {
    if (str1[i] !== str2[i]) diff++;
  }

  return diff;
}
