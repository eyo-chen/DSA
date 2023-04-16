//////////////////////////////////////////////////////
// *** Additive Number ***
//////////////////////////////////////////////////////
/*
An additive number is a string whose digits can form an additive sequence.

A valid additive sequence should contain at least three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

Given a string containing only digits, return true if it is an additive number or false otherwise.

Note: Numbers in the additive sequence cannot have leading zeros, so sequence 1, 2, 03 or 1, 02, 3 is invalid.

Example 1:
Input: "112358"
Output: true
Explanation: 
The digits can form an additive sequence: 1, 1, 2, 3, 5, 8. 
1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

Example 2:
Input: "199100199"
Output: true
Explanation: 
The additive sequence is: 1, 99, 100, 199. 
1 + 99 = 100, 99 + 100 = 199
 
Constraints:
1 <= num.length <= 35
num consists only of digits.
*/
/*
This problem is kinda harder than others

One thing to make the problem more clear
If str is "1235"
Once choosing "1" and "2", then we have to let "2" be the first in next recursive calls
In other words, once first + seond = third
In next recursive calls, 
second has to be first, third has to be second
so on and so forth
That's the Additive Number

This solution is kinda like bruth force
We just decompose the whole string to find the all the possible first and second
For example, str = "12345",
first "1", second "2", subStr "345"
first "1", second "23", subStr "45"
first "1", second "234", subStr "5"
first "1", second "2345", subStr ""
first "12", second "3", subStr "45"
first "12", second "34", subStr "5"
first "12", second "345", subStr ""
first "123", second "4", subStr "5"
first "123", second "45", subStr ""
first "1234", second "5", subStr ""
We try all of these possibilities, and do the recursion on each of these possibilities

Along with the process, we need to check first, second, subStr have leading zero
We won't go on recursion if one of them has leading zero

Also, another tricky part is .slice() method, have to care about the index like i and j, so that we can have all correct decomposition
*/
/**
 * @param {string} num
 * @return {boolean}
 */
function isAdditiveNumber(str) {
  // simple case
  if (str.length <= 2) return false;
  if (str.length === 3)
    return Number(str[0]) + Number(str[1]) === Number(str[2]);

  // outer loop find the first number
  for (let i = 1; i <= str.length; i++) {
    const first = str.slice(0, i);

    // skip any possibilities with leadingZero
    if (checkLeadingZero(first)) continue;

    // inner loop find the second number
    for (let j = i + 1; j <= str.length; j++) {
      const second = str.slice(i, j);

      // skip any possibilities with leadingZero
      if (checkLeadingZero(second)) continue;

      // remainder is the subString after substract first and second
      const remainder = str.slice(j);

      if (checkLeadingZero(remainder)) continue;

      /*
      If remainder is "", then we know this possbility is no way being additive number
      For example, str = "12345"
      First "123", Second "45", Remainder ""
      This is no way to be additive number
      */
      if (remainder.length > 0 && recursiveHelper(first, second, remainder))
        return true;
    }
  }

  return false;
}

function recursiveHelper(first, second, subStr) {
  // base case
  if (subStr === '') return true;

  for (let i = 1; i <= subStr.length; i++) {
    /*
    Given the subString(aka, remainder), find any third subString(number) is matching the formula
    first + second = third
    For example, str = "12358"
    First "1", Second "2", subStr "358"
    Try to find third subString in the "358" is matching the 1 + 2
    */
    const third = subStr.slice(0, i);

    /*
    Instead of using plus, using subtraction to avoid overflow for very large input integers
    This is kinda important
    */
    const validIndex = Number(third) - Number(second) === Number(first);

    // also need to check remainder have no leading zero
    const remainder = subStr.slice(i);
    if (checkLeadingZero(remainder)) continue;

    /*
    As we can see here,
    second become first in the next recursive call stack
    third become second
    */
    if (validIndex && recursiveHelper(second, third, remainder)) return true;
  }

  return false;
}

function checkLeadingZero(str) {
  if (str.length > 1 && str[0] === '0') return true;
  return false;
}

console.log(isAdditiveNumber('011235'));
