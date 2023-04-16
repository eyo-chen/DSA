//////////////////////////////////////////////////////
// *** Validate Subsequence ***
//////////////////////////////////////////////////////
/* 
Given two non-empty arrays of integers, write a function that determines whether the second array is a subsequence of the first one.
A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array but that are in the same order as they appear in the array. 
For instance, the numbers [1, 3, 4] form a subsequence ofthe array [1, 2, 3, 4] ,
and so do the numbers  [2, 4].
Note that a single number in an array and the array itself are both valid subsequences of the array.

Sample Input
array = [5, 1, 22, 25, 6, -1, 8, 10]
sequence = [1, 6, -1, 10]
*/
/*
This problem is fairly easy
It's kind of like using two pointers

We iterate the sequence array
Our goal is finding all the element of sequence array in the array array
For example,
i = 0, val = 1
=> using index to find the val 1 in the array
=> index = 1, val = 1
=> index++

i = 1, val = 6
=> using index to find the val 6 in the array
=> keep updating the index until find the val 6
=> index = 4, val = 6
=> index++

i = 2, val = -1
=> using index to find the val 6 in the array
=> keep updating the index until find the val 6
=> index = 5, val = -1
=> index++

i = 3, val = 10
=> using index to find the val 6 in the array
=> keep updating the index until find the val 6
=> index = 7, val = 10
=> index++

After iterating sequence, just return true,
which means we find all the element of sequence in the array

Another example, 
array = [5, 1, 22, 25, 6, -1, 8, 10]
sequence = [10, 1]
i = 0, val = 10
=> using index to find the val 6 in the array
=> keep updating the index until find the val 6
=> index = 7, val = 10
=> index++

i = 1, val = 1
=> index nos is 8 which is the length of array
=> index++
=> index > array.length, so return false
=> It means we've looped through all the array, and we can't find the element of sequence

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
function isValidSubsequence(array, sequence) {
  let index = 0;

  for (let i = 0; i < sequence.length; i++) {
    while (index < array.length && array[index] !== sequence[i]) {
      index++;
    }

    index++;

    if (index > array.length) {
      return false;
    }
  }

  return true;
}
