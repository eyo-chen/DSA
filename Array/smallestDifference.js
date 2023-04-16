//////////////////////////////////////////////////////
// *** Smallest Difference ***
//////////////////////////////////////////////////////
/*
Write a function that takes in two non-empty arrays of integers, finds the pair of numbers (one from each array) whose absolute difference is closest to zero, and returns an array containing these two numbers, with the number from the first array in the first position.
Note that the absolute difference of two integers is the distance between them on the real number line. For example, the absolute difference of -5 and 5 is 10, and the absolute difference of -5 and -4 is 1.
You can assume that there will only be one pair of numbers with the smallest
difference.

Sample Input
arrayOne = [-1, 5, 10, 20, 28, 3]
arrayTwo = [26, 134, 135, 15, 17]

Sample Output
[28, 26]
*/
/*
************************************************************
n = arrayOne.length, m = arrayTwo.length

Time compelxity: O(n * m)
Space comelxity: O(1)
*/
function smallestDifference(arrayOne, arrayTwo) {
  let min = Infinity;
  let output = null;
  for (let i = 0; i < arrayOne.length; i++) {
    const num1 = arrayOne[i];
    for (let j = 0; j < arrayTwo.length; j++) {
      const num2 = arrayTwo[j];

      if (Math.abs(num1 - num2) < min) {
        output = [num1, num2];
        min = Math.abs(num1 - num2);
      }
    }
  }

  return output;
}

/*
************************************************************
n = arrayOne.length, m = arrayTwo.length

Time compelxity: O(nlog(n) + mlog(m))
Space comelxity: O(1)
*/
function smallestDifference(arrayOne, arrayTwo) {
  let min = Infinity;
  let output = null;
  let index1 = 0,
    index2 = 0;

  arrayOne.sort((a, b) => a - b);
  arrayTwo.sort((a, b) => a - b);

  while (index1 < arrayOne.length && index2 < arrayTwo.length) {
    const num1 = arrayOne[index1];
    const num2 = arrayTwo[index2];

    if (Math.abs(num1 - num2) < min) {
      min = Math.abs(num1 - num2);
      output = [num1, num2];
    }

    if (num1 > num2) {
      index2++;
    } else {
      index1++;
    }
  }

  return output;
}
