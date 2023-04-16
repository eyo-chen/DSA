//////////////////////////////////////////////////////
// *** Move Element To End *** (algo)
//////////////////////////////////////////////////////
/*
You're given an array of integers and an integer
Write a function that moves all instances of that integer in the array to the end of the array and returns the array.
The function should perform this in place (i.e., it should mutate the input array) and doesn't need to maintain the order of the other integers.

Sample Input
array = [2, 1, 2, 2, 2, 3, 4, 2]
toMove = 2
Sample Output
[1, 3, 4, 2, 2, 2, 2, 2] // the numbers 1, 3, and 4 could be ordered differently
*/
/*
This solution works, but the code is bad and not cleanr
Go to see the second solution

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
function moveElementToEnd(array, toMove) {
  let index = array.length - 1;
  for (let i = 0; i < array.length; i++) {
    if (array[i] !== toMove) {
      continue;
    }

    while (index >= 0 && array[index] === toMove) {
      index--;
    }

    if (index < 0 || i >= index) {
      break;
    }

    const tmp = array[i];
    array[i] = array[index];
    array[index] = tmp;
    index--;
  }

  return array;
}

/*
Note that the main idea of this solution is as same as previous one
But this solution is more consice and cleaner
And it's easier to understand the main logic

We're using two pointers
left
=> I try to find the value toMove, so that I can swap with the right pointer
=> If I'm not a value toMove, I just keep going inward

right
=> I try to find the non-value toMove, so that I can swap with the left pointer
=> If I'm a value toMove, I just keep going inward

We just based on these two pointers to swap the value
Note that
We only do the swap operation when
1) the value of right pointer is not toMove
2) the value of left pointer is toMove
There's no sense to swap in other case

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
function moveElementToEnd(array, toMove) {
  let left = 0;
  let right = array.length - 1;

  while (left < right) {
    // if right is toMove, just go inward
    if (array[right] === toMove) {
      right--;
    }

    // if left is not toMove, just go inward
    if (array[left] !== toMove) {
      left++;
    }

    // swap
    if (array[left] === toMove && array[right] !== toMove) {
      const tmp = array[left];
      array[left] = array[right];
      array[right] = tmp;

      // also go inward at the same time
      right--;
      left++;
    }
  }

  return array;
}
