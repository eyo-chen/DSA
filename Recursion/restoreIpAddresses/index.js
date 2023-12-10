/*
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
