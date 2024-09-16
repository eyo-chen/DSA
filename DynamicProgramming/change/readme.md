# Problem Explanation

## Brute-force Approach
The brute-force approach is to use recursion to solve the problem.<br>
For every recursive call, it represents a subproblem:<br>
"How many ways can I make up the current amount using the current coins?"<br>

There is on note:<br>
We have to avoid the duplicate combination

Suppose we have the coins: [1,2,5], and the amount is 5.<br>
These two combinations are the same:
- [1,2,2]
- [2,1,2]

We want to avoid count these two combinations as different.<br>
The approach to avoid that is to use index to limit the coins we can use.<br>
Once we choose a coin, we can only use the coins after it.<br>
For example, if we choose coin 2, we can only use coin 2 and 5.<br>
If we choose coin 5, we can only use coin 5.

### Complexity Analysis
#### Time Complexity O(n^m)
- n is the number of coin denominations
- m is the amount we're trying to make change for
- The depth of the recursion tree can go up to m (in the worst case, if we have a coin of value 1).
- At each level, we can potentially use any of the n coins.

#### Space Complexity O(m)
- where m is the amount we're trying to make change for.
- The recursion stack. In the worst case, the recursion can go as deep as the amount we're trying to make change for (if we have a coin of value 1 and keep subtracting 1 from the amount).

## Top-down Approach
This is basically the same as the brute-force approach, but we use memoization to avoid the duplicate calculation.<br>

There's one thing to note:<br>
We have to use amount and index as key to avoid the duplicate calculation.<br>

We can't simply use the amount as the key<br>
Here's an example to illustrate:<br>
Let's say we have coins [1, 2, 5] and we're trying to make an amount of 5.<br>
If we only used amount as the key, we might memoize the result for amount 5 after considering all coins.<br>
But the number of ways to make 5 using all coins [1, 2, 5] is different from the number of ways to make 5 using only [2, 5] or only [5].<br>
By including the index in the key, we're able to distinguish between these different scenarios:<br>
"5-0" represents making 5 using all coins [1, 2, 5]<br>
"5-1" represents making 5 using only [2, 5]<br>
"5-2" represents making 5 using only [5]<br>
This approach ensures that we correctly count all unique combinations without duplicates, which is essential for solving this particular variation of the coin change problem.<br>

### Complexity Analysis
#### Time Complexity O(n*m)
- n is the number of coin denominations
- m is the amount we're trying to make change for
- Each subproblem (unique combination of amount and index) is computed only once and then stored in the memo.
- There are m * n possible subproblems (m possible amounts * n possible starting indices).
- Each subproblem takes O(n) time to compute (due to the for loop over coins).

#### Space Complexity O(n * m)
- The memoization map (memo) can store up to m * n entries in the worst case.
- The recursion stack can go up to depth m in the worst case.
- The space used by the map dominates, so the overall space complexity is O(n * m).

## Bottom-up Approach
The bottom-up approach is to build up the solution from the smallest subproblems to the largest problem.<br>
We first create a 2D array like following:
```
0   1   2   3   4   5
[]  1   0   0   0   0   
[1] 1   0   0   0   0   
[2] 1   0   0   0   0   
[5] 1   0   0   0   0   
```
The row is the coin index, and the column is the amount.<br>
Each cell represents the subproblem:<br>
"How many ways can I make up the current amount using the current coins?"<br>

For example,<br>
```
0   1   2   3   4   5
[]  1   0   0   0   0   
[1] 1   0   0   0   0   
[2] 1   0   *   0   0   
[5] 1   0   0   0   0   
```
`*` means "How many ways can I make up the amount 3 using the coin [1,2]?"<br>

```
0   1   2   3   4   5
[]  1   0   0   0   0   
[1] 1   0   0   0   0   
[2] 1   0   0   0   0   
[5] 1   0   0   *   0   
```
`*` means "How many ways can I make up the amount 4 using the coin [1,2,5]?"<br>

```
0   1   2   3   4   5
[]  1   0   0   0   0   
[1] 1   0   0   0   0   
[2] 1   0   0   0   0   
[5] 1   0   0   0   *   
```
`*` means "How many ways can I make up the amount 5 using the coin [1,2,5]?"<br>
This is our final answer, and all other cells are the subproblems.<br>

How can we fill in the cell `*`?<br>
For each subproblem, we have two choices:<br>
1. Do not use the current coin
   - If we do not use the current coin, we can use the previous coins to make up the amount. So the number of ways to make up the amount is the same as the number of ways to make up the amount without the current coin.
   - doNotUseAmount = table[i-1][j]
2. Use the current coin
   - If we use the current coin, we need to make up the amount minus the current coin. So the number of ways to make up the amount is the same as the number of ways to make up the amount minus the current coin.
   - useAmount = table[i][j-coins[i]]

For example,<br>
```
0   1   2   3   4   5
[]  1   0   0   0   0   
[1] 1   0   ^   0   0   
[2] -   0   *   0   0   
[5] 1   0   0   0   0   
```
The total amount of `*` is the sum of `table[i-1][j]` and `table[i][j-coins[i]]`.
- `table[i-1][j]` is represented by `^`
   - It means "Because I don't use coin 2, so how many ways can I make up the amount 3 using the coin [1]?"
- `table[i][j-coins[i]]` is represented by `-`
   - It means "Because I use coin 2, so how many ways can I make up the amount 1 using the coin [1,2]?"
So `*` = `^` + `-`

### Complexity Analysis
#### Time Complexity O(n*m)
- n is the number of coin denominations
- m is the amount we're trying to make change for
- We have two nested loops:
   - The outer loop iterates over all coins (n iterations)
   - The inner loop iterates over all amounts from 1 to m (m iterations)
- Inside these loops, we perform constant time operations.

#### Space Complexity O(n * m)
- We create a 2D table of size (n+1) * (m+1).
- No additional space that grows with input size is used.