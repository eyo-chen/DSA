//////////////////////////////////////////////////////
// *** Decode Ways ***
//////////////////////////////////////////////////////
/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:

"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.

The answer is guaranteed to fit in a 32-bit integer.


Example 1:

Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
Example 3:

Input: s = "0"
Output: 0
Explanation: There is no character that is mapped to a number starting with 0.
The only valid mappings with 0 are 'J' -> "10" and 'T' -> "20", neither of which start with 0.
Hence, there are no valid ways to decode this since all digits need to be mapped.
Example 4:

Input: s = "06"
Output: 0
Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").
*/

/*
There's two main mistakes or traps in this question
1. We don't need create a hash table to map number and character because the problem only ask us the total counts
2. We don't need to find all the decomposition(like one of the qeustion in recursive section)(Palindromic Decompositions)
   Because the constraints of the problem, which is each string should map to chaeacter (from "1" ~ "26")
   For example, we're given string "123",
   In first call stack, it's okay to decompose "1" and "12"
   But there's NO reason to find "123" because three digits number are always invalid in this problem
   This point is very important
   Again, if we're given long string "1234567890"
   In Palindromic Decompositions, we need do sth like this
                                                            "1234567890"
              "1", "234567890"   "12", "34567890"        "123", "4567890"     "1234", "567890"     "12345", "7890"
   As we could see, the branching factor would be the lenght of string, which is bad

   In this problem, we only need do this 
                                                             "1234567890"
              "1", "234567890"   "12", "34567890"        ................
   Because any subString after "12", would be invalid
*/
/*
The process is very similar to the question of Palindromic Decompositions
but the branching factor is always be 2, which means we only explore any possibilities less than or equal to two digits

There are two things need to be careful along with coming up the solution
1. use "index" to keep tracking "where am i starting decomposing the string"
   For example, the string "12345"
   If index is 0, then the possibilities are "1" and "12"
   If index is 2, then the possibilities are "3" and "34"

   Note that index is 0-based
  
2. When invoke the next recursive function, we have to pass "index + i"
   For example, the string "12345", in the first for-loop, the i is 1, and index is 2
   subStr is "3"
   In the next call stack, we only want to start doing the decomposition at the index 3
   So again in the first for-loop, the i is 1, and index is 3
   subStr is "4"

3. We need this if-statement if (index + i <= s.length)
   In some edge case, index + i will out of the bound of the length of string
   For example, the string "123", we have the case is index is 2, and in the second for-loop, so the i is 2
   Note that now haven't reached the based case (index === s.length)
   And now index + i is 4
   then we do this line of code 
   const subStr = s.substring(index, index + i);
   const subStr = s.substring(2, 4);
   It's obvious that it's out of the bound, but slice function does NOT care
   It still gives us "3"
   Note that it's the same result of this line of code const subStr = s.substring(2, 3);
   So it will create duplicate element, which is not we want
   Try to get rid of this if-statement, and see the result


If we're given string "123", along with the prcoess
Use ("") to represent the subString of each call stack
use | to represent the position of index

                                                   |123
                       ("1") 1|23                                       ("12") 12|3
            ("2") 12|3            ("23") 123|                   ("3") 123|
        ("3") 123|

The result will be the sum of each path
"1", "2", "3"
"1", "23"
"12", "3"

It's very important to understand what's the meaning in this recursive tree?

For the top root node |123, it asks "what's the total decomposition amount of ths string "123"?"
We don't know the answer, but we do know we can decompose the first one digit and two digits if we start from the beginning
So the recusive functions starts
"Hey, I don't know the answer yet, so I ask the sub problem after decomposing the first two digits"
"I(root node) just wait here, and ask two sub problems. I'll get the answer once those two sub problems finish"
Sub Problem1: ("1") 1|23
Sub Problem2: ("12") 12|3

For the node ("1") 1|23, it's after decomposing one digit
This subproblem ask "what's the total decomposition amount of ths string "23"?"
Again, we don't know yet, so first decompose first tow digits, and ask the sub problem again
Sub Problem1: ("2") 12|3
Sub Problem2: ("23") 123|

For the node ("23") 123|, it's after decomposing two digits
This subproblem ask "what's the total decomposition amount of ths string ""?"
It hit's the base case, which means we get the answer, so throw the answer up to the previos call stack

So on and so forth.....


This tree leas us the time and space complexity
************************************************************
n is the length of string
Time: O(2 ^ n)
The depth of tree would be n deep long at worst, and brancing factor is 2

Space: O(n)


************************************************************
However, we can see the overlaping sub problem 12|3
So we can use memoization to optimize the solution

                                                   |123
                       ("1") 1|23                                       has answer already
            ("2") 12|3            ("23") 123|                   
        ("3") 123|

Time: O(n)
The depth of tree would be n deep long at worst

Space: O(n)
Keep the memo
*/
/**
 * @param {string} s
 * @return {number}
 */
function numDecodings(s) {
  return numDecodingsHelper(s, 0, {});
}

function numDecodingsHelper(s, index, memo) {
  if (memo[index] !== undefined) return memo[index];

  // finish decomposition, one possibility
  if (index >= s.length) return 1;

  let result = 0;

  // always decomposed two digits at max
  for (let i = 1; i <= 2; i++) {
    if (index + i <= s.length) {
      const subStr = s.substring(index, index + i);

      if (checkValidResult(subStr)) {
        result += numDecodingsHelper(s, index + i, memo);
      }
    }
  }

  memo[index] = result;

  return result;
}

function checkValidResult(str) {
  /*
    based on the description of this problem,
    "06" !== "6"
    so have to check if first character is "0"
    */
  if (str[0] === '0' || Number(str) >= 27 || Number(str) === 0) return false;

  return true;
}
/*
  This is bottom-up approach (tabulation) By myself
  
  For example, s = "1324"
  
  1. Initialize table with all 0 value
  
    1    3    2    4    ""
    0    0    0    0     0
  
  2. Base case is setting empty string to 1
     Because it's always decompose empty string
  
    1    3    2    4    ""
    0    0    0    0     1
  
  3. Based on the given constraint, set the last digit of string to 0 or 1
     If it's "0", set 0
     If not, set 1
     Because we can only map 1 ~ 26 to A ~ Z
     "0" is not allowed
  
    1    3    2    4    ""
    0    0    0    1     1
  
  4. Loop through from the last two digit to the first digit
     For each element, we can either decompose one digit or two digits
  
     for i = 2, the string is gonna decompose is "24"
     we can decompose "2" or "24"
     => decompose "2" (one digit)
        the subString after decompose "2" is "4", go check the cell of "4"(subProblem)
        Sub Problem => what's the total counts to decompose "4": 1
     => decompose "24" (two digits)
        the subString after decompose "24" is "", go check the cell of ""(subProblem)
        Sub Problem => what's the total counts to decompose "": 1
     => decompose one digit + decompose two digits = 2
  
      1    3    2    4    ""
      0    0    2    1     1
  
     for i = 1, the string is gonna decompose is "324"
     we can decompose "3" or "32"
     => decompose "3" (one digit)
        the subString after decompose "3" is "24", go check the cell of "24"(subProblem)
        Sub Problem => what's the total counts to decompose "24": 1 
        (go back to see above)(the string is gonna decompose is "24"?)(already solved the subProblem)
     => decompose "32" (two digits)
        Invalid: 0
     => decompose one digit + decompose two digits = 2
  
      1    3    2    4    ""
      0    2    2    1     1
  
     for i = 0, the string is gonna decompose is "1324"
     we can decompose "1" or "13"
     => decompose "1" (one digit)
        the subString after decompose "1" is "324", go check the cell of "324"(subProblem)
        Sub Problem => what's the total counts to decompose "324": 2
        (go back to see above)(already solved the subProblem)
     => decompose "24" (two digits)
        the subString after decompose "13" is "24", go check the cell of "24"(subProblem)
        Sub Problem => what's the total counts to decompose "24": 2
        (go back to see above)(already solved the subProblem)
     => decompose one digit + decompose two digits = 4
     
      1    3    2    4    ""
      4    2    2    1     1
  
                                                     |1324
                         ("1") 1|324                                       ("13") 13|24
              ("3") 13|24            ("32") 132|4                   ("2") 132|4        ("24") 1324|
          ("2") 132|4  ("24") 1324| ("4") 1324|                    ("4") 1324|
      ("4") 1324| 
      
      Decomposition => ["1", "3", "2", "4"], ["1", "3", "24"] ,["13", "2", "4"], ["13", "24"]
  
      Compare the table and recursive tree, one is top-down, the other one is bottom-up
      Recursive tree: Start from the beginning. |1324. "what's the total decomposition amount of ths string "1324"?". I don't know, but we can ask two subproblm after decomposing first two digits. So it follow asking "what's the total decomposition amount of ths string "324"?"(decompose "1") and "what's the total decomposition amount of ths string "24"?"(decompose "13"). Keep asking the subproblem till hit the base case which is index passing out the length of string or hits the empty string. So the question is asking from top to down
  
      Table: Start from the end. 1324. "what's the total decomposition amount of ths string "1324"?". I don't know yet, but let's start from the end which is empty string.
  
      Example:
      Table -> the string is gonna decompose is "24", the answer is 2
      Recusive tree -> go to see the node is 13|24, which means the same thing as the string is gonna decompose is "24"
                       it has two unique path
      
      Same thing, same logic as "324" (1|324)
  
      This is very important to review!!!
  
  ************************************************************
  s is the length of string
  Time: O(s)
  Iterate each cell of table, which length is equal to the length of given string
  
  Space: O(s)
  Construct DP table
  */
function numDecodings1(s) {
  const table = new Array(s.length + 1).fill(0);
  // empty string
  table[s.length] = 1;

  // last digit of string
  table[s.length - 1] = s[s.length - 1] === '0' ? 0 : 1;

  for (let i = s.length - 2; i >= 0; i--) {
    // decompose one digit and valid subString
    if (s[i] !== '0') table[i] = table[i + 1];

    // decompose two digits and valid subString
    if (s[i] === '1' || (s[i] === '2' && Number(s[i + 1]) <= 6))
      table[i] += table[i + 2];
  }

  return table[0];
}

/*
  This problem is very similar to fibonacci problem
  Which can be optimize to constant space
  
  Use variable, use same example as above
  
  twoDigitVar = 1, oneDigitVar = 1
  r = result, t = twoDigitVar, o = oneDigitVar
  
     for i = 2, 
                r    o    t
      1    3    2    4    ""
      0    0    2    1     1
  
     for i = 1, 
           r    o    t
      1    3    2    4    ""
      0    2    2    1     1
  
     for i = 0, 
      r    o    t
      1    3    2    4    ""
      4    2    2    1     1
  
  ************************************************************
  s is the length of string
  Time: O(s)
  Iterate each cell of table, which length is equal to the length of given string
  
  Space: O(1)
  */
function numDecodings2(s) {
  // empty string
  let twoDigitVar = 1;

  // last digit of string
  let oneDigitVar = s[s.length - 1] === '0' ? 0 : 1;

  let result = 0;

  for (let i = s.length - 2; i >= 0; i--) {
    // decompose one digit and valid subString
    if (s[i] !== '0') result += oneDigitVar;

    // decompose two digits and valid subString
    if (s[i] === '1' || (s[i] === '2' && Number(s[i + 1]) <= 6))
      result += twoDigitVar;

    // update varaible
    twoDigitVar = oneDigitVar;
    oneDigitVar = result;
    result = 0;
  }

  return oneDigitVar;
}
// console.log(numDecodings1('1324'));
// console.log(numDecodings2('1324'));
