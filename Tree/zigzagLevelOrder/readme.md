# Solution Explanation

## Solution 1: Use Queue and Stack
This solution uses a **queue** for standard BFS traversal and an auxiliary **stack** to control the order of child node insertion. The key insight is that by alternating the order in which we add left and right children to the queue, we naturally achieve the zigzag pattern in the next level.

- Queue stores the "correct" order of nodes to be processed
  - Note that queue is very much like typical BFS, it stores the "correct" order of nodes to be processed
  - We just pop nodes from the queue and add their values to the current level array
- Stack temporarily holds nodes to reverse the order of child insertion based on the current level (even or odd)
  - The reason we need a stack is because it's LIFO (Last In First Out) behavior allows us to reverse the order of child insertion
  - Look at the below example
    - When we are the second level [20,9]
    - When we treat it as queue, we pop 20 first, then 9. And add them to current level array in that order
    - However, for the next level, we actually want to pop 9 first, and process its children first (5,3), then pop 20 and process its children (15,7)
    - This is exactly what stack does for us
    - We basically stores the same orders of nodes as processed from queue, but when we pop from stack, we get the reverse order

**Core Idea:**
- Use queue for level-by-level processing
- Use stack as temporary storage to reverse the order of child insertion
- For even levels (0, 2, 4...): add right child first, then left child
- For odd levels (1, 3, 5...): add left child first, then right child
- The values are collected in natural queue order, but the zigzag effect comes from the reversed child insertion

```
Tree structure:
       3
     /   \
    9     20
   / \   /  \
  5   3 15   7
```

**Initial State:**
- `queue = [3]`, `stack = []`, `level = 0`, `ans = []`

**Level 0 (even):**
1. Process queue: `node = 3`, `curLevel = [3]`
2. Add to stack: `stack = [3]`
3. Pop from stack and add children (right first for even level):
   - Add 20, then 9: `queue = [20, 9]`
4. Result: `ans = [[3]]`, `level = 1`

**Level 1 (odd):**
1. Process queue: `node = 20`, `node = 9`, `curLevel = [20, 9]`
2. Add to stack: `stack = [20, 9]`
3. Pop from stack (9 first, then 20) and add children (left first for odd level):
   - From 9: Add 5, then 3: `queue = [5, 3]`
   - From 20: Add 15, then 7: `queue = [5, 3, 15, 7]`
4. Result: `ans = [[3], [20, 9]]`, `level = 2`

**Level 2 (even):**
1. Process queue: `curLevel = [5, 3, 15, 7]`
2. No more children to process
3. Result: `ans = [[3], [20, 9], [5, 3, 15, 7]]`

### Complexity Analysis
#### Time Complexity O(n) 
- Each node is visited exactly once

#### Space Complexity O(w) 
- Where w is the maximum width of the tree
  - Queue can hold at most one level of nodes
  - Stack temporarily holds the same nodes as processed from queue
  - In the worst case (complete binary tree), w = n/2



## Solution 2: Position-Based Approach
This solution uses a single **queue** for BFS but calculates the **position** where each node's value should be placed in the current level array. The zigzag effect is achieved by alternating the position calculation formula.

**Core Idea:**
- Use standard BFS with queue
- Pre-allocate level array with known size
- Calculate position based on direction flag:
  - Right-to-left: `position = i` (normal insertion)
  - Left-to-right: `position = size - i - 1` (reverse insertion)
- Always add children in the same order (right first, then left)

Let's see the calculation of positions in detail:<br>
- Second Level (right to left):
  - For size = 2, queue = [20, 9], indices = [0, 1]
  - i = 0 → position = 0 -> curLevel[0] = 20
  - i = 1 → position = 1 -> curLevel[1] = 9
- Third Level (left to right):
  - For size = 4, queue = [7, 15, 3, 5], indices = [0, 1, 2, 3]
  - i = 0 → position = 4 - 0 - 1 = 3 -> curLevel[3] = 7
  - i = 1 → position = 4 - 1 - 1 = 2 -> curLevel[2] = 15
  - i = 2 → position = 4 - 2 - 1 = 1 -> curLevel[1] = 3
  - i = 3 → position = 4 - 3 - 1 = 0 -> curLevel[0] = 5


```
Tree structure:
       3
     /   \
    9     20
   / \   /  \
  5   3 15   7
```

**Initial State:**
- `queue = [3]`, `rightToLeft = false`, `ans = []`

**Level 0:**
- `size = 1`, `curLevel = [0]` (pre-allocated)
- Process node 3: `position = 1 - 0 - 1 = 0`
- `curLevel[0] = 3` → `curLevel = [3]`
- Add children: `queue = [20, 9]`
- Result: `ans = [[3]]`, `rightToLeft = true`

**Level 1:**
- `size = 2`, `curLevel = [0, 0]` (pre-allocated)
- Process node 20: `position = 0` → `curLevel[0] = 20`
- Process node 9: `position = 1` → `curLevel[1] = 9`
- Add children: `queue = [15, 7, 5, 3]`
- Result: `ans = [[3], [20, 9]]`, `rightToLeft = false`

**Level 2:**
- `size = 4`, `curLevel = [0, 0, 0, 0]` (pre-allocated)
- Process node 15: `position = 4 - 0 - 1 = 3` → `curLevel[3] = 15`
- Process node 7: `position = 4 - 1 - 1 = 2` → `curLevel[2] = 7`
- Process node 5: `position = 4 - 2 - 1 = 1` → `curLevel[1] = 5`
- Process node 3: `position = 4 - 3 - 1 = 0` → `curLevel[0] = 3`
- Result: `ans = [[3], [20, 9], [3, 5, 7, 15]]`

### Complexity Analysis
#### Time Complexity O(n)
- Each node is visited exactly once

#### Space Complexity O(w)
- Where w is the maximum width of the tree
- Queue holds at most one level of nodes
- Level arrays are allocated with exact size needed
- More space-efficient than Solution 1 (no auxiliary stack needed)