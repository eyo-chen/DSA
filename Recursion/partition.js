//////////////////////////////////////////////////////
// *** Palindrome Partitioning ***
//////////////////////////////////////////////////////
/*
Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.

A palindrome string is a string that reads the same backward as forward.

Example 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:
Input: s = "a"
Output: [["a"]]
 
Constraints:
1 <= s.length <= 16
s contains only lowercase English letters.
*/
/*
I came up the solution by myself (second time write this problem)
Once finding the backtracking patter, some of problems are not that hard

The core idea is to find all the partition of the input string
Along with the process, we only keep exploring any further partition if the current snippet is palindrome

For example, input is "aab"
For index 0, we can decompose "a", "aa" and "aab"
"a" is palindrome, so can keep exploring
now remaining string is "ab"
we can decompose "a" and "ab"
"a" is palindrome, so can keep exploring
now remaining string is "b"
we can decompose "b"
"a" is palindrome, and we hit the base case

hope this is clear, this is depth-first search

Choice 
=> decompose from index ~ length of input
=> For example, input is "aab"
=> If we're at index 0, we can decompose "a", "aa" and "aab"
=> If we're at index 1, we can decompose "aa" and "aab"

Constraint
=> Each of snippet have to be palindrome

Gaol
=> index is out of the bound of string, aka is equal to the length of string
 
                                                "aab"
                   "a, ab"                     "aa, b"                       "aab"
        "a, a, b"          "a, ab"   


Note how i decompose the string(for-loop part)
The idea is 
For any working index, I know I can decompose all the way to the length of string
For example, "aab"
At index 0, i can decompose
s.slice(0, 1) -> "a"
s.slice(0, 2) -> "ab"
s.slice(0, 3) -> "aab"

At index 1, i can decompose
s.slice(1, 2) -> "a"
s.slice(1, 3) -> "ab"

At index 2, i can decompose
s.slice(2, 3) -> "b"

As we can see, i can go all the way to 3, which is the length of string
The main difference is the working index 

So i just said, i always tryna to decompose all the way to the length of string
BUT, i have to check if (index + i) <= s.length
Because if I'm at index 2, and i will go from 1 ~ 3
It's okay s.slice(2,3)
BUT, it's not okay to keep decomposing, like
s.slice(2,4) or s.slice(2,5)
That's the reason we need if (index + i) <= s.length
Also, the index is gonna throw down to furhter recursive calls is (index + i)
For example, if we finish decomposing s.slice(2, 4)
In next call stack, we want to decompose remaining snippet, which is 
s.slice(4, ....)

Note that .slice(begin, end)
begin -> inclusive
end -> exclusive
For example, s = "1234"
s.slice(0, 4) = "1234"
s.slice(1, 3) = "23"

I know here the index is kinda tricky, but have to fully figure this out 
so that the answer is correct

************************************************************
n = the legnth of s
Time complexity: O(N * 2 ^ N)
=> res.push([...tmp]); and isPalindrome are both O(n) works

Space complexity: O(n)
=> the deepest height of recursive tree is n

Detailed explanation of runtime complexity

A string of length N will have (N, N-1, N-2, ...,1) substrings at positions (0, 1, 2, ..., N-1) respectively. So the total number of substrings of a string is N+N-1+...+1 = O(N2). It is not exponential.

The number 2N in complexity analysis above is in fact the number of nodes in the search tree - not the number of substrings. It is the number of possible partitionings (each partitioning is a way to partition the string into substrings).

This can be derived as follows - Imagine the string as a sequence of N chars separated by a pipe between neighbors, such as a string "abcde" = a|b|c|d|e. Such a representation will have N-1 pipes - in this example, 4 pipes.
If you want the partitioning to have 4 substrings, then you can ask, "how many ways can I select 3 pipes out of the 4 pipes?" - answer is 4 choose 3, i.e. 4C3 = 4. The 4 ways to partition are: { {"a", "b", "c", "de"}, {"a", "b", "cd", "e"}, {"a", "bc", "d", "e"}, {"ab", "c", "d", "e"}
Arguing like the above, the total number of ways to partition this example is when we ask all questions "how many ways can I select 0 or 1 or 2 or 3 or 4 pipes?" = 4C0 + 4C1 + 4C2 + 4C3 + 4C4 = 24 = 16 partitionings
In general a string of length N will have N-1C0 + N-1C1 + ... +N-1CN-2 = 2N-1 = 2N-1 = O(2N) partitionings
*/
/**
 * @param {string} s
 * @return {string[][]}
 */
var partition = function (s) {
  const res = [];

  recursiveHelper(s, 0, [], res);

  return res;

  function recursiveHelper(s, index, tmp, res) {
    // base case (working index is out of the bound)(goal)
    if (index === s.length) {
      res.push([...tmp]);
      return;
    }

    // choice
    for (let i = 1; i <= s.length; i++) {
      if (index + i <= s.length) {
        const subStr = s.slice(index, index + i);

        // constraint (we won't go any further recursive call if it's not palindrome)
        if (isPalindrome(subStr)) {
          tmp.push(subStr);
          recursiveHelper(s, index + i, tmp, res);
          tmp.pop();
        }
      } else break;
    }
    return;
  }
};

function isPalindrome(str) {
  if (str.length === 1) return true;

  let begin = 0;
  let end = str.length - 1;

  while (begin < end) {
    if (str[begin] !== str[end]) return false;

    begin++;
    end--;
  }

  return true;
}

/*
Basically same idea as previos one
The main difference is how we decompose the string

Instead of looping from 0 ~ str.length every time
we use index as starting point
For example, if input is "223"

                                                "aab"
                   "a"                   "aa"
        "a"          "ab"                 "b"
    "b"
For "aab", I can decompose three times, "a", "aa", "aab"
s.slice(0,1)
s.slice(0,2)
s.slice(0,3)
For "a", remaming part is "ab", we don't need to conside "a" (first "a")
s.slice(1,2)
s.slice(1,3)
=> See
=> We don't need to start at 0 anymore

The main point is once finish decomposing, we past i + 1 to the next call stack
which means hey, i've finished decomposing every thing before i, just start at i + 1
Again,
s.slice(0,1) means decompose "a", once we keep recure on
we pass i + 1, 0 + 1, to the next call stack
it means i've finish decomposing "a", just keep decomposing afterwards

This idea makes the code cleaner

It's hard to describe, but it's better to write some easy examples on papers
One thing to note that is both of approach
The next index is gonna pass to next recursive calls is the end idnex in .slice(start, end)


Final note
Because of how we deal with this problem, the main difference is here
is how we decompose the string 
First solution: const subStr = s.slice(index, index + i);
Second solution: const subStr = str.slice(index, i + 1);

Again, try to write on paper to see the difference clearly
*/
function partition1(str) {
  const res = [];

  helperFunction(str, res, [], 0);

  return res;

  function helperFunction(str, res, tmp, index) {
    if (index === str.length) {
      res.push([...tmp]);
      return;
    }

    for (let i = index; i < str.length; i++) {
      /*
      note (i + 1) pass to the recursive calls
      */
      const subStr = str.slice(index, i + 1);
      if (isPalindrome(subStr)) {
        tmp.push(subStr);

        helperFunction(str, res, tmp, i + 1);

        tmp.pop();
      }
    }
  }
}
console.log(partition('abcc'));
// console.log(partition1('aabbccaa'));
