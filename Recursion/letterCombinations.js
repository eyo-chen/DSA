//////////////////////////////////////////////////////
// *** Letter Combinations of a Phone Number ***
//////////////////////////////////////////////////////
/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]
 
Constraints:
0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
*/
/*
I came up this solution by myself after watching the instruction video only in 1 minutes

The point got me to this solution is 
"illustrate this problem as decision space"
For example, digits = "23", we can think sth like this
the final output is just having two digits space
2 3
_ _  -> decision space
a d  ->
b e  -> these three are possible decisions 
c f  ->

For each slot or decision space, we know we can use some logic to fill up those slots

Recursive Tree
Start at first index, "2", we know it can have "abc"
                               "23"
                "a"            "b"            "c"
First index "2" can have three decisions
Then, each decisions can have another more decisions
                               "23"
                "a"            "b"            "c"
  "ab"    "ad"   "ae"  "bd"   "ba"  "be"   "cd"  "ca"  "ce"

************************************************************
n = the legnth of digits
Time complexity: O(4 ^ n)
=> Closely look the hashTable, then we know that the maximum branching factor is just 4
=> Which means for each single digit, we'll explore four decisions at most, so the iteration in each call stack is just O(1) works
=> It means the works will not scall with the input grow
=> The deepest height of recursive tree is O(n)
=> If the length of input digit is 4, then we'll have 4 height of recursive tree because we have to fill up four slots

=> Note that the first solution has res.push(tmp.join('')); 
=> It may be O(n * 4 ^ n) if we consider .join() do O(n) work

Space complexity: O(n)
*/
/**
 * @param {string} digits
 * @return {string[]}
 */
const hashTable = [
  '', // 0
  '', // 1
  'abc', // 2
  'def', // 3
  'ghi', // 4
  'jkl', // 5
  'mno', // 6
  'pqrs', // 7
  'tuv', // 8
  'wxyz', // 9
];

var letterCombinations = function (digits) {
  // this is how the question defined when the length of input is 0
  if (digits.length === 0) return [];
  const res = [];
  recursiveHelper(digits, 0, [], res);

  return res;

  function recursiveHelper(digits, index, tmp, res) {
    /*
    Base case
    If input digit is 2
    _ _ -> we have two slots, and it's zero base, so if now index is 2, then we know we out of the bound
    */
    if (index === digits.length) {
      res.push(tmp.join(''));
      return;
    }

    // map to corrsponding character
    const charArr = hashTable[Number(digits[index])];

    // fill up decisions space
    for (let i = 0; i < charArr.length; i++) {
      tmp.push(charArr[i]);
      recursiveHelper(digits, index + 1, tmp, res);
      tmp.pop();
    }

    return;
  }
};

/*
Same logic
But do NOT use array, use string instead
*/
function letterCombinations1(digits) {
  if (digits.length === 0) return [];
  const res = [];
  recursiveHelper(digits, 0, '', res);

  return res;

  function recursiveHelper(digits, index, tmp, res) {
    if (index === digits.length) {
      res.push(tmp);
      return;
    }

    const charArr = hashTable[Number(digits[index])];

    for (let i = 0; i < charArr.length; i++) {
      recursiveHelper(digits, index + 1, tmp + charArr[i], res);
    }

    return;
  }
}

/*
This is completely different solution
It's not using the recursion backtracking
Instead, using the concept of queue

For example, digits = "23"

1st outer loop (loop through "23")
=> "2" -> "abc"
=> remeber the length of current result array(1)

1st mid loop (loop through current result array)("")
=> pre = ""

1st inner loop (loop through "abc")
=> add "" + "a" at the beginning
=> add "" + "b" at the beginning
=> add "" + "c" at the beginning

2nd outer loop
=> "3" -> "def"
=> remeber the length of current result array(3)

2nd mid loop (loop through current result array)(["a", "b", "c"])
=> pre = "a"

2nd inner loop (loop through "def")
=> add "a" + "d" at the beginning
=> add "a" + "e" at the beginning
=> add "a" + "f" at the beginning

2nd inner loop (loop through "def")
=> add "b" + "d" at the beginning
=> add "b" + "e" at the beginning
=> add "b" + "f" at the beginning

2nd inner loop (loop through "def")
=> add "c" + "d" at the beginning
=> add "c" + "e" at the beginning
=> add "c" + "f" at the beginning

As we could see, this is the main idea of this solution
There are three things to note
1. Add empty string in the res array at the very biginning
=> Why?
=> This is kina like the default case
=> The core idea is poping out the prev element in the res array, and then add new character
=> We have to make sure the very first poping out element is at least a string, so that we can add another new string
=> const prev = res.pop(); and  res.unshift(prev + strArr[k]);
=> aslo const resLen = res.length; need the res has to be one length

2. const resLen = res.length; is necessary and important
=> Why?
=> Because inside inner loop, we will keep adding new added element in the result array
=> which means the length of result array is not fixed, it's dynamic
=> it will keep increasing, so we have to reemeber the current length of array
=> for example, the input digits = "23"
=> Now we're at the "3"
=> So now the res array shoule be ["a", "b", "c"]
=> Now we have to remeber the length of current array
=> which means we only can pop the array three times later
=> ["cd", "a", "b"]
=> ["ce" ,"cd", "a", "b"]
=> ["cf", "ce" ,"cd", "a", "b"]
=> ["bd" ,"cf", "ce" ,"cd", "a"]
=> so on and so forth

3. This is not the most efficent way to implement
=> array.pop() is okay
=> res.unshift(prev + strArr[k]); is not okay, it will take O(n) underneath the hood
=> In order to have both O(1) time to add at the begining, and remove at the end is
=> Link-List
=> Here just temporarily use array
=> But can say it would be better if using linklist in interview

************************************************************
n = the legnth of digits
Time complexity: O(n * (3 ^ n))
=> the outer loop is O(n)
=> the mid loop should be O(3 ^ n)
=> Why?
=> If input digits = "234"
=> First mid loop -> res = ["a", "b", "c"]
=> Second mid loop -> res = ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
=> As we could see, mid loop will scale as the prcoess keep going
=> not sure this is correct
=> the inner loop is always O(1) times, either 3 or 4 

Space complexity: O(1)
*/
var letterCombinations2 = function (digits) {
  // base case
  if (digits.length === 0) return [];
  const res = [''];

  // O(n)
  for (let i = 0; i < digits.length; i++) {
    const strArr = hashTable[Number(digits[i])];
    const resLen = res.length;

    // O(3 ^ n)
    for (let j = 0; j < resLen; j++) {
      const prev = res.pop();

      // O(1)
      for (let k = 0; k < strArr.length; k++) {
        res.unshift(prev + strArr[k]);
      }
    }
  }

  return res;
};
console.log(letterCombinations2('3456').length);
// console.log(letterCombinations1('31'));
