# 0-1 Knapsack Problem - Dynamic Programming Solution

## Problem Understanding

The 0-1 Knapsack Problem is a classic optimization problem where we need to select items to maximize value while staying within a weight constraint. The "0-1" means we can either take an item completely or leave it entirely - we cannot take fractions of items.

**Given:**
- An array of item values: `values[i]` represents the value of item `i`
- An array of item weights: `weights[i]` represents the weight of item `i`
- A maximum weight constraint: `maxWeight`

**Goal:** Find the subset of items that maximizes total value while keeping total weight ≤ maxWeight.

## Dynamic Programming Approach

Dynamic programming works exceptionally well for this problem because it exhibits two key properties: **optimal substructure** and **overlapping subproblems**.

### 1. Identifying Subproblems

The key insight is to think about the problem in terms of decisions we make for each item.<br>
Let's say the problem is given as:
```
values = [v1, v2, ..., vn]
weights = [w1, w2, ..., wn]
maxWeight = W(5)
```
It's really hard to know if we see all the items at once.<br>
Instead, we can start thinking the small cases first, like:
- What if we only had the first item(v1, w1) and maxWeight = 1?
- What if we had the first two items(v1, w1) and (v2, w2) with maxWeight = 2?

Something like this<br>
Because these smaller cases are easier to solve, we can build up to the full problem by combining solutions to these smaller cases.
And these are our subproblems.

At each step, we face a binary choice: **include the current item or exclude it**.<br>
We can break down the original problem into smaller subproblems:<br>
***"What is the maximum value we can get using only the current items(first `i` items) with the current max weight `w`?"***

This leads us to define our DP state as:
```
dp[i][w] = maximum value we can get using first i items with weight limit w
```

### 2. Understanding the Recurrence Relation

For each item and weight combination, we have two choices:

**Choice 1: Include the current item** (if it fits)
- Again, let's say the problem is given as:
```
values = [v1, v2, ..., vn]
weights = [w1, w2, ..., wn]
maxWeight = W(5)
```
- Suppose the current available item is (v2, w2) and we are at weight limit `w`.
- If we do choose this item, what's the maximum value we can get?
- we can think of as two parts:
  - The value of the current item: `v2` (since we are including it) 
  - The subproblem.
    - What's the subproblem?
    - The maximum value we can get after we choose this item.
    - AKA, the maximum value we can get using the first (v1, w1)(`i-1`) items with the remaining weight limit `w - w2`.
  - This is represented as `dp[i-1][w - current_weight] + current_value`
  - `i-1` represents the remaining items after including the current item.
    - If we initially had `n` items, after including the current item, we have `n-1` items left to consider.
  - `w - current_weight` represents the remaining weight limit after including the current item.

**Choice 2: Exclude the current item**
- If we do not choose the current item, the subproblem becomes:
  - What's the maximum value we can get using the first (v1, w1)(`i-1`) items with the same weight limit `w`?
  - This is represented as `dp[i-1][w]`
  - `i-1` represents the remaining items after excluding the current item.
  - `w` remains the same because we are not including the current item, so the weight limit does not change.

We take the maximum of these two choices

## Step-by-Step Walkthrough

Let's trace through the classic example:
```
values = [60, 50, 70, 30]
weights = [5, 3, 4, 2]
maxWeight = 8
```

### Setting Up the DP Table

We create a table with dimensions `(number of items + 1) × (maxWeight + 1)`. The extra row and column represent the base case of "no items" and "no weight capacity".

```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   ?   ?   ?   ?   ?   ?   ?   ?
[50, 3]  0   ?   ?   ?   ?   ?   ?   ?   ?
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```

**Why the extra row and column?** The empty array (first row) represents our base case where we have no items to choose from, so the maximum value is always 0. This prevents us from accessing negative indices when we look at `dp[i-1][w-weight[i-1]]`.

### Check the subproblem structure
```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   ?   ?   ?   ?   ?   ?   ?   ?
[50, 3]  0   ?   ?   ?   *   ?   ?   ?   ?
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```
When we're at `*`, we are trying to solve the subproblem of<br>
"What's the maximum value we can get using the first two items ([60, 5] and [50, 3]) with a weight limit of 3?"<br>


```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   ?   ?   ?   ?   ?   ?   ?   ?
[50, 3]  0   ?   ?   ?   ?   ?   ?   ?   ?
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   *   ?   ?
```
When we're at `*`, we are trying to solve the subproblem of<br>
"What's the maximum value we can get using the first two items ([60, 5], [50, 3] and [30, 2]) with a weight limit of 6?"

### Filling the Table - Item by Item

**Processing Item 1: [60, 5]**

```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0  60  60  60  60
[50, 3]  0   ?   ?   ?   ?   ?   ?   ?   ?
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```

For the first item with value 60 and weight 5:
- For weights 1-4: Cannot include the item (too heavy), so value remains 0
- For weights 5-8: Can include the item, so value becomes 60

Let's see an example:
```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0   *   ?   ?   ?
[50, 3]  0   ?   ?   ?   ?   ?   ?   ?   ?
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```
When we're at `*`, we are trying to solve the subproblem of<br>
"What's the maximum value we can get using the first item ([60, 5]) with a weight limit of 5?"<br>
Again, we have two choices:
1. We can choose the first item
   - The formula is `dp[i-1][w - current_weight] + current_value`
   - `dp[0][5 - 5] + 60 = 0 + 60 = 60`
   - For `dp[0][5-5]`, it basically means the subproblem after including the first item.
2. We can not choose the first item
   - The formula is `dp[i-1][w]`
   - `dp[0][5] = 0`
We take the maximum of these two choices: `max(60, 0) = 60`


**Processing Item 2: [50, 3]**

```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0  60  60  60  60
[50, 3]  0   0   0   *   *   ?   ?   ?   ?
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```
Suppose we are at `*`(3 and 4), we are trying to solve the subproblem of<br>
"What's the maximum value we can get using the first two items ([60, 5] and [50, 3]) with a weight limit of 3 and 4?"<br>
Again, we have two choices:
1. We can choose the current(second) item
   - The formula is `dp[i-1][w - current_weight] + current_value`
   - For weight 3: `dp[1][3 - 3] + 50 = 0 + 50 = 50`
   - For weight 4: `dp[1][4 - 3] + 50 = 0 + 50 = 50`
2. We can not choose the current(second) item
   - The formula is `dp[i-1][w]`
   - For weight 3: `dp[1][3] = 0`
   - For weight 4: `dp[1][4] = 0`

We take the maximum of these two choices:
- For weight 3: `max(50, 0) = 50`
- For weight 4: `max(50, 0) = 50`

```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0  60  60  60  60
[50, 3]  0   0   0   50  50  *   *   *   *
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```
Suppose we are at `*`(5, 6, 7 and 8), we are trying to solve the subproblem of<br>
"What's the maximum value we can get using the first two items ([60, 5] and [50, 3]) with a weight limit of 5, 6, 7 and 8?"<br>
Again, we have two choices:
1. We can choose the current(second) item
   - The formula is `dp[i-1][w - current_weight] + current_value`
   - For weight 5: `dp[1][5 - 3] + 50 = 0 + 50 = 50`
   - For weight 6: `dp[1][6 - 3] + 50 = 0 + 50 = 50`
   - For weight 7: `dp[1][7 - 3] + 50 = 0 + 50 = 50`
   - For weight 8: `dp[1][8 - 3] + 50 = 60 + 50 = 110`
2. We can not choose the current(second) item
   - The formula is `dp[i-1][w]`
   - For weight 5: `dp[1][5] = 60`
   - For weight 6: `dp[1][6] = 60`
   - For weight 7: `dp[1][7] = 60`
   - For weight 8: `dp[1][8] = 60`
We take the maximum of these two choices:
- For weight 5: `max(50, 60) = 60`
- For weight 6: `max(50, 60) = 60`
- For weight 7: `max(50, 60) = 60`
- For weight 8: `max(110, 60) = 110`

```
         0   1   2   3   4   5   6   7   8
[]       0   0   0   0   0   0   0   0   0
[60, 5]  0   0   0   0   0  60  60  60  60
[50, 3]  0   0   0   50  50 60  60  60 110
[70, 4]  0   ?   ?   ?   ?   ?   ?   ?   ?
[30, 2]  0   ?   ?   ?   ?   ?   ?   ?   ?
```


**The Pattern Emerges**

Notice how each cell depends on:
1. The cell directly above (excluding current item)
2. The cell at `dp[i-1][current_weight - current_item_weight]` plus current item's value (including current item)


## Complexity Analysis
**Time Complexity: O(n × w)**
- We have nested loops: `n` items × `w` weight capacities
- Each cell computation takes constant time
- Total operations: n × w

**Space Complexity: O(n × w)**
- We maintain a 2D DP table of size (n+1) × (w+1)
- Each cell stores one integer value

