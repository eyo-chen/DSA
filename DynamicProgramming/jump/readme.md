# Problem Explanation

## Top-Down Approach(Recursion + Memoization)
The idea to solve this problem is very similar to the `canJump` problem.

Here's the thought process:<br>
Suppose the input is [2,3,1,1,4]<br>
For the fist index, we have two choices, jump 1 step to index 1, or jump 2 steps to index 2.<br>
For jump 1 step, the subproble becomes `jump([3,1,1,4])`, and for jump 2 steps, the subproble becomes `jump([1,1,4])`.<br>
When we got the answer from `jump([3,1,1,4])` and `jump([1,1,4])`, we just choose the smaller one and plus 1.

Base case:
- If the current index is the last index, return 0.
  - return 0 because we are at the last index, no need to jump.
  - and the caller will plus 1 to the result.
- If the current index is 0, return -1.

Recursive case:
- Iterate through the possible steps from the current index.
- If the helper function returns a valid result, return the result plus 1.
- If no steps lead to a valid result, return -1.

### Complexity Analysis
#### Time Complexity: O(n^2)
- In the worst case, the function `helper` can be called for each index of the array.
- For each call, we might iterate through all the possible steps (up to the length of the array in the worst case).
- This leads to a nested loop structure, resulting in O(n^2) time complexity.

#### Space Complexity: O(n)
- The `memo` slice is created with the same length as the input array, which is O(n).
- The recursion stack in the worst case (when we need to jump to each element one by one) can go up to n levels deep.

## Greedy Approach
The idea of greedy approach is kind of break down the input array into different intervals.<br>
The total number of intervals is the minimum number of jumps required to reach the end of the array.<br>

Let's take an example to understand this approach:<br>
Input: [2,3,1,5,3,2,1]

1. We're at the index 0, where max steps we can take is 2.
   - rightPtr = 0 + 2 = 2, In this case, the interval is [0, 2].
   - jumps = 1, which means the first interval is first jump(jump from index 0 to index 2)
   - leftPtr = rightPtr = 2
2. We're at the index 1, where max steps we can take is 3.
   - rightPtr = max(rightPtr, 1 + 3) = 4. Now, we only move the rightPtr to 4.
   - i != leftPtr, it means that we're still in the first interval.
   - so we don't need to increase the jumps, and we don't need to update the leftPtr.
   - the reason we need to update the rightPtr is that, we want to make sure that we're covering the maximum possible steps during every interval.
3. We're at the index 2, where max steps we can take is 1.
   - rightPtr = max(rightPtr, 2 + 1) = 4.
   - i == leftPtr, it means that we're at the end of the first interval.
   - we need to increase the jumps by 1, which means we move to the next interval.
   - leftPtr = rightPtr = 4, updating the leftPtr to the start of the next interval.
4. We're at index 3, where max steps we can take is 5.
   - rightPtr = max(rightPtr, 3 + 5) = 8.
   - i != leftPtr, it means that we're still in the second interval.
   - so we don't need to increase the jumps, and we don't need to update the leftPtr.
5. We're at index 4, where max steps we can take is 3.
   - rightPtr = max(rightPtr, 4 + 3) = 8.
   - i == leftPtr, it means that we're at the end of the second interval.
   - we need to increase the jumps by 1, which means we move to the next interval.
   - leftPtr = rightPtr = 8, updating the leftPtr to the start of the next interval.
   - leftPtr >= n-1, so we can break the loop.
   - it means that the start of the current new interval is already beyond the last index, so we don't need to continue the loop.

- return jumps

### Complexity Analysis
#### Time Complexity: O(n)
#### Space Complexity: O(1)
  
