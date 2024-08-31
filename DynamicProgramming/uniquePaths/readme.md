# Problem Explanation

The critical part of this problem is to know that we can only move right or down.

## Brute Force Approach
At first, we can easily come up with a brute force approach by using recursion.<br>

The base case is <br>
- when the row and col is equal to the target row and col, we return 1.<br>
- when the row and col is out of the boundary, we return 0.<br>

The recursive case is to move right and down, and we add the result of the recursive call of moving right and moving down.<br>

Basically, for each recursive state, we're asking:<br>
How many ways can we get to this specific cell?<br>
- Go right to find the answer
- Go down to find the answer
- Sum the answer of going right and going down

Let's walk through an example to see how this works.<br>
```
 1  2  3
 4  5  6
 7  8  9
```
From 1 to 9, we have 6 ways to get there.
- 1 -> 4 -> 7 -> 8 -> 9
- 1 -> 4 -> 5 -> 8 -> 9
- 1 -> 4 -> 5 -> 6 -> 9
- 1 -> 2 -> 5 -> 8 -> 9
- 1 -> 2 -> 5 -> 6 -> 9
- 1 -> 2 -> 3 -> 6 -> 9

Look at the 5 this spot, we know that the unique path from 5 to 9 is 2.<br>
But we still need to duplicate calculate the unique path from 5 to 9 many times.<br>

Therefore, the problem with this approach is that it will cause a lot of duplicate calculation, and the calculation grows exponentially.<br>

### Complexity Analysis
#### Time Complexity O(2^(m+n))
- For each cell, we have 2 choices, either go right or go down.
- The branching factor is 2, and the depth of the tree is m+n.
- Therefore, the time complexity is O(2^(m+n)).


#### Space Complexity O(m+n)
- We need to store the recursive stack.
- The maximum depth of the recursive stack is m+n.
- Therefore, the space complexity is O(m+n).

## Optimized Approach
From the brute force approach, we can see that the duplicate calculation is the problem to the efficiency.<br>
So we can use a hash table to store the result of the unique path at a specific cell to the target cell.<br>

The key is to store the result of the unique path for each cell, so we don't need to recalculate it again.<br>

Let's walk through an example to see how this works.<br>
```
 1  2  3
 4  5  6
 7  8  9
```
First callstack
- memo = {}
- We're at 1. 
- And ask the question: How many ways can we get to 9 from 1?
- We don't know, go left and right to find the answer.

Second callstack
- memo = {}
- We're at 2.
- And ask the question: How many ways can we get to 9 from 2?
- We don't know, go left and right to find the answer.

Third callstack
- memo = {}
- We're at 5. (suppose we go down first)
- And ask the question: How many ways can we get to 9 from 5?
- We don't know, go left and right to find the answer.
- Suppose we know that the unique path from 5 to 9 is 2.
- We store it in the memo.
- memo = {
  "5": 2
}

....<br>

N callstack
- memo = {"5": 2}
- We're at 5 again. (Suppose we come from 4)
- And ask the question: How many ways can we get to 9 from 5?
- We already know the answer from the memo, which is 2.
- We return the answer.
- In this case, we don't need to go right or down, because we already know the answer from the memo.


### Complexity Analysis
#### Time Complexity O(m*n)
- For each cell, we only calculate it once.
- The maximum number of recursive calls is m*n.
- Therefore, the time complexity is O(m*n).

#### Space Complexity O(m*n)
- We need to store the recursive stack.
- The maximum depth of the recursive stack is m+n.
- Therefore, the space complexity is O(m*n).

## Dynamic Programming Approach
Above two approaches are kind of top-down approach, which is memoization.<br>
We can also use a bottom-up approach to solve this problem.<br>
Instead of building up the answer from the target cell to the start cell, we can find the answer from the start cell to the target cell.<br>

Let's see an example:
```
 1  2  3
 4  5  6
 7  8  9
```
In the previous approach, we start from start cell(1) and ask the question: How many ways can we get to 9 from 1?<br>
1 -> I don't know, let me ask 2.<br>
2 -> I don't know, let me ask 3.<br>
3 -> I don't know, let me ask 6.<br>
6 -> I don't know, let me ask 9.<br>
9 -> I'm here, so the answer is 1, let me pass it to 6.<br>
In this approach, we figure out the answer at the target cell(9), and pass it to the previous cell(6).<br>


However, we can figure out the answer at the start cell(1), and use it to build up the answer to the target cell(9).<br>
For [1, 2, 3] and [1, 4, 7], we know that there is only one way to get to each cell.
So we can initialize the table as:
```
 1  1  1
 1  0  0
 1  0  0
```
It represents that there's only one path to get to each cell from the start cell.

For 5, what's the number of ways to get to it from the start cell?<br>
The answer is 2.<br>
How do we get 2?<br>
The total number of ways to get to 5 is the sum of the number of ways to get to the cell above it and the cell left to it.<br>
So it's 1(from 1) + 1(from 2) = 2.<br>
We can fill the table as:
```
 1  1  1
 1  2  0
 1  0  0
```

Then, we can use the same logic to fill the rest of the table.
cell[r][c] = cell[r-1][c] + cell[r][c-1]

```
 1  1  1
 1  2  3
 1  3  6
```

When we using bottom-up approach, we need to figure out two things:<br>
- The base case
- The logic to solve the subproblem

In this example, we don't know what's the unique path from the start cell to the target cell.<br>
However, we know 
- Base case
  - The first row and first column only have one path to get to each cell.
- The logic to solve the subproblem
  - The number of ways to get to the current cell is the sum of the number of ways to get to the cell above it and the cell left to it.

We can use this logic to build up the final answer.<br>

### Complexity Analysis
#### Time Complexity O(m*n)
- We need to fill the table, which is m*n.
- Therefore, the time complexity is O(m*n).

#### Space Complexity O(m*n)
- We need to store the table.
- The space complexity is O(m*n).
