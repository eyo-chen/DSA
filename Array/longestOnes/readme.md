# Problem Explanation

## Brute Force
The idea is to iterate through the array and for each element, we try to find the longest subarray that contains only 1s by flipping at most k 0s.<br>
For example, if we are at index i, we try to find the longest subarray that contains only 1s by flipping at most k 0s starting from index i.<br>
E.g. [1, 1, 0, 1], k = 1<br>
We start from index 0, we try to find the longest subarray that contains only 1s by flipping at most 1 zero.<br>
We then move to index 1, we try to find the longest subarray that contains only 1s by flipping at most 1 zero starting from index 1.<br>
So on and so forth.

### Complexity Analysis
#### Time Complexity O(n^2)
#### Space Complexity O(1)

## Sliding Window
The core idea is to maintain a window that contains at most k 0s and try to find the longest subarray that contains only 1s.<br>
We keep updating the right pointer and whenever we encounter a 0, we update the flip count.<br>
Then we ensure that the flip count does not exceed k by shrinking the window from the left.<br>
Finally, we update the answer with the current valid window size.<br>

Let's summarize the steps:
1. Initialize the left and right pointers to 0.
2. Initialize the flip count to 0.
3. Initialize the answer to 0.
4. Iterate through the array with the right pointer.
5. Whenever we encounter a 0, we update the flip count.
6. If the flip count exceeds k, we shrink the window from the left by moving the left pointer forward.
7. Update the answer with the current valid window size.

### Complexity Analysis
#### Time Complexity O(n)
#### Space Complexity O(1)
