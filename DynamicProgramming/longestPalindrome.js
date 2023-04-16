var longestPalindrome = function (s) {
  return recursiveHelper(s, 0, s.length - 1);

  function recursiveHelper(s, left, right, index) {
    if (left > right) return 0;
    else if (left === right) {
      if (index) return 1;
      else return 0;
    }

    if (s[left] === s[right])
      return recursiveHelper(s, left + 1, right - 1, true) + 2;
    else {
      return Math.max(
        recursiveHelper(s, left + 1, right, false),
        recursiveHelper(s, left, right - 1, false)
      );
    }
  }
};

function longestPalindrome1(s) {
  return recursiveHelper(s, 0, s.length - 1);

  function recursiveHelper(s, left, right) {
    if (left > right) return '';
    else if (left === right) {
      return s[left];
    }

    let middle, goLeft, goRight;
    if (s[left] === s[right]) {
      middle = s[left] + recursiveHelper(s, left + 1, right - 1) + s[right];
    } else {
      goLeft = recursiveHelper(s, left + 1, right, false);
      goRight = recursiveHelper(s, left, right - 1, false);
    }

    if (middle) return middle;
    else if (goLeft.length > goRight.length) return goLeft;
    else return goRight;
  }
}

// console.log(longestPalindrome1('cabdac'));

function longestPalindrome2(s) {
  let res = '';

  for (let i = 0; i < s.length; i++) {
    for (let j = i + 1; j <= s.length; j++) {
      const subStr = s.slice(i, j);
      if (checkPalindrome(subStr) && subStr.length > res.length) {
        res = subStr;
      }
    }
  }

  return res;
}

function checkPalindrome(s) {
  let left = 0,
    right = s.length - 1;

  while (left < right) {
    if (s[left] !== s[right]) return false;
    left++;
    right--;
  }

  return true;
}

function longestPalindrome3(s) {
  let start = 0,
    end = 0;

  for (let i = 0; i < s.length; i++) {
    const len1 = expandOutward(s, i, i);
    const len2 = expandOutward(s, i, i + 1);

    const tmpMaxLen = Math.max(len1, len2);

    if (tmpMaxLen > end - start) {
      start = i - Math.trunc((tmpMaxLen - 1) / 2);
      end = i + tmpMaxLen / 2;
    }
  }

  return s.slice(start, end + 1);
}

function expandOutward(s, left, right) {
  if (left > right) return 0;

  while (left >= 0 && right < s.length && s[left] === s[right]) {
    left--;
    right++;
  }

  return right - left - 1;
}

// console.log(longestPalindrome3('cbbd'));
