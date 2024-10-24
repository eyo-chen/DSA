# Problem Explanation

## Using Hash Table
The idea is pretty simple, we first create a direction array to represent the direction of the spiral.<br>
Then we iterate through the matrix, and for each element, we check if the next element is out of bound or has been visited.<br>
If it is, we change the direction.<br>
We also use a hash table to record the visited elements to avoid revisiting.<br>

For example,<br>
```
1  2  3  4  x
5  6  7  8  
9 10 11 12
```
We start from 1, and the initial direction is right.<br>
We go right until we hit 'X' which means that we're out of bound.<br>
So, we change the next direction to down.<br>
Again, we go down until we're out of bound.<br>
Then we change the direction to left.<br>
After hitting 9, we change the direction to up.<br>
We'll stop at 1 because we've already visited it.<br>

### Complexity Analysis
#### Time Complexity O(m*n)
- We iterate through the matrix once, so the time complexity is O(m*n).

#### Space Complexity O(m*n)
- We use a hash table to record the visited elements, so the space complexity is O(m*n).


## Using Variable to Control Direction
The idea is to use four variables to control the boundary of the matrix.<br>
- rowStart: The start row index.
- rowEnd: The end row index.
- colStart: The start column index.
- colEnd: The end column index.

We start from the top-left corner of the matrix, and we go right until we hit the `colEnd`.<br>
After that, we need to move the `rowStart` down by 1, which means the first row is done exploring.<br>
Then we go down until we hit the `rowEnd`.<br>
After that, we need to move the `colEnd` left by 1, which means the last column is done exploring.<br>
Then we go left until we hit the `colStart`.<br>
After that, we need to move the `rowEnd` up by 1, which means the last row is done exploring.<br>
Then we go up until we hit the `rowStart`.<br>
After that, we need to move the `colStart` right by 1, which means the first column is done exploring.<br>
We repeat the above process until we've explored the entire matrix.

For example,
```
1  2  3  4
5  6  7  8  
9 10 11 12
```
We start from 1, and the initial boundary is `rowStart = 0`, `rowEnd = 2`, `colStart = 0`, `colEnd = 3`.<br>
We go right until we hit 4, then we move `rowStart` down by 1, so `rowStart = 1`.<br>
That means the first row is done exploring.<br>

There's gotcha here, we need to check if `rowStart <= rowEnd` and `colStart <= colEnd` before each direction change.<br>
Imagine we're at 5. At this point, `rowStart = 1`, `rowEnd = 1`, `colStart = 1`, `colEnd = 2`.<br>
After we hit 7, we're done exploring this row, so we move `rowStart` down by 1, so `rowStart = 2`.<br>
Now, `rowStart` is greater than `rowEnd`, that means we can't explore column anymore.<br>
That means we can't go left anymore.<br>
Same logic applies to `colStart` and `colEnd`.<br>
Therfore, after moving `rowStart` or `colEnd`, we need to check if `rowStart <= rowEnd` and `colStart <= colEnd` again before we change direction.<br>
In short, after going right and down, we need to check if we can go left and up.<br>

### Complexity Analysis
#### Time Complexity O(n)
- We iterate through the matrix once, so the time complexity is O(n).

#### Space Complexity O(1)
- We don't use any extra space, so the space complexity is O(1).