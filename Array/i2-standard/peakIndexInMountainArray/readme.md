# Problem Statement (Peak Index in a Mountain Array)

You are given an integer mountain array arr of length n where the values increase to a peak element and then decrease.<br>
Return the index of the peak element.<br>
Your task is to solve it in O(log(n)) time complexity.<br>

Example 1:<br>
Input: arr = [0,1,0]<br>
Output: 1<br>

Example 2:<br>
Input: arr = [0,2,1,0]<br>
Output: 1<br>

Example 3:<br>
Input: arr = [0,10,5,2]<br>
Output: 1<br>


Constraints:
- 3 <= arr.length <= 105
- 0 <= arr[i] <= 106
- arr is guaranteed to be a mountain array.

# Solution Explanation
The idea of solution is very simple, let's see one example<br>
arr = [1,2,3,4,6,5,1]<br>
mid=3 (value is 4)<br>

It's obvious that index 3 is not the peak index, so we have to move either right or left pointer.<br>
Let's first compare arr[2]=3 (mid-1) and arr[3]=4 (mid).<br>
Because 4 > 3, that means it's increasing order. We know that the peak is NOT at the left side.
So we move the LEFT pointer to mid (left = mid), searching in the right half.<br>