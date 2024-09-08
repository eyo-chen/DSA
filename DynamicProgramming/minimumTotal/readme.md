# Problem Explanation

To solve this problem, have to first understand that we can go either `index` or `index + 1` to the nexy layer.<br>
For example, if we have a triangle like this:
```
     2
    3 4
   6 5 7
  4 1 8 3
```
When we're at 2, we can go to either `3` or `4`.<br>
When we're at 3, we can go to either `6` or `5`.<br>
When we're at 5, we can go to either `1` or `8`.<br>
So we can see that we can go to either `index` or `index + 1` to the next layer.<br>


At first, we might think that we can just find the minimum value for each layer and sum them up.<br>
But this won't work, and let's find out why.<br>
If we do that, we might get something like this:
```
      -1
     2   3
    1  -1  -3
```
We know the minimum path sum is `-1 + 3 + (-3) = -1`<br>
However, when we're at `-1`, we choose `2` instead of `3` because `2` is the minimum value for the next layer, and this won't lead us to the minimum path sum.<br>

Therefore, if we start from the top, we have to explore every possible path to find the minimum path sum.<br>

## Top-Down Approach
The idea is pretty straightforward, we start from the top of the triangle and explore every possible path to the bottom by using recursion.<br>

Base Case: <br>
When we reach the bottom of the triangle, we return the value of the current cell.<br>

Recursive Case:<br>
For each callstack, we have two choices, either go to the `index` or `index + 1` to the next layer.<br>

Along with the process, we use a map to store the result of the current cell to avoid duplicate calculation.<br>

Let's walk through an example:
```
      -1
     2   3
    1  -1  -3
```
First Callstack:
- depth = 0, index = 0
- val = -1
- go to depth = 1, index = 0

Second Callstack:
- depth = 1, index = 0
- val = 2
- go to depth = 2, index = 0

Third Callstack:
- depth = 2, index = 0
- val = 1
- go to depth = 3, index = 0
- It will hit the base case and return the value of the current cell, which is `1`.

Second Callstack:
- depth = 2, index = 0
- val = 1
- left = 1
- go to depth = 3, index = 1

Third Callstack:
- depth = 3, index = 1
- val = -1
- go to depth = 4, index = 1
- It will hit the base case and return the value of the current cell, which is `-1`.

Second Callstack:
- depth = 2, index = 0
- val = 2
- left = 1
- right = -1
- return 2 + min(1, -1) = 1

First Callstack:
- depth = 1, index = 0
- val = 2
- left = 1
- go to depth = 2, index = 1

Second Callstack:
- depth = 2, index = 1
- val = 3
- go to depth = 3, index = 1

Third Callstack:
- depth = 3, index = 1
- val = -1
- go to depth = 4, index = 1
- It will hit the base case and return the value of the current cell, which is `-1`.

Second Callstack:
- depth = 2, index = 1
- val = 3
- left = -1
- go to depth = 3, index = 2

Third Callstack:
- depth = 3, index = 2
- val = -3
- go to depth = 4, index = 2
- It will hit the base case and return the value of the current cell, which is `-3`.

Second Callstack:
- depth = 2, index = 1
- val = 3
- left = -1
- right = -3
- return 3 + min(-1, -3) = 0

First Callstack:
- depth = 1, index = 0
- val = -1
- left = 1
- right = 0
- return -1 + min(1, 0) = -1 

### Complexity Analysis
#### Time Complexity O(n^2)
- where n is the number of rows in the triangle.
- In the worst case, the function will visit each element in the triangle once.
- The triangle has n rows, and the i-th row has i elements.
- The total number of elements in the triangle is 1 + 2 + 3 + ... + n, which is n(n+1)/2.
- This sum is O(n^2).
- Each visit involves constant time operations (map lookups, simple arithmetic).
- Therefore, the overall time complexity is O(n^2).

#### Space Complexity O(n^2)
- where n is the number of rows in the triangle.a
- The space complexity is dominated by the memoization map.
- In the worst case, we store a result for each unique (depth, index) pair.
- The number of unique pairs is equal to the number of elements in the triangle, which is n(n+1)/2.
- This is O(n^2) space.
- The recursion stack will go as deep as the number of rows, which is O(n), but this is overshadowed by the memoization map.


## Bottom-Up Approach
The idea is to start from the bottom of the triangle and explore every possible path to the top by using iteration.<br>

The process is as follows:
1. Create a table to store the minimum path sum for each layer.
2. Initialize the table with the last layer of the triangle.
3. For each layer, starting from the second last layer, update the table with the minimum path sum for each cell.
4. For each cell, we find the minimum value between the `cell below it` and the `cell below it to the right` and add the value of the current cell to it.
5. The value of the first cell in the table will be the minimum path sum for the triangle after we finished iterating through the table.

Let's explain the code:
```go
	for i := l - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			v := triangle[i][j]
			m := min(table[j], table[j+1])
			table[j] = v + m
		}
	}
```
- `i` represent the current layer, `j` represent the current index in the current layer.
- `i := l - 2` means we start from the second last layer.
   - because we initialized the table with the last layer of the triangle, we don't need to consider the last layer.
- `i >= 0` means we iterate from the second last layer to the top layer.
- `j <= i`, in the second(nested) for loop, what we want to do is to iterate through all the cells in the current layer.
   - `j` starts from 0, which is easy to understand
   - because the problem is given a triangle, the number of cell in each layer is just the layer number.
   - For example, the first layer has 1 cell, the second layer has 2 cells, the third layer has 3 cells, and so on.
   - So `j <= i` means we iterate through all the cells in the current layer.
- `v := triangle[i][j]` get the value of the current cell.
- `m := min(table[j], table[j+1])` get the minimum value between the `cell below it` and the `cell below it to the right`.
- `table[j] = v + m` update the value 


Let's walk through an example:
```
      -1
     2   3
    1  -1  -3
```
After initialization:
table = [1, -1, -3]

Start from the second last layer:
i = 1
- j = 0
- curVal = triangle[1][0] = 2
- minVal = min(table[0], table[1]) = min(1, -1) = -1
- table[0] = curVal + minVal = 2 + (-1) = 1
- table = [1, -1, -3]
- j = 1
- curVal = triangle[1][1] = 3
- minVal = min(table[1], table[2]) = min(-1, -3) = -3
- table[1] = curVal + minVal = 3 + (-3) = 0
- table = [1, 0, -3]

i = 0
- j = 0
- curVal = triangle[0][0] = -1
- minVal = min(table[0], table[1]) = min(1, 0) = 0
- table[0] = curVal + minVal = -1 + 0 = -1
- table = [-1, 0, -3]

After iteration:<br>
table = [-1, 0, -3]<br>

Return table[0] = -1

### Complexity Analysis
#### Time Complexity O(n^2)
- where n is the number of rows in the triangle.
- The first loop initializes the table with the bottom row of the triangle, which takes O(n) time.
- The main nested loops iterate through the triangle from bottom to top:
  - The outer loop runs n-1 times (from l-2 to 0).
  - For each iteration i of the outer loop, the inner loop runs i+1 times.
- The total number of iterations is (n-1) + (n-2) + ... + 2 + 1, which sums up to n(n-1)/2.
- This sum is O(n^2).
- Each iteration performs constant time operations.
- Therefore, the overall time complexity is O(n^2).

#### Space Complexity O(n)
- where n is the number of rows in the triangle.
- The space complexity is determined by the `table` slice.
- The size of `table` is equal to the number of rows in the triangle, which is n.
- No other data structures grow with the input size.
- The space used is constant apart from the `table` slice.
- Therefore, the space complexity is O(n).