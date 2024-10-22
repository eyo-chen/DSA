# Problem Description

Given an array A of integers and integer K, return the maximum S such that there exists i < j with A[i] + A[j] = S and S < K. If no i, j exist satisfying this equation, return -1.<br>

Example 1:<br>
Input: A = [34,23,1,24,75,33,54,8], K = 60<br>
Output: 58<br>
Explanation: <br>
We can use 34 and 24 to sum 58 which is less than 60.<br>

Example 2:<br>
Input: A = [10,20,30], K = 15<br>
Output: -1<br>
Explanation: <br>
In this case it's not possible to get a pair sum less that 15.<br>

Note:<br>
- 1 <= A.length <= 100
- 1 <= A[i] <= 1000
- 1 <= K <= 2000

# Problem Explanation
The solution is very similar to two sum sorted problem.<br>
We first sort the array, then use two pointers to find the maximum sum that is less than K.<br>
If the current sum is less than K, we update the maximum sum and move the left pointer to the right, so that the sum is increased.<br>
If the current sum is greater than or equal to K, we move the right pointer to the left, so that the sum is decreased.<br>

# Complexity Analysis
## Time Complexity O(n*log(n))
## Space Complexity O(1)