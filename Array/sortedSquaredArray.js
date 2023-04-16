//////////////////////////////////////////////////////
// *** Sorted Squared Array ***
//////////////////////////////////////////////////////
/*
Write a function that takes in a non-empty array of integers that are sorted in ascending order and returns a new array of the same length with the squares of the original integers also sorted in ascending order.

Sample Input
array = [1, 2, 3, 5, 6, 8, 9]
Sample Output
[1, 4, 9, 25, 36, 64, 81]
*/
/*
Using Math.abs to sort the array first
Then do the caculation

************************************************************
Time compelxity: O(n * log(n))
Space comelxity: O(n)
*/
function sortedSquaredArray(array) {
  return array
    .sort((a, b) => Math.abs(a) - Math.abs(b))
    .map(element => element * element);
}

/*
Using two pointers
Because we know that the input array is sorted
So that we can make sure the largest number after caculation is at the two side of array
For example, [-5, -3, -1, 0, 5, 6]
=> After caculation, [25, 9, 1, 0, 25, 36]
=> As we can see, the first two largest value is 25 and 36
=> Which is the first and last element in the input array

So it's great to use two pointers
We just start at the first and last element
And using Math.abs to compare which value is greater
And first put the value in the output array

Note that we first deal with larger number
which means we have to build the array from the end to the beginning
It's okay
Just using index to build the array
************************************************************
Time compelxity: O(n)
Space comelxity: O(n)
*/
function sortedSquaredArray(array) {
  const output = new Array(array.length).fill(null);
  let left = 0;
  let right = array.length - 1;

  // build the index from the end
  let index = array.length - 1;

  while (right >= left) {
    const rightVal = array[right];
    const leftVal = array[left];

    // find the larger value
    if (Math.abs(rightVal) >= Math.abs(leftVal)) {
      output[index] = rightVal * rightVal;
      right--;
    } else {
      output[index] = leftVal * leftVal;
      left++;
    }

    // updating index
    index--;
  }

  return output;
}
