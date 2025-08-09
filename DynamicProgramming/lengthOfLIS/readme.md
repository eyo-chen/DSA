## Recursive Approach
The idea is that we brute force all possible subsequences of the input array `nums` and check if they are increasing subsequences.<br>
We use a index to track the current position in the array and an array to store the current subsequence.<br>
For each index `idx`, we have two choices:
1. Not choose the current element: We simply move to the next index without adding the current element to the subsequence.
2. Choose the current element: We need to check if the current element can be added to the subsequence (i.e., it is greater than the last element in the current subsequence). If it can, we add it to the subsequence and move to the next index.
Finally, we just return the maximum value between the two choices.<br>

### Complexity Analysis
#### Time Complexity O(2^n)
- At each index, the algorithm makes a binary decision: take the current element or skip it. This creates a binary tree of recursive calls where:
  - At index 0: 2 choices (take or skip)
  - At index 1: 2 choices for each previous choice = 4 total paths
  - At index 2: 2 choices for each of the 4 paths = 8 total paths
  - And so on...
- In the worst case, this leads to 2^n recursive calls.
```
For nums = [1, 2, 3]:
                    helper([], 0)
                   /            \
            helper([], 1)      helper([1], 1)
           /          \        /            \
    helper([], 2)  helper([2], 2)  helper([1], 2)  helper([1,2], 2)
      /    \        /      \         /        \         /         \
   len([]) len([3]) len([2]) len([2,3]) len([1]) len([1,3]) len([1,2]) len([1,2,3])
Each level doubles the number of calls, resulting in 2^n total calls.
```

#### Space Complexity O(n)

## Dynamic Programming Approach
The idea is create a 1D array `dp` where `dp[i]` represents the number of ways to reach the i-th step.<br>
For each `dp[i]`, it means what's the longest increasing subsequence that ends at step `i`, which is our subproblem.<br>

How can we find the answer for each subproblem?<br>
For each `dp[i]`, we can iterate through all previous steps `k` (where `k < i`) and check if the step `i` can be reached from step `k`. If it can, we can update `dp[i]` as follows:
```go
dp[i] = max(dp[i], dp[k] + 1)
```
It basically means that the answer for the subproblem `dp[i]` is the maximum of its current value and the value of `dp[k]` plus one (to account for the step `i` itself).<br>
It basically means that ***"Hey, I'm currently at step `i`, and I can extend the longest increasing subsequence from step `k`"***<br>
In other words, for each step `i`, it says ***"I need to check all the previous steps before me to see if I can extend the longest increasing subsequence from them, and I just get the maximum value from them"***

For example, nums=[0,1,0,3,2,3]<br>
- We first initialize the `dp` array with all elements set to 1, since each step can be reached in at least one way (by itself).
- For `i=0`
  - We intentionally skip this step since the first element is always a valid longest increasing subsequence.(Since there's no value before it)
  - dp=[1,1,1,1,1,1]
- For `i=1`, value is 1
  - It can be reached from step 0
  - dp=[1,2,1,1,1,1]
- For `i=2`, value is 0
  - It can't be reached from step 0 or 1 because it's value is 0
- For `i=3`, value is 3
  - It can be reached from step 0, 1 and 2
  - The maximum value from previous steps is 2 (from step 1), so we add 1 to it
  - dp=[1,2,1,3,1,1]
- For `i=4`, value is 2
  - It can be reached from step 0, 1 and 2
  - The maximum value from previous steps is 2 (from step 1), so we add 1 to it
  - dp=[1,2,1,3,3,1]
- For `i=5`, value is 3
  - It can be reached from step 0, 1, 2 and 4
  - The maximum value from previous steps is 3 (from step 4), so we add 1 to it
  - dp=[1,2,1,3,3,4]
Finally, the answer is the maximum value in the `dp` array, which is 4 in this case.<br>

### Complexity Analysis
#### Time Complexity O(2^n)
#### Space Complexity O(n)