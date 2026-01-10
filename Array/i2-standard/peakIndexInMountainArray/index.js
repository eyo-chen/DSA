//////////////////////////////////////////////////////
// *** Peak Index in a Mountain Array ***
//////////////////////////////////////////////////////
/*
Let's call an array arr a mountain if the following properties hold:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
Given an integer array arr that is guaranteed to be a mountain, return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

Example 1:
Input: arr = [0,1,0]
Output: 1

Example 2:
Input: arr = [0,2,1,0]
Output: 1

Example 3:
Input: arr = [0,10,5,2]
Output: 1

Constraints:
3 <= arr.length <= 104
0 <= arr[i] <= 106
arr is guaranteed to be a mountain array.
 
Follow up: Finding the O(n) is straightforward, could you find an O(log(n)) solution?
*/
/**
 * @param {number[]} arr
 * @return {number}
 */
/*
Linear solution

************************************************************
Time compelxity: O(n)
Space comelxity: O(1)
*/
var peakIndexInMountainArray = function (arr) {
  for (let i = 0; i < arr.length - 1; i++) {
    if (arr[i] > arr[i + 1]) {
      return i;
    }
  }
};

/*
Using binary search logic

We have to first understand the pattern of mountain array
[1,2,5,4,3]
This is perfect mountain array
The first part [1,2] is increasing
The second part [4,3] is descending

In this case, mid = arr[2] = 5
arr[2] > arr[2 + 1]
5 > 4
What does this mean?
It means it's descending order
It guarantees that the peak is not in the [4,3]
So move right pointer to mid

mid = arr[1] = 2
arr[1] > arr[1 + 1]
2 > 5, false
What does this mean?
It means it's increasing order
It guarantees that the peak is now in the [5]
So move left pointer to mid + 1

I know it may be confused why right = mid and left = mid + 1
But I guess just have to use few test case to make the index right

Let's look at the other example to see why right = mid and left = mid + 1
nums = [1,3,5,30,4,2]

nums = [1,3,5,30,4,2]
left = 0, right = 5, mid = 2
nums[2] < nums[2 + 1]
left = mid + 1
Why we can say left = mid + 1?
Look at the condition
We compare mid & mid + 1
And we know the order of mid & mid + 1 is increasing
What does that mean?
It means mid is definitely NOT the peak
But mid + 1 could be peak
So we just move left to mid + 1

nums = [1,3,5,30,4,2]
left = 3, right = 5, mid = 4
nums[4] > nums[4 + 1]
right = mid
Why we can say right = mid?
Look at the condition
We compare mid & mid + 1
And we know that mid & mid + 1 is descending
What does that mean?
It mean mid + 1 is definitely NOT the peak
But mid could be the peak
So we just move right to mid

nums = [1,3,5,30,4,2]
left = 3, right = 4, mid = 3
nums[3] > nums[3 + 1]
Again, move right to mid

nums = [1,3,5,30,4,2]
left = 4, right = 4, 
DONE

Hope the logic is clear

************************************************************
Time compelxity: O(log n)
Space comelxity: O(1)
*/
var peakIndexInMountainArray = function (arr) {
  let right = arr.length - 1;
  let left = 0;

  while (right > left) {
    const mid = left + Math.trunc((right - left) / 2);

    if (arr[mid] > arr[mid + 1]) {
      right = mid;
    } else {
      left = mid + 1;
    }
  }

  return left;
};
