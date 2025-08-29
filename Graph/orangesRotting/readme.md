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

# One Small Gotcha
I implemented a solution on 2025/08/23 that was very similar to the correct solution. <br>
However, I didn't include the condition `freshCount > 0` in the outer for-loop, which resulted in an incorrect approach. <br>
Why does this happen? Let's use an example to find out.<br>

**Initial Grid:**
```
[2, 1, 1]
[1, 1, 0]
[0, 1, 1]
```

**Setup:**
- `queue = [[0,0]]` (position of the initial rotten orange)
- `freshCount = 6` (counting all the `1`s)

## First Solution (INCORRECT) - Trace:

### Minute 0 → 1:
- **Before:** `queue = [[0,0]]`, `freshCount = 6`
- Process rotten orange at `[0,0]`
- Rot adjacent fresh oranges: `[0,1]` and `[1,0]`
- **After processing:** `queue = [[0,1], [1,0]]`, `freshCount = 4`
- **Grid becomes:**
```
[2, 2, 1]
[2, 1, 0]
[0, 1, 1]
```
- `minute++` → `minute = 1`

### Minute 1 → 2:
- **Before:** `queue = [[0,1], [1,0]]`, `freshCount = 4`
- Process `[0,1]`: rots `[0,2]`
- Process `[1,0]`: rots `[1,1]`
- **After processing:** `queue = [[0,2], [1,1]]`, `freshCount = 2`
- **Grid becomes:**
```
[2, 2, 2]
[2, 2, 0]
[0, 1, 1]
```
- `minute++` → `minute = 2`

### Minute 2 → 3:
- **Before:** `queue = [[0,2], [1,1]]`, `freshCount = 2`
- Process `[0,2]`: no adjacent fresh oranges
- Process `[1,1]`: rots `[2,1]`
- **After processing:** `queue = [[2,1]]`, `freshCount = 1`
- **Grid becomes:**
```
[2, 2, 2]
[2, 2, 0]
[0, 2, 1]
```
- `minute++` → `minute = 3`

### Minute 3 → 4:
- **Before:** `queue = [[2,1]]`, `freshCount = 1`
- Process `[2,1]`: rots `[2,2]`
- **After processing:** `queue = [[2,2]]`, `freshCount = 0`
- **Grid becomes:**
```
[2, 2, 2]
[2, 2, 0]
[0, 2, 2]
```
- `minute++` → `minute = 4`

### Minute 4 → 5:
- **Before:** `queue = [[2,2]]`, `freshCount = 0` ⚠️
- **The problem:** Even though `freshCount = 0` (no fresh oranges left), we still enter the loop because `len(queue) > 0`
- Process `[2,2]`: no adjacent fresh oranges to rot
- **After processing:** `queue = []`, `freshCount = 0`
- `minute++` → `minute = 5` ❌

**First solution returns: 5** (WRONG!)

## Second Solution (CORRECT) - Key Difference:

The trace is identical until Minute 4, but then:

### Minute 4:
- **Before:** `queue = [[2,2]]`, `freshCount = 0`
- **Loop condition check:** `len(queue) > 0 && freshCount > 0`
- This is `true && false = false`
- **We don't enter the loop!**
- `minute` stays at 4

**Second solution returns: 4** (CORRECT!)
