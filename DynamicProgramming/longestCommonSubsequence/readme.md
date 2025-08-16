# Recursive Approach
Let's walk through an example with text1 = "abcde" and text2 = "ace"<br>
The question is: "What's the longest common subsequence of 'abcde' and 'ace'?"<br>
If we focus only on the last characters of both strings, "e" and "e"<br>
We immediately find one common character!<br>
So we can break this down into a subproblem:<br>
"What's the longest common subsequence of 'abcd' and 'ac'?"<br>
Once we get the answer, we simply add 1<br>

Now, let's solve this subproblem again<br>
What's the longest common subsequence of "abcd" and "ac"?<br>
If we focus only on the last characters of both strings, "d" and "c"<br>
They don't match<br>
So we know we can't use both of them<br>
We have two options:<br>
1. Ignore the last character of text1 and find the LCS of "abc" and "ac"<br>
2. Ignore the last character of text2 and find the LCS of "abcd" and "a"<br>
We take the maximum of both options<br>

That's the core idea behind solving this problem<br>
Let's summarize the approach:<br>
1. If the last characters of both strings match, we can include them in the LCS and move to the remaining characters.
2. If the last characters don't match, we have two options:
   - Ignore the last character of text1 and find the LCS of text1[:-1] and text2
   - Ignore the last character of text2 and find the LCS of text1 and text2[:-1]
   - We take the maximum of both options as the result for the current subproblem.

## Complexity Analysis
### Time Complexity O(2^(m+n))
- where m and n are the lengths of text1 and text2 respectively.
- In the worst case (when no characters match), each recursive call branches into two more calls. This creates a binary tree of recursive calls with a depth of (m+n), leading to exponential time complexity. For example, with strings of length 1000 each, we could have up to 2^2000 recursive calls, which is computationally infeasible.

### Space Complexity O(m+n)
- The maximum depth of recursion is (m+n) when we need to process all characters from both strings. Each recursive call uses constant space, so the total space complexity is proportional to the maximum call stack depth.

# DP Approach
The idea is very similar to the recursive approach, but instead of solving the same subproblems repeatedly, we build a DP table and fill each cell based on the results of subproblems we've already solved.<br>
Let's use the same example with text1 = "abcde" and text2 = "ace". The DP table looks like this:<br>
```
   ""  a   c   e
"" 0   0   0   0
a  0   0   0   0
b  0   0   0   0
c  0   0   0   0
d  0   0   0   0
e  0   0   0   0
```
The formula for filling each cell is as follows:<br>
- If the characters match, we take the diagonal value and add 1
  - `dp[r][c] = dp[r-1][c-1] + 1`
- If the characters don't match, we take the maximum of the left and top cells
  - `dp[r][c] = max(dp[r-1][c], dp[r][c-1])`

## Complexity Analysis
### Time Complexity O(m×n)
- where m and n are the lengths of text1 and text2 respectively.
- We need to fill every cell in the DP table exactly once. Since the table has (m+1) × (n+1) dimensions, and each cell takes constant time to compute, the total time complexity is O(m×n). This is a massive improvement over the exponential recursive approach.

### Space Complexity O(m×n)
- We create a DP table of size (m+1) × (n+1) to store the results of all subproblems. Each cell stores an integer representing the LCS length for the corresponding substrings. The space complexity can be optimized to O(min(m,n)) if we only keep track of the current and previous rows/columns, since we only need the diagonal, left, and top values to compute each cell.