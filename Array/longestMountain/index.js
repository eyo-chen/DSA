//////////////////////////////////////////////////////
// *** Longest Mountain in Array ***
//////////////////////////////////////////////////////
/*
You may recall that an array arr is a mountain array if and only if:

arr.length >= 3
There exists some index i (0-indexed) with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given an integer array arr, return the length of the longest subarray, which is a mountain. Return 0 if there is no mountain subarray.

Example 1:
Input: arr = [2,1,4,7,3,2,5]
Output: 5
Explanation: The largest mountain is [1,4,7,3,2] which has length 5.

Example 2:
Input: arr = [2,2,2]
Output: 0
Explanation: There is no mountain.
 
Constraints:
1 <= arr.length <= 104
0 <= arr[i] <= 104
 
Follow up:
Can you solve it using only one pass?
Can you solve it in O(1) space?
*/
/**
 * @param {number[]} arr
 * @return {number}
 */
/*
The main idea of this solution is using two arrays
up
=> represent the length of going up
[2,1,4,7,3,2,5]
[0,0,1,2,0,0,0]
=> 0 means I don't go up
=> 1 means I go up one length
=> 2 means I go up two length
=> Note that we finish this array from the beginning to the end

down
=> represent the length of going down
[2,1,4,7,3,2,5]
[0,0,0,2,1,0,0]
=> Basically same thing as above
=> Note that we finish this array from the end to the beginning

After that, if two arrays are not 0, then we caculate the length

************************************************************
Time complexity: O(n)
Space complexity: O(n)
*/
var longestMountain = function (arr) {
  const up = new Array(arr.length).fill(0);
  const down = new Array(arr.length).fill(0);
  let output = 0;

  for (let i = 1; i < arr.length - 1; i++) {
    if (arr[i] > arr[i - 1]) {
      up[i] = up[i - 1] + 1;
    }
  }

  for (let i = arr.length - 2; i > 0; i--) {
    if (arr[i] > arr[i + 1]) {
      down[i] = down[i + 1] + 1;
    }
  }

  for (let i = 1; i < arr.length - 1; i++) {
    if (up[i] && down[i]) {
      output = Math.max(output, up[i] + down[i] + 1);
    }
  }

  return output;
};

/*
This is optimize solution
This problem is kind of hard

One edge case to note that is 
flat state connot be count as part of mountain
For example, arr = [1,2,3,3,3,3,1,2], output = 0
Tho there's a increasing order and descending order
But there's no peak (one is highest value), so it's not a mountain
So a mountain has to be stricly increasing order and descending order
Like this one, arr = [1,2,3,4,3,2,1], output = 7

This solution is I reference from discuss
This is the most intutive one, I can't understant other solutions....

The main idea of this solution is very straightforward
1. Keep going up, just like climbing the mountain
2. check if start === index
   => What does this mean?
   => Note that we set start = index before the while-loop(going up)
   => If start is still equalt to index
   => It means we never go up, which also means arr[index] === arr[index + 1]
   => It's flat, not the part of mountain
   => so we just update index, and skip the later execution
   => Because we know this part is never in our output

2. set peak = index
   => In correct case, now index should be the highest value, which is peak
   => And now we should immediately go down

3. Keep doing down
4. Check peak === index
   => This logic is as same as above
   => If peak is equal to index
   => I means we never move, which means arr[index] === arr[index + 1]
   => Again, we want to skip it because it's never part of the mountain

5. Finally, use index - start + 1 to count the length of mountain
   => Noe that index is just like end index

************************************************************
Time complexity: O(n)
Space complexity: O(1)
*/
var longestMountain = function (arr) {
  const len = arr.length;
  let index = 0;
  let output = 0;

  while (index < len) {
    // set starting point
    let start = index;

    // going up
    while (index + 1 < len && arr[index] < arr[index + 1]) {
      index++;
    }

    // we never go up, just skip it and update the index
    if (start === index) {
      index++;
      continue;
    }

    // find the peak
    let peak = index;

    // go down immediately
    while (index + 1 < len && arr[index] > arr[index + 1]) {
      index++;
    }

    // we never go up, just skip it and update the index
    if (peak === index) {
      index++;
      continue;
    }

    // count the length of mountain
    output = Math.max(output, index - start + 1);
  }

  return output;
};
