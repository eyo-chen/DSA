# Solution Explanation

This problem asks us to reverse node values at odd levels in a perfect binary tree. Two approaches are provided:

**Approach 1: Recursive DFS (Two-Pointer)**
The key insight is that in a perfect binary tree, we can use a recursive approach that simultaneously traverses from both the left and right sides of each level. At odd levels, we swap values between corresponding nodes from opposite ends. This is exactly what the diagram illustrates - **swapping values from the very left and right** of each odd level.

**Approach 2: Iterative BFS (Level-Order)**
Uses a queue to perform level-order traversal. At each odd level, we collect all nodes in an array and swap values from both ends toward the center.

## Walkthrough with Example

Let's trace through the example shown in the diagram with root = [2,3,5,8,13,21,34]:

### Initial Tree Structure:
```
Level 0 (even):            2
                        /     \
Level 1 (odd):         3       5
                    /   \    /   \
Level 2 (even):    8    13  21   34
                  /|    |\  /|    |\
Level 3 (odd):  30 31 32 33 34 35 36 37
```

### Approach 1: Recursive DFS

Starting from the root, we call `helper(root.Left, root.Right, 1)`:

**Level 1 (odd):**
- `helper(node3, node5, 1)` is called
- Since level 1 is odd, **swap values**: node3.Val ↔ node5.Val
- Tree becomes: [2, **5**, **3**, 8,13,21,34,...]
- Then recursively process:
  - `helper(node8, node34, 2)` - level 2 (even), no swap
  - `helper(node13, node21, 2)` - level 2 (even), no swap

**Level 3 (odd):**
- Multiple recursive calls swap leaf nodes:
  - **30 ↔ 37** (leftmost with rightmost)
  - **31 ↔ 36** (second-left with second-right)
  - **32 ↔ 35** (third-left with third-right)
  - **33 ↔ 34** (center pairs)

The critical pattern here is that the recursive calls `helper(left.Left, right.Right, level+1)` and `helper(left.Right, right.Left, level+1)` ensure we're always **pairing nodes from opposite ends** of each level, which is exactly what the diagram shows.

### Approach 2: Iterative BFS

**Level 0:** Process root (2), no swap (even level)

**Level 1:** 
- Queue contains [3, 5]
- Odd level, so collect nodes: [3, 5]
- Swap: 3 ↔ 5
- Result: [2, **5**, **3**, ...]

**Level 2:**
- Queue contains [8, 13, 21, 34]
- Even level, no swap

**Level 3:**
- Queue contains [30, 31, 32, 33, 34, 35, 36, 37]
- Odd level, swap from both ends:
  - 30 ↔ 37
  - 31 ↔ 36
  - 32 ↔ 35
  - 33 ↔ 34

### Final Result:
```
         2
        / \
       5   3
      / \ / \
     8 13 21 32
    /| |\ /| |\
   37 36 35 34 33 32 31 30
```

## Time and Space Complexity

### Approach 1: Recursive DFS
- **Time Complexity:** O(n)
  - We visit each node exactly once
  - Each swap operation is O(1)
  - n = total number of nodes

- **Space Complexity:** O(log n) = O(h)
  - Recursion stack depth equals the height of the tree
  - Since it's a perfect binary tree: h = log₂(n)
  - No additional data structures needed

### Approach 2: Iterative BFS
- **Time Complexity:** O(n)
  - We visit each node exactly once in the BFS traversal
  - For each odd level with k nodes, swapping takes O(k)
  - Total: O(n) across all levels

- **Space Complexity:** O(n/2) = O(n)
  - Queue can hold up to all nodes at the last level
  - In a perfect binary tree, the last level has n/2 nodes
  - Additional array to store current level nodes: O(width of level)
  - Worst case: O(n)

**Key Takeaway:** The recursive approach leverages the perfect binary tree structure to swap corresponding nodes from opposite ends without explicitly storing level information, making it the more elegant solution.