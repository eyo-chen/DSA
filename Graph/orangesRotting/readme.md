# Problem Explanation

When we saw ***matrix*** and ***minimum*** in the problem description, we should think of ***BFS***.<br>

The idea should be clear once we know that we can use BFS to solve this problem.<br>
We just need to put all the rotten oranges into the queue, and then we can do BFS to explore layer by layer to find the minimum time.<br>

However, there is one edge case we should consider<br>
See the following example:<br>
```
2  1  1  0 
1  1  0  0
0  0  0  1
```
In this case, we know that we start at (0, 0) to do the BFS<br>
However, the fresh orange at (3, 3) will never be rotten because it is isolated from the rotten oranges.<br>
How can we handle this case?<br>
We can simply count the number of fresh oranges before we start the BFS.<br>
If the number of fresh oranges is greater than 0 after the BFS, we can return -1, which means there are some fresh oranges that will never be rotten.<br>

Summarize how to solve this problem:<br>
1. Loop through the matrix
  - Count the number of fresh oranges
  - Put the rotten oranges into the queue
2. Do BFS
  - For each rotten orange, we can explore the four directions
  - If we find a fresh orange, we can make it rotten and put it into the queue
  - Decrease the number of fresh oranges by 1
  - Increase the time(minute) by 1


Let's walk through the example:<br>
```
2  1  1
1  1  0
0  1  1
```
After loop through the matrix, we have:<br>
- fresh = 6
- queue = [(0, 0)]

First Iteration:<br>
- (0, 0) -> (1, 0), (0, 1)
- queue = [(1, 0), (0, 1)]
- fresh = 4
- minute = 1
```
2  2  1
2  1  0
0  1  1
```

Second Iteration:<br>
- (1, 0) -> (2, 0), (1, 1)
- (0, 1) -> x
- queue = [(2, 0), (1, 1)]
- fresh = 2
- minute = 2
```
2  2  2
2  2  0
0  1  1
```

Third Iteration:<br>
- (2, 0) -> x
- (1, 1) -> (2, 1)
- queue = [(2, 1)]
- fresh = 1
- minute = 3
```
2  2  2
2  2  0
0  2  1
```

Fourth Iteration:<br>
- (2, 1) -> (2, 2)
- queue = [(2, 2)]
- fresh = 0
- minute = 4
```
2  2  2
2  2  0
0  2  2
```

After the BFS, we have:<br>
- fresh = 0
- minute = 4

# Complexity Analysis
## Time Complexity: O(n * m)
- n is the number of rows
- m is the number of columns
- At least, we need to loop through the matrix once to count the number of fresh oranges

## Space Complexity: O(n * m)
- n is the number of rows
- m is the number of columns
- The space complexity is the space used by the queue
- The worst case is that all the oranges are rotten, so the space complexity is O(n * m)


