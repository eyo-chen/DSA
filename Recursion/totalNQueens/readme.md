# Problem Explanation

The core idea of solution is similar to N Queens problems<br/>
The only difference is that we don't need to return the board, we just need to return the number of solutions<br/>

In this solution, we don't need `vector<string>` to store the board<br/>
Instead, we simply use `vector<int>` to store the column index of each row<br/>
For example,
[0, 1, 2, 3] means
<pre>
[0, 0, 1, 0]
[1, 0, 0, 0]
[0, 0, 0, 1]
[0, 1, 0, 0]
</pre>
The index of vector<int> means the row<br/>
The value of vector<int> means the column<br/>

By utilizing this `vector<int>`, we can know the position of each queen<br/>

Because changing the data structure, the isValid function is also changed<br/>
Look at the following board
<pre>
[0, 0, 0, 0]
[0, 1, 0, 0]
[0, 0, 0, 0]
[0, 0, 0, 1]
</pre>
Position [1, 1] and [3, 3] are at the same diagonal<br/>
How can we know the above position is valid or not?<br/>
We can use the following formula to check if it's the same diagonal<br/>
```
if (abs(rowDiff) == abs(colDiff))
```
abs(3 - 1) == abs(3 - 1) => true<br/>
If the above formula is true, it means the two positions are at the same diagonal<br/>
Note that we don't need to check the same row<br/>
Because different row is guaranteed by the vector<int> data structure<br/>

It also applies to the positive diagonal<br/>
<pre>
[0, 0, 0, 0]
[0, 0, 0, 1]
[0, 0, 0, 0]
[0, 1, 0, 0]
</pre>
Position [1, 3] and [3, 1] are at the same diagonal<br/>
abs(3 - 1) == abs(3 - 1) => true<br/>

## Choices and Constraints

- **Choice:** At any given row, we can place a queen at any column
  - Once we place a queen at r row, we can't place a queen at the same row
- **Constraint:** Can't place two queens at the same row, column, or diagonal
- **Goal:** When we reach the last row, we have a valid combination

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