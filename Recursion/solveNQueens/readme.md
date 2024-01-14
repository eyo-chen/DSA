# Problem Explanation

The idea to solve this problem is to try any possible combination of the queens' positions, and check if the combination is valid or not.

Let's walk through the thought process.<br/>
At the very beginning, place queens [0, 0] <br/>

What's the next placement?<br/>
[0, 1] => No. It's obvious that two queens are at the same row<br/>
**That's the first intuition**<br/>
We don't need to try any possible position at the same row since it's obvious violation<br/>

Okay, so let's move on to the next row<br/>
[1, 0] => No. It's obvious that two queens are at the same column<br/>
This check is straightforward<br/>

Okay, let's try on next position within the same row<br/>
[1, 1] => No. Because it's the same diagonal<br/>
How can we check if it's the same diagonal?<br/>
Let's look the example below
<pre>
[0, 0] [0, 1] [0, 2] [0, 3] -> row 0
[1, 0] [1, 1] [1, 2] [1, 3] -> row 1
[2, 0] [2, 1] [2, 2] [2, 3] -> row 2
[3, 0] [3, 1] [3, 2] [3, 3] -> row 3
</pre>
Suppose we're at [2, 2] (row 2)<br/>
Look at the previous row(row 1), [1, 1], [1, 3] is at the same diagonal<br/>
Look at the next previous row(row 0), [0, 0] is at the same diagonal<br/>
It's the negative diagonal<br/>
When [r - rowDiff, r - rowDiff] <br/>
It's the positive diagonal<br/>
When [r - rowDiff, r + rowDiff] <br/>
**That's the second intuition**<br/>

Note that there are multiple ways to check the constraint<br/>
Here is just one of the ways<br/>
Lookt at the second solution of Javascript<br/>
Thought the checking logic is different, the core logic is the same<br/>

## Choices and Constraints

- **Choice:** At any given row, we can place a queen at any column
  - Once we place a queen at r row, we can't place a queen at the same row
- **Constraint:** Can't place two queens at the same row, column, or diagonal
- **Goal:** When we reach the last row, we have a valid combination

## Recursive Tree Visualization
<pre>
                                        top
                    (0,0)                    (0,1)           (0,2)           (0,3)
        (1,0)   (1,1)   (1,1)   (1,2)  
(2,0)(2,1)(2,2)(2,3)
</pre>
It's abvious we won't have perfect tree like this, which means the branching factor won't always be 4, because we have the constraint, can't be same row and col, also diagonal cases <br/>
It's just a visualization to help us understand the problem


# Complexity Analysis

n = the legnth of input string

## Time Complexity: O((n^n) * n)
- Branching Factor = n
  - At worst, we can choose n times
- Depth = n
  - At worst, we have n rows
- Each call stack = O(n)
    - Because of the palindrome check

This is the very rough estimation, because we have the constraint, can't be same row and col, also diagonal cases <br/>
So the branching factor won't always be n, it's less than n<br/>
Another time complexity analysis is O(n! * n)<br/>
First row, we have n choices<br/>
Second row, we can have at least n - 2 choices

## Space Complexity: O(n)

## Detail Explanation
To understand the time complexity of this solution, let's examine the recursive function helper which is essentially a backtracking algorithm to solve the N Queens problem.

For the N Queens problem, a brute force approach would generate all possible combinations of queens on the board, which would be O(n^n) since for every row there are n possibilities, and there are n rows to consider. However, the backtracking approach tries to place a queen in a valid position row by row and skips entire sub-trees of invalid board states, which means it does not explore all n^n possibilities.

Analyzing backtracking time complexities can be tricky because the algorithm doesn't explore all possible paths; it prunes many paths that lead to invalid solutions. The worst-case time complexity is difficult to determine, but it is significantly less than O(n^n).

The number of solutions for N Queens grows exponentially with N, but it's not as simple as O(2^n) or O(n!). The exact number of solutions is not known for arbitrary n, but there are established results for small values of n.

The isValid function checks three conditions for each row (up to the current one), which takes O(n) time. Since isValid is called for every cell in a row before the recursive call, the time complexity for a single row is O(n^2). However, this does not simply translate to O(n^3) for the entire board because of the backtracking pruning.

Without an exact formula for the number of valid arrangements of queens, we can't provide a precise time complexity. However, a common way to express the complexity of the N Queens backtracking algorithm is O(n!), recognizing that the algorithm has to potentially explore every permutation of row placements (even though many permutations are pruned).

For space complexity, the algorithm requires O(n) space for the temp vector (which stores the current state of the board), and the recursive call stack could go up to O(n) deep in the case of placing a queen in each row. Therefore, the space complexity is O(n).

In summary, the time complexity is less than O(n^n) and often estimated as O(n!) due to backtracking pruning, while the space complexity is O(n) due to the board and the recursive call stack.