# Problem Explanation

Given the root of a binary tree, we need to find the maximum width of the tree, where the width of a level is the number of nodes (including null nodes) between the leftmost and rightmost non-null nodes, as if the tree were complete. The answer must fit within a 32-bit signed integer.

We can't just assume the last level has the maximum width. For example, consider this tree:
```
       1
      / \
     3   2
    / \   \
   5   3   9
  / 
 6           
```

In this case, the maximum width might not be at the deepest level, so we need to compute the width for each level and find the maximum.

## Core Idea

To solve this, we treat each level as an array of nodes, including null nodes, and calculate the width as the difference between the largest and smallest indices of non-null nodes plus one. We use **Breadth-First Search (BFS)** to traverse the tree level by level, assigning indices to nodes as if they were in a complete binary tree.

### How to Calculate the Width for Each Level?

We view each level as an array. For example, for the tree above:
- First level: `[1]`
- Second level: `[3, 2]`
- Third level: `[5, x, x, 9]`
- Fourth level: `[6]`

To find the width, we only need the smallest and largest indices of non-null nodes in each level’s array. We don’t need to build the actual array—just track these indices.

### How to Calculate Indices for Each Node?

We assign indices based on a complete binary tree:
- Root node starts at index `0`.
- Left child: `2 * parent_index`.
- Right child: `2 * parent_index + 1`.

For example, in the second level, node `3` (left child of `1`) has index `2 * 0 = 0`, and node `2` (right child of `1`) has index `2 * 0 + 1 = 1`.

### Step-by-Step Breakdown

1. Initialize a queue with the root node and index `0`.
2. For each level:
   - Record the minimum and maximum indices of non-null nodes.
   - Calculate the width as `max_index - min_index + 1`.
   - Enqueue children with indices `2 * parent_index` (left) and `2 * parent_index + 1` (right).
   - Normalize indices by subtracting the first node’s index in the next level to avoid overflow, as indices can grow exponentially (`2^depth`).
3. Update the global maximum width if the current level’s width is larger.
4. Continue until the queue is empty.
5. Return the maximum width.

The key to preventing overflow is normalizing indices by subtracting the first node’s index in the next level, keeping numbers manageable.

## Example Walkthrough

**Input**: `[1,3,2,5,null,null,9,6,null,7]`

**Tree Structure**:

```
       1
      / \
     3   2
    /     \
   5       9
  /         \
 6           7
```

**Step-by-Step**:

1. **Level 1**:
   - Queue: `[{node: 1, index: 0}]`.
   - Min index = `0`, max index = `0`.
   - Width = `0 - 0 + 1 = 1`.
   - Enqueue: `3` (index `2 * 0 = 0`), `2` (index `2 * 0 + 1 = 1`).
   - Queue: `[{node: 3, index: 0}, {node: 2, index: 1}]`.
   - Max width = `1`.
   - Normalize: Subtract `0` (first index in next level). Queue: `[{node: 3, index: 0}, {node: 2, index: 1}]`.

2. **Level 2**:
   - Queue: `[{node: 3, index: 0}, {node: 2, index: 1}]`.
   - Min index = `0`, max index = `1`.
   - Width = `1 - 0 + 1 = 2`.
   - Enqueue: `5` (index `2 * 0 = 0`), `9` (index `2 * 1 + 1 = 3`).
   - Queue: `[{node: 5, index: 0}, {node: 9, index: 3}]`.
   - Max width = `max(1, 2) = 2`.
   - Normalize: Subtract `0`. Queue: `[{node: 5, index: 0}, {node: 9, index: 3}]`.

3. **Level 3**:
   - Queue: `[{node: 5, index: 0}, {node: 9, index: 3}]`.
   - Min index = `0`, max index = `3`.
   - Width = `3 - 0 + 1 = 4`.
   - Enqueue: `6` (index `2 * 0 = 0`), `7` (index `2 * 3 + 1 = 7`).
   - Queue: `[{node: 6, index: 0}, {node: 7, index: 7}]`.
   - Max width = `max(2, 4) = 4`.
   - Normalize: Subtract `0`. Queue: `[{node: 6, index: 0}, {node: 7, index: 7}]`.

4. **Level 4**:
   - Queue: `[{node: 6, index: 0}, {node: 7, index: 7}]`.
   - Min index = `0`, max index = `7`.
   - Width = `7 - 0 + 1 = 8`.
   - Enqueue: (none, all children are null).
   - Queue: `[]`.
   - Max width = `max(4, 8) = 8`.

**Output**: `8` (maximum width at level 4).

### Why Width is 8 at Level 4?

At level 4, nodes `6` (index `0`) and `7` (index `7`) are the leftmost and rightmost non-null nodes. In a complete binary tree at level 4, indices range from `0` to `7`. The width includes null nodes between them: `[6, x, x, x, x, x, x, 7]`, so `7 - 0 + 1 = 8`.

## Notes

- **Normalization**: Subtracting the first node’s index per level prevents overflow, crucial for deep trees where indices grow as `2^depth`.
- **Edge Cases**:
  - Single node: Width = `1`.
  - Skewed tree: Width may be `1` per level.
  - Deep tree: Normalization ensures indices stay within bounds.
- **Index Tracking**: Using `2 * parent_index` and `2 * parent_index + 1` mimics a complete binary tree, correctly accounting for null nodes.

## Complexity Analysis

- **Time Complexity**: O(n), where `n` is the number of nodes, as we visit each node once during BFS.
- **Space Complexity**: O(w), where `w` is the maximum width of the tree, reflecting the maximum queue size at any level.