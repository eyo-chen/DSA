//////////////////////////////////////////////////////
// *** Restore IP Addresses ***
//////////////////////////////////////////////////////
/*
A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order. 

Example 1:
Input: s = "25525511135"
Output: ["255.255.11.135","255.255.111.35"]

Example 2:
Input: s = "0000"
Output: ["0.0.0.0"]

Example 3:
Input: s = "101023"
Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

Constraints:
0 <= s.length <= 20
s consists of digits only.
*/
/*
I came up this solution by myself(second time)

This solution is actually intuitive
For each call stack, we always have three choices, decompose 1, 2 or three substr
Because the vaild ID address is always gonna between 0 ~ 255
For example, "123456789"
First choice, decompose one substr -> "1" ("1.23456789")
Second choice, decompose two substr -> "12" ("12.3456789")
Third choice, decompose three substr -> "123" ("123.456789")
We won't decompose four substr -> "1234" (X)
Because that's invalid ID address

After decomposing one substr, we can keep doing the same thing
"1.23456789"
First choice, decompose one substr -> "2" ("1.2.3456789")
Second choice, decompose two substr -> "23" ("1.23.456789")
Third choice, decompose three substr -> "234" ("1.234.56789")

so on and so forth

Another constraint to keep in mind is that
we can only have four segements in ID address
For example, 
"123.456.78.9" -> this is valid
"123.4.5.67.89" -> this is invalid
though each section is between 0 ~ 255, but the whole ID address only can have four segements

So that means we won't always need to decompose the whole string to the end
We only need to decompose three times to get four segements
See the example above, we only do two decomposition, but now we have three segments
("1.2.3456789")
("1.23.456789")
("1.234.56789")
=> All of them are three segements, and only do two decomposition

Now all it's become clear
Choice 
=> three decomposition, 1, 2 or 3 
Constraint
=> each segements should be between 0 ~ 255, and can't have leading zero
Goal 
=> decompose three times in total to get the four segments
=> this is usually our base case

                                        "245123888"
                   "2. ....."          "24. ....."         "245. ....."
 "2.4. ..."  "2.45. ..."  "2.451. ..."
so on and so forth

************************************************************
n = the legnth of s
Time complexity: O(1)
=> Branching factor -> 3
=> No matter how large or how long the input s is, we're always gonna have three choices
=> Again, because the constraint said 0 ~ 255
=> deepest height of recursive tree -> 4
=> we're always gonna have three recursive calls to have four segements
=> So right now all is O(1) works
=> Note that res.push(tmp.join('.')); and const subStr = s.slice(lenIndex, lenIndex + i); seems O(n) works
=> But if we really think about that, our longest ID address is gonna be 3 * 4 length long, three digits, and four segements
=> That means .join(".") and .slice() is gonna have O(1) work no matter how large input is
=> Note that O(1) doesn't mean it works less amount of time
=> It means no matter how big the input is, our operation is always running in constant time

Space complexity: O(1)
=> the deepest height of recursive tree is 4
*/
/**
 * @param {string} s
 * @return {string[]}
 */
var restoreIpAddresses = function (s) {
  // simple case when length is four
  if (s.length === 4) return [s.split('').join('.')];

  const res = [];

  recursiveHelper(s, 0, 0, [], res);

  return res;

  /*
  lenIndex -> keep tracking where should i keep decomposing (branching factor)
  heiIndex -> how many times i've decomposed (height of recursive tree)
  */
  function recursiveHelper(s, lenIndex, heiIndex, tmp, res) {
    // have decomposed three times (Goal)
    if (heiIndex === 4) {
      // if have fully decomposed the input
      if (lenIndex === s.length) {
        res.push(tmp.join('.'));
      }
      return;
    }

    // Choices
    for (let i = 1; i <= 3; i++) {
      // decompose (note tha index of .slice())
      const subStr = s.slice(lenIndex, lenIndex + i);

      // Constraint
      if (validAddress(subStr)) {
        tmp.push(subStr);
        recursiveHelper(s, lenIndex + i, heiIndex + 1, tmp, res);
        tmp.pop();
      }
    }

    return;
  }
};

function validAddress(str) {
  if (str.length > 3) return false;
  if (str.length > 1 && str[0] === '0') return false;
  if (Number(str) > 255) return false;

  return true;
}

/*
Another way to write the algorithm

Same complexity
*/
function restoreIpAddresses1(s) {
  if (s.length === 4) return [s.split('').join('.')];

  const res = [];

  recursiveHelper(s, 0, 0, '', res);

  return res;

  function recursiveHelper(s, lenIndex, heiIndex, tmp, res) {
    if (heiIndex === 4) {
      if (lenIndex === s.length) {
        // remove the last dot
        tmp = tmp.slice(0, tmp.length - 1);
        res.push(tmp);
      }

      return;
    }

    for (let i = lenIndex; i < lenIndex + 3; i++) {
      tmp += s[i];

      if (validAddress1(tmp)) {
        recursiveHelper(s, i + 1, heiIndex + 1, tmp + '.', res);
      }
    }

    return;
  }
}

function validAddress1(str) {
  let subStr = '';
  let index = str.length - 1;

  // if str is "123.125", we know that we've check "123", so now we only want to check "125"
  // so try to build "125"
  while (index >= 0) {
    if (str[index] === '.') break;
    subStr += str[index];

    index--;
  }

  // "521" -> "125"
  subStr = subStr.split('').reverse().join('');

  if (subStr.length > 3) return false;
  if (subStr.length > 1 && subStr[0] === '0') return false;
  if (Number(subStr) > 255) return false;

  return true;
}
console.log(restoreIpAddresses('1239130131'));
console.log(restoreIpAddresses1('1239130131'));
