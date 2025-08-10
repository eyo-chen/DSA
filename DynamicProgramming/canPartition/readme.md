# Brute Force Solution (Recursion)
The idea is quite simple.<br>
We can try to find all the possible subsets, and check if the sum of the subset is equal to the target sum.<br>
If we find a subset with the target sum, we return true.<br>
If we exhaust all subsets and do not find any, we return false.<br>

What's the critiria for a subset to be valid?<br>
If the `total_sum - current_sum == current_sum`, then we have a valid subset.<br>
It basically means `total_sum = 2 * current_sum`.<br> 

How to find all subsets?<br>
We can use recursion to find all subsets.<br>
We can either include the current element in the subset or not.<br>
If we include the current element, we add it to the current sum and move to the next element.<br>
If we do not include the current element, we move to the next element without adding it to the current sum.<br>

Let's visualize this with an example:<br>
nums=[1,5,11,5]<br>
sum=22<br>
```
                                               0
                                1                             0              <- choose 1 or not
                      6              1               5               0       <- choose 5 or not
               17          6    12       1      16      5       11       0   <- choose 11 or not
```
For the current sum 11, we found that `22 - 11 == 11`, so we return true.<br>

## Complexity Analysis
### Time Complexity O(2^n)
- where n is the number of elements
- We explore 2 choices (include/exclude) for each of the n elements
- In the worst case, we explore all 2^n possible subsets
### Space Complexity O(n)
- Maximum depth of recursion is n (when we process all elements)


# Memoization Solution
The idea is to use recursion with memoization to avoid recomputing the same subproblems.<br>
But what exactly are the subproblems?<br>
Before we answer that, let's shift our mindset a bit.<br>
Back to our formula before, `total_sum - current_sum == current_sum`.<br>
We can rewrite it as `total_sum == 2 * current_sum`.<br>
Also, we can rewrite `total_sum / 2 == current_sum`.<br>
And we can see the `current_sum` as our target sum.<br>
So, we can think of the problem as ***finding a subset of the array that sums up to `total_sum / 2`***.<br>
If we can find such a subset, then we can partition the array into two subsets with equal sum.<br>

Now, let's back to the diagram to understand the subproblems.<br>
```
                                               0
                                1(*)                             0                  <- choose 1 or not
                      6(**)              1               5               0          <- choose 5 or not
               17          6    12       1      16      5       11       0          <- choose 11 or not
```
At the * position, after choosing 1, the problem becomes finding a subset of the remaining elements(5,11,5) that sums up to `total_sum / 2 - 1 = 10`.<br>
At the ** position, after choosing 5, the problem becomes finding a subset of the remaining elements(11,5) that sums up to `total_sum / 2 - 5 = 5`.<br>
Can you see the pattern?<br>
The subproblem is to ***find a subset of the remaining elements that sums up to the new target sum***.<br>
We have two variables to track the state of the problem:<br>
1. The index of the current element we are considering.
   - AKA, the remaining elements.
   - For example, if we're * position, we're at index 0, so the remaining elements are [5,11,5].
2. The target sum we are trying to achieve.<br>

So, we can use these two variables as the key for our memoization.<br>
We can use a map to store the results of the subproblems.<br>
If we have already computed the result for a given key, we can return it directly.<br>
If not, we can compute the result and store it in the map for future reference.<br>

## Complexity Analysis
### Time Complexity O(n × sum)
- where n is array length and sum is total sum
- We have at most n × (sum/2) unique subproblems
- Each subproblem is computed only once and cached

### Space Complexity O(n × sum)
- O(n × sum) for memoization table + O(n) for recursion stack

# Iterative Solution
The idea is different from dynamic programming solutions for other problems.<br>
We basically try to find all the possible sums, and stores them in a set(hashTable).<br>
If we find a sum that is equal to `total_sum / 2`, we return true.<br>
If we exhaust all elements and do not find any, we return false.<br>

How to find all possible sums?<br>
The logic is following:
1. We start with a set with a single element 0, which represents the sum of an empty subset.
2. For each element in the array, we iterate through the current sums in the set and
   - Add the current element to each sum in the set.
   - Add the new sums to the set.

For example, for the array [1,5,11,5], we can visualize the process as follows:
```
Initial set: {0}
After adding 1: {0, 1}
After adding 5: {0, 1, 5(0+5), 6(1+5)}
After adding 11: {0, 1, 5, 6, 11(0+11), 12(1+11), 16(5+11)}
....
```
If we find `total_sum / 2` in the set, we return true.<br>

## Complexity Analysis
### Time Complexity O(n × sum) 
- where n is array length and sum is total sum
- For each of the n numbers, we iterate through all current possible sums
- Maximum number of possible sums is sum/2

### Space Complexity O(sum)
- In worst case, we might store up to sum/2 different sums

# Dynamic Programming
This is basically a 0/1 knapsack problem.<br>
We can use dynamic programming to solve this problem.<br>
Before explaining the solution, go to check the knapsack problem to understand the basic idea.<br>
Because this problem is very similar to the knapsack problem, we can use the same approach.<br>

The subproblem is the same as the memoization solution.<br>
The subproblem is to ***find a subset of the remaining elements that sums up to the new target sum***.<br>

Suppose the problem is given as:<br>
nums = [1, 5, 11, 5]<br>
We know that the total sum is 22, so the target sum is 11.<br>
So, the problem is to find a subset of the array that sums up to 11.<br>
We can define our DP state as:<br>
`dp[i][j] represents if we can find a subset of the first i elements that sums up to j`<br>
- `i` represents the index of the current element we are considering.<br>
- `j` represents the target sum we are trying to achieve.<br>

dp[1][5] represents if we can find a subset in [1] that sums up to 5.<br>
dp[2][6] represents if we can find a subset in [1, 5] that sums up to 6.<br>
....<br>
dp[4][11] represents if we can find a subset in [1, 5, 11, 5] that sums up to 11.<br>
And this is the final answer we are looking for.<br>


The initial DP table looks like this:
```
                0   1   2   3   4   5   6   7   8   9   10  11
[]              t   f   f   f   f   f   f   f   f   f   f   f
[1]             t   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?
[1, 5]          t   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?
[1, 5, 11]      t   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?
[1, 5, 11, 5]   t   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?   ?
```
The first row(except the first column) is all false because we can't find a subset that sums up to any positive number with an empty array.<br>
The first column is all true because we can always find a subset that sums up to 0 with any array (the empty subset).<br>

For each cell `dp[i][j]`, we have two choices:
1. **Include the current item**:
   - If we include the current item, we check if we can find a subset of the first `i-1` items that sums up to `j - current_item`.
   - The formula is `dp[i-1][j - current_item]`.
   - For example, if we are at `dp[2][6]`, the problem is "Can we find a subset of the first 2 items([1, 5]) that sums up to 6?"<br>
   - If we include the current item, we check if we can find a subset of the first 1 item([1]) that sums up to `6 - 5 = 1`.
   - That's the answer of `dp[1][1]`
2. **Exclude the current item**:
   - If we exclude the current item, we check if we can find a subset of the first `i-1` items that sums up to `j`.
   - The formula is `dp[i-1][j]`.
   - For example, if we are at `dp[2][6]`, the problem is "Can we find a subset of the first 2 items([1, 5]) that sums up to 6?"<br>
   - If we exclude the current item, we check if we can find a subset of the first 1 item([1]) that sums up to `6`.
   - That's the answer of `dp[1][6]`  

We take the logical OR of these two choices to fill the DP table:
```
dp[i][j] = dp[i-1][j - current_item] || dp[i-1][j]
```

However, we need to handle the case where `j - current_item` is negative.<br>
If the current item is greater than `j`, we cannot include it<br>

## Complexity Analysis
### Time Complexity O(n × sum)
### Space Complexity O(n × sum)