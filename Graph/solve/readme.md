# Problem Explanation

We know we need to model this problem as a graph problem. However, how to really solve this problem after modeling it as a graph problem. Let's find out.<br>

The problem is to find the point where it's not connected with the border. At first, we might think of using DFS or BFS to traverse the graph. However, there's a problem need to be clarified:<br>
*** The whole matrix can be composed of multiple graphs ***<br>

Now, the next step is to think how to efficiently traverse all the graphs in the matrix. We can have two options:<br>
1. Traverse the graph inwards to outwards
2. Traverse the graph outwards to inwards

```
      1     2    3    4
  1  ['O', 'X', 'X', 'X']
  2  ['X', 'O', 'O', 'X']
  3  ['X', 'X', 'X', 'X']
  4  ['X', 'O', 'O', 'X']
  5  ['X', 'X', 'O', 'O']
```
In first option, we might start from (2,2), (2,3), (4,2) or (4,3)
In second option, we only start from (1,1), (5,3) and (5,4)

The second option is more efficient than the first option because
1. If we traverse from inwards, we might flip the wrong point to 'X' which is connected with the border
  - For example, suppose we start from (4,2), and we traverse all the adjacent points where it's 'O', and we flip it to 'X'. However, the point (4,2) is actually connected with the border, so we did the wrong operation.
2. If we traverse from outwards, we can guarantee that all the points we traverse are connected with the border
  - In other words, these points don't need to be flipped to 'X'
  - We also can say these points are the *safe* points
  - After we know what the safe points are, we can just simply flip the points where it's 'O' but not in the safe points to 'X'


Now, we can summarize the steps to solve this problem:
1. Traverse the graph outwards to inwards
   - Traverse the first row and the last row
     - If it's 'O', we need to mark all the connected points where it's 'O' to safe points
   - Traverse the first column and the last column
     - If it's 'O', we need to mark all the connected points where it's 'O' to safe points
2. Traverse the whole matrix
   - If it's 'O' and not in the safe points, we need to flip it to 'X'


Note that we have two ways to mark the safe points:
1. Using Hash Set
  - use the location as the key
2. Mutating the matrix
  - we can mark the safe points as 'S' to point out that it's safe


# Complexity Analysis
## Time Complexity O(m*n)
- where m is the number of rows and n is the number of columns in the matrix

## Space Complexity O(m*n)
- where m is the number of rows and n is the number of columns in the matrix
