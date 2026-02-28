# Problem Explanation

## Create Array To Store Zero Position
The idea is to use an array to store the position of zero<br>
Then, we can loop through the array to set the row and col to zero<br>

For example, input matrix
```
[1, 2, 3, 4]
[5, 0, 7, 8]
[0, 10, 11, 12]
[13, 14, 15, 0]
```
We can simply put all the zero position into an array
```
[(1,1), (2,0), (3,3)]
```
Then, we can loop through the array to set the row and col to zero

There are two drawbacks of this solution<br>
1. It uses extra space to store the zero position<br>
2. It set one position to zero multiple times<br>
   - Look at the above example, `matrix[2][1](10)` will be set to zero twice because `(1,1)` and `(2,0)` both are in the array

### Complexity Analysis
#### Time Complexity O(m*n(m + n))
- where m is the number of row, n is the number of col
- First loop to find zero positions: O(m*n)
- Second loop, for each zero position found, we traverse:
  - Up and down the column: O(m)
  - Left and right across the row: O(n)
  - If we have k zeros in the matrix (worst case k = m*n), then setting zeros takes O(k (m + n))
- Therefore, total time complexity is O(m*n) + O(m*n (m + n)) = O(m*n (m + n))

#### Space Complexity O(m*n)
- where m is the number of row, n is the number of col
- In the worst case, where all elements are 0, the zeroPosition slice will store coordinates for all elements
- Therefore, space complexity is O(m*n)



## Create Two Arrays(First Row and First Col) To Store Zero Position
In order to avoid mutating one position multiple times, we can use the first row and first col to store the zero position<br>
We create two arrays, `firstRow` and `firstCol` to store the zero position<br>
`firstRow` represents whether the r-th column need to be set to zero<br>
`firstCol` represents whether the c-th row need to be set to zero<br>

For example, input matrix
```
       true   true   false   false <- firstRow
false    1      2      3      4
false    5      0      7      8
true     0      10     11     12
true     13     14     15     0

^
|
|-> firstCol
```
Once we have the `firstRow` and `firstCol`, we can traverse the matrix again to set the zero
- If `firstRow[r]` is true, set all element in row r to zero
- If `firstCol[c]` is true, set all element in col c to zero

### Complexity Analysis
#### Time Complexity O(m*n)
- where m is the number of row, n is the number of col
- First loop to find zeros: O(m*n)
- Second loop for rows: O(m*n) in worst case
- Third loop for columns: O(m*n) in worst case
- Therefore, total time complexity is O(m*n) + O(m*n) + O(m*n) = O(m*n)

#### Space Complexity O(m + n)
- where m is the number of row, n is the number of col
- firstRow array: O(m)
- firstCol array: O(n)
- Therefore, space complexity is O(m + n)

## Using First Row and First Col To Store Zero Position
This idea is a little bit tricky<br>
The overall idea is to NOT create any extra space to store the zero position<br>
Instead, we use the first row and first col to store the zero position<br>

For example, input matrix
```
1  2  3  4
5  0  7  8
0  10 11 12
13 14 15 0
```
We convert the matrix to
```
0  0  3  0
0  0  7  8
0  10 11 12
0  14 15 0
```
As we can see, the first row and first col are mutated to 0 to indicate whether the specific row or col need to be set to zero<br>

At first, we might summarize the solution as follows<br>
1. Traverse the matrix, if matrix[r][c] is 0, set matrix[r][0] and matrix[0][c] to 0
2. Loop through the first row, if matrix[0][c] is 0, set all element in col c to zero
3. Loop through the first col, if matrix[r][0] is 0, set all element in row r to zero

However, this solution is incorrect because it will set the first row and first col to zero first, which will affect our later logic<br>
For example, consider if the input matrix is
```
0  1  2  0
3  4  5  2
1  3  1  5
```
After the first traversal, the matrix will be converted to
```
0  1  2  0
3  4  5  2
1  3  1  5
```
Then, loop through the first row, if matrix[0][c] is 0, set all element in col c to zero
```
0  1  2  0
0  4  5  0
0  3  1  0
```
Then, loop through the first col, if matrix[r][0] is 0, set all element in row r to zero
```
0  0  0  0
0  0  0  0
0  0  0  0
```
As we can see, the result is incorrect<br>
The reason is that when looping through the first row, we are also changing the value of the first column, which will affect our later logic<br>

In order to fix this, the core idea is that we DO NOT mutate the value of the first row and first column when we are traversing the matrix<br>
For example, input matrix is
```
0  1  2  0
3  4  5  2
4  0  1  5
```

After the first traversal, we mutating the first row and first col to 0 when we encounter a zero, so the matrix is
```
0  0  2  0
3  4  5  2
0  4  1  5
```
Then, we traverse the matrix again<br>
BUT, now we only traverse from matrix[1][1] to matrix[m-1][n-1], so we will not be affected by the mutation of the first row and first col<br>
After the second traversal, the matrix will be converted to
```
0  0  2  0
3  0  5  0
0  0  0  0
```

After mutating the innner matrix(excluding the first row and first col), we can mutate the first row and first col<br>
But how do we know whether to set the first row and first col to zero?<br>
We can use two boolean variable to store if the first row and first col need to be set to zero<br>
- `isZeroAtFirstRow` to store whether the first row need to be set to zero
- `isZeroAtFirstCol` to store whether the first col need to be set to zero

In order to determine if the first row and first col need to be set to zero, we can traverse the first row and first col at the beginning<br>
- Traverse the first row, if matrix[0][c] is 0, set `isZeroAtFirstRow` to true. It indicates that we should set the first row to zero.
- Traverse the first col, if matrix[r][0] is 0, set `isZeroAtFirstCol` to true. It indicates that we should set the first col to zero.

Let's summarize the whole process
1. Traverse the first row, if matrix[0][c] is 0, set `isZeroAtFirstRow` to true. It indicates that we should set the first row to zero later.
2. Traverse the first col, if matrix[r][0] is 0, set `isZeroAtFirstCol` to true. It indicates that we should set the first col to zero later.
3. Traverse whole matrix, if matrix[r][c] is 0, set matrix[r][0] and matrix[0][c] to 0. This is to indicate that the specific row and col need to be set to zero.
4. Traverse the innner matrix(excluding the first row and first col), if either it's corresponding first row(matrix[0][c]) or first col(matrix[r][0]) is 0, set matrix[r][c] to 0.
5. After traversing the innner matrix, if `isZeroAtFirstRow` is true, set all element in first row to zero.
6. If `isZeroAtFirstCol` is true, set all element in first col to zero.

### Complexity Analysis
#### Time Complexity O(m*n)
- where m is the number of row, n is the number of col
- First loop to traverse the first row: O(n)
- Second loop to traverse the first col: O(m)
- Third loop to traverse the matrix: O(m*n)
- Fourth loop to traverse the innner matrix: O(m*n)
- Fifth loop to set the first row to zero: O(n)
- Sixth loop to set the first col to zero: O(m)
- Therefore, total time complexity is O(n) + O(m) + O(m*n) + O(m*n) + O(n) + O(m) = O(m*n)

#### Space Complexity O(1)
- We do not use any extra space, only use two boolean variable to store if the first row and first col need to be set to zero

## Mutate The Matrix To Special Value
The idea is to loop through the whole matrix, if matrix[r][c] is 0, set all element in row r and col c to a special value<br>
After that, we can loop through the matrix again to set the special value to zero<br>

Number things to note
- We can't set the special value to 0 because it will affect our later logic(just like the first row and first col in the previous solution)
- In dynamic programming, we can set the special value to '.'. In static programming, we can set the special value to a value that will not be in the matrix, for example, math.MaxInt
  - The idea is that the special value should be a value that will not be in the matrix, so that we can identify if the value is the special value or not
- We can only mutate the value to special value when the value is NOT 0
```
1  2  3  4
5  0  7  8
0  10 11 12
13 14 0  0
```
For example, if we DO set the special value on matrix[3][3] when encountering matrix[2][3]<br>
Then, the matrix will be converted to
```
1  2  3  4
5  0  7  8
0  10 11 12
13 14 0  .
```
Later, when we encounter matrix[3][3], we won't set it's row and col to special value because matrix[3][3] is '.'<br>
However, this is WRONG!<br>
In the correct scenario, the row and column of matrix[3][3] should be set to special value, so that we can identify that the value in row 3 and col 3 need to be set to zero<br>

### Complexity Analysis
#### Time Complexity O((m*n) * (m+n))
- where m is the number of row, n is the number of col
- For each cell in the matrix(m*n), if we find a zero:
  - We need to traverse the whole row to set the special value: O(n)
  - We need to traverse the whole col to set the special value: O(m)
- Therefore, total time complexity is O((m*n) * (m+n)) = O((m*n) * (m+n))

#### Space Complexity O(1)
- We do not use any extra space.