//////////////////////////////////////////////////////
// *** Backspace String Compare ***
//////////////////////////////////////////////////////
/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

Example 1:
Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
 
Constraints:
1 <= s.length, t.length <= 200
s and t only contain lowercase letters and '#' characters.
 
Follow up: Can you solve it in O(n) time and O(1) space?
*/
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
/*
This is my first solution
The idea is simple, just seperately generate the edtied string

Let's focus on the generateString function
the key is using skip variable to note us when do we need to skip the element
We loop the string from the end to the beginning
If it's not "#" and skip is greater than 0
It means it's time to skip the current string
If it's "#", we just simply increment the skip varaible
else, just add the character to ouput string

For example, s = "a#cde##h"
i = 7, s[7] = h, skip = 0
=> just add "h"
output = "h"

i = 6, s[6] = #, skip = 1
=> increment skip variable
output = "h"

i = 5, s[5] = #, skip = 2
=> increment skip variable again
output = "h"

i = 4, s[4] = e, skip = 2
=> current character is not "#" and skip is greater than 0
=> skip the current character
output = "h"

i = 3, s[3] = d, skip = 1
=> current character is not "#" and skip is greater than 0
=> skip the current character
output = "h"

i = 2, s[2] = c, skip = 0
=> just add "c"
output = "hc"

i = 1, s[1] = #, skip = 1
=> increment skip variable
output = "hc"

i = 0, s[0] = a, skip = 1
=> current character is not "#" and skip is greater than 0
=> skip the current character
output = "hc"


The idea of this solution is simple, but it's not the most optimize solution

************************************************************
n = s.length, m = t.length
Time complexity: O(n + m)
Space complexity: O(n + m)
*/
var backspaceCompare = function (s, t) {
  return generateString(s) === generateString(t);
};

function generateString(s) {
  let output = '';
  let skip = 0;

  for (let i = s.length - 1; i >= 0; i--) {
    // If it's not "#" and skip is greater than 0, just skip current character
    if (s[i] !== '#' && skip > 0) {
      skip--;
      continue;
    }

    // increment the skip variable
    if (s[i] === '#') {
      skip++;
    }
    // add the current character
    // it's important to add the character in this way
    // because we iterate the string from the end to the beginning
    else {
      output = s[i] + output;
    }
  }

  return output;
}

/*
This is another solution I reference from YT
The idea is very similar to previous solution
But this solution using stack to generate edtied string

Using stack in the problem is great
If we hit the "#", we just remove the last added character

After generating two stack, we just compare them

************************************************************
n = s.length, m = t.length
Time complexity: O(n + m)
Space complexity: O(n + m)
*/
var backspaceCompare = function (s, t) {
  const newS = generateStack(s);
  const newT = generateStack(t);

  if (newS.length !== newT.length) {
    return false;
  }

  for (let i = 0; i < newS.length; i++) {
    if (newS[i] !== newT[i]) {
      return false;
    }
  }

  return true;
};

function generateStack(str) {
  const stack = [];

  for (const s of str) {
    if (s === '#') {
      stack.pop();
    } else {
      stack.push(s);
    }
  }

  return stack;
}

/*
I wrote the solution after know that it has to use two pointer

Again, we loop through the string from the end to the beginning

The main idea of this solution is this part of code
        if (s[ptrS] === '#') {
        skipS++;

        while (s[--ptrS] === '#' || skipS !== 0) {
            if (s[ptrS] !== '#') {
            skipS--;
            } else {
            skipS++;
            }
        }
        }

Let's see in detail
If s[ptrS] is "#", we have to skip the later character
first increment the skip variable
Go inside the while-loop
At the time going inside the while-loop, we have to immediately minus ptrS
Because we've already know s[ptrS] is "#"
(I know there are other way to write the code, but this way is great)
If s[ptrS] is not "#", we just minus the skip variable
else, we want to keep increment skip varible
Note that we'll keep minusing ptrS variable when while-loop condition check

The main idea is 
If we encounter "#", we want to skip later character
And we use skip variable to help us to know how many characters are gonna skip

Note the code above is exactly the same as this one
        if(s[ptrS] === "#"){
            skipS++;
            ptrS--;
            
            while(s[ptrS] === "#" || skipS !== 0){
                if(s[ptrS] !== "#"){
                    skipS--;
                }else{
                    skipS++;
                }
                ptrS--;
            }
        }

************************************************************
n = s.length, m = t.length
Time complexity: O(Max(n, m))
Space complexity: O(1)
*/
var backspaceCompare = function (s, t) {
  let ptrS = s.length - 1;
  let ptrT = t.length - 1;
  let skipS = 0;
  let skipT = 0;

  // keep iterating until both of ptr over the first index
  while (ptrS >= 0 || ptrT >= 0) {
    // gonna skip later character
    if (s[ptrS] === '#') {
      skipS++;

      while (s[--ptrS] === '#' || skipS !== 0) {
        if (s[ptrS] !== '#') {
          skipS--;
        } else {
          skipS++;
        }
      }
    }

    // gonna skip later character
    if (t[ptrT] === '#') {
      skipT++;

      while (t[--ptrT] === '#' || skipT !== 0) {
        if (t[ptrT] !== '#') {
          skipT--;
        } else {
          skipT++;
        }
      }
    }

    // after skipping, if character is not the same, return false
    if (s[ptrS] !== t[ptrT]) {
      return false;
    }

    // this is testing their length
    // if one of pointer is finish iterating, and the other one is still keep iterating
    // we know that the length of both valid string is not the same
    // just return false
    if (ptrS >= 0 !== ptrT >= 0) {
      return false;
    }

    ptrS--;
    ptrT--;
  }

  return true;
};
