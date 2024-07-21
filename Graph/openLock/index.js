//////////////////////////////////////////////////////
// *** Open the Lock ***
//////////////////////////////////////////////////////
/*
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

Example 1:
Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation: 
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".

Example 2:
Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation: We can turn the last wheel in reverse to move from "0000" -> "0009".

Example 3:
Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation: We cannot reach the target without getting stuck.

Constraints:
1 <= deadends.length <= 500
deadends[i].length == 4
target.length == 4
target will not be in the list deadends.
target and deadends[i] consist of digits only.
*/
/**
 * @param {string[]} deadends
 * @param {string} target
 * @return {number}
 */
/*
Using DFS, queue

This problem is not that hard, the main idea is still using DFS
The harder part is how to update the next string
For example,
current string is "0000"
It can connect with
"1000", "9000", "0100", "0900", "0010", "0090", "0001", "0009"
The hardest part is to find these updated string

The idea is 
1. loop through four digits
2. and pass string, index(digit) and queue to helper function
It will help us to generate correct updated string
For example,
updateStr("0000", 1, queue)
After executing, queue = ['0900', '0100']

updateStr("0009", 3, queue)
After executing, queue = ['0008']

updateStr("1111", 2, queue)
After executing, queue = ['1121', '1101']

After that, we just directly pass into queue
And 
      if (seen[str] || hashTable[str]) {
        continue;
      }

      if (str === target) {
        return count;
      }
will help us to catch the base case
*/
var openLock = function (deadends, target) {
  // convert deadends array to object, so that it's O(1) access
  const hashTable = deadends.reduce((acc, cur) => {
    acc[cur] = true;
    return acc;
  }, {});

  const seen = {};

  // initial value is "0000" because problem said always start at "0000"
  const queue = ['0000'];
  let count = 0;

  while (queue.length > 0) {
    let len = queue.length;

    while (len > 0) {
      const str = queue.shift();

      len--;

      // if it has been seen or it's deadend, just skip it
      if (seen[str] || hashTable[str]) {
        continue;
      }

      // find the target
      if (str === target) {
        return count;
      }

      // has been seen
      seen[str] = true;

      // update string according to each digit
      for (let i = 0; i < 4; i++) {
        updateStr(str, i, queue);
      }
    }

    count++;
  }

  return -1;
};

/*
If current digit is "0", then we
1. set to "9"
2. set to "1"

If current digit is "9", then we
1. set to "9 - 1"

Other case, then we
1. set to "digit + 1"
2. set to "digit - 1"
*/
function updateStr(str, index, queue) {
  const arr = [...str];
  const char = str[index];

  switch (arr[index]) {
    case '0': {
      arr[index] = '9';
      queue.push(arr.join(''));

      arr[index] = '1';
      queue.push(arr.join(''));

      break;
    }

    case '9': {
      arr[index] = String(arr[index] - 1);
      queue.push(arr.join(''));

      break;
    }

    default: {
      arr[index] = String(+arr[index] + 1);
      queue.push(arr.join(''));

      // recover back to orginal character
      arr[index] = char;

      arr[index] = String(arr[index] - 1);
      queue.push(arr.join(''));

      break;
    }
  }
}

/*
Another way to update the string
For example,
"0000", index = 1
(0 + 1) % 10 = 1
((0 - 1) + 10) % 10 = 9

"0900", index = 1
(9 + 1) % 10 = 0
((9 - 1) + 10) % 10 = 8

As we can see,
If it's +, we can use + 1, then % 10, so that if it's 9 + 10, we also get 0, which is what we want
If it's -, we can use + 10, then % 10, so that if 0 - 1, we can also get (-1 + 10) % 10 = 9, which is what we want to
*/
function updateStr1(str, index, queue) {
  const arr = [...str];
  const char = str[index];

  arr[index] = String((+arr[index] + 1) % 10);
  queue.push(arr.join(''));

  arr[index] = char;

  arr[index] = String((arr[index] - 1 + 10) % 10);
  queue.push(arr.join(''));
}
