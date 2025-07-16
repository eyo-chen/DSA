# Smallest Subtree with Deepest Nodes

## Problem Statement

Given the root of a binary tree, we need to find the smallest subtree that contains all the deepest nodes in the tree. The depth of each node is defined as the shortest distance to the root.

**Key Definitions:**
- **Deepest nodes**: Nodes that have the largest depth possible among any node in the entire tree
- **Subtree**: A tree consisting of a node plus the set of all its descendants
- **Smallest subtree**: Among all subtrees that contain all deepest nodes, the one with the fewest nodes

**Examples:**

**Example 1:** Tree `[3,5,1,6,2,0,8,null,null,7,4]` → Output: `[2,7,4]`
```
        3 (depth 0)
       / \
      5   1 (depth 1)
     /|   |\
    6 2   0 8 (depth 2)
     /|
    7 4 (depth 3) ← deepest nodes
```
Node 2's subtree contains both deepest nodes 7 and 4.

**Example 2:** Tree `[1]` → Output: `[1]`
```
1 (depth 0) ← only node, hence deepest
```

**Example 3:** Tree `[0,1,3,null,2]` → Output: `[2]`
```
    0 (depth 0)
   / \
  1   3 (depth 1)
   \
    2 (depth 2) ← deepest node
```

## Thought Process

### Initial Observations

The problem can be broken down into two key insights:

1. **Identify the target**: We need to find all nodes that are at the maximum depth in the tree
2. **Find the optimal container**: Among all subtrees that contain these deepest nodes, we want the smallest one

### Key Insight: Lowest Common Ancestor (LCA)

The crucial realization is that the "smallest subtree containing all deepest nodes" is actually the **Lowest Common Ancestor (LCA)** of all deepest nodes. Here's why:

- Any subtree that contains all deepest nodes must be rooted at some ancestor of these nodes
- The LCA is the "lowest" (closest to leaves) such ancestor
- Any ancestor of the LCA would give us a larger subtree, making it suboptimal

### From LCA to Algorithm Design

This insight transforms our problem into: "Find the LCA of all deepest nodes." But we can solve this more elegantly by processing the tree bottom-up and making decisions at each node based on where the deepest nodes are located.

## Solution Approach

### High-Level Strategy

We use a **depth-first search (DFS)** with a **bottom-up approach**:

1. For each node, determine the maximum depth in its subtree
2. Based on the depths of left and right subtrees, decide where the deepest nodes are located
3. Return the appropriate result based on the distribution of deepest nodes

### What Information Do We Track?

For each node, our recursive function returns:
- **Depth**: The maximum depth of any node in this subtree
- **Node**: The root of the smallest subtree containing all deepest nodes in this subtree

### Decision Logic at Each Node

At any given node, we have three possible scenarios:

1. **Deepest nodes only in left subtree**: `leftDepth > rightDepth`
   - Pass up the answer from the left subtree
   
2. **Deepest nodes only in right subtree**: `rightDepth > leftDepth`
   - Pass up the answer from the right subtree
   
3. **Deepest nodes in both subtrees**: `leftDepth == rightDepth`
   - Current node is the LCA of all deepest nodes
   - Return the current node as the answer

### Why This Logic Works

The algorithm works because of this fundamental property:

**When both left and right subtrees have the same maximum depth, it means the deepest nodes are distributed across both subtrees, making the current node their LCA.**

Conversely, if one subtree has a greater maximum depth, all the deepest nodes must be in that subtree, so we delegate the answer to that subtree.

## Detailed Solution Breakdown

### Step 1: Base Case Handling

When we encounter a `null` node:
- Return depth as -1 (indicating no nodes in this subtree)
- Return `null` as the node

### Step 2: Recursive Depth Calculation

For each node:
1. Recursively calculate results for left and right subtrees
2. Current node's depth = `max(leftDepth, rightDepth) + 1`

### Step 3: Determining the Answer Node

Based on the comparison of left and right depths:

```
if leftDepth > rightDepth:
    // All deepest nodes are in left subtree
    return {depth: currentDepth, node: leftResult.node}
else if rightDepth > leftDepth:
    // All deepest nodes are in right subtree  
    return {depth: currentDepth, node: rightResult.node}
else:
    // Deepest nodes in both subtrees, current node is LCA
    return {depth: currentDepth, node: currentNode}
```

### Step 4: Propagating Results Up

The beauty of this approach is that information flows naturally upward:
- If we find the answer at a lower level, it gets passed up through the recursive calls
- If the answer is at a higher level, the comparison logic will identify it correctly

## Why This Solution is Optimal

### Time Complexity: O(n)
- We visit each node exactly once in our DFS traversal
- At each node, we perform constant-time operations

### Space Complexity: O(h)
- Where h is the height of the tree
- This is due to the recursive call stack

### Correctness Guarantee

The algorithm guarantees correctness because:
1. **Bottom-up processing**: We have complete information about subtrees before making decisions
2. **Proper depth tracking**: We accurately identify where the deepest nodes are located
3. **LCA identification**: The equal-depth condition correctly identifies the LCA

## Edge Cases Handled

The solution elegantly handles all edge cases:

1. **Single node tree**: The node itself is both deepest and the answer
2. **All deepest nodes under same parent**: That parent becomes the answer
3. **Deepest nodes scattered across tree**: The algorithm finds their LCA correctly
4. **Unbalanced trees**: The depth comparison logic works regardless of tree structure

## Algorithm Intuition

Think of the algorithm as asking at each node: "Where are the deepest nodes in my subtree?"

- If they're all in my left subtree, pass up the left answer
- If they're all in my right subtree, pass up the right answer  
- If they're in both subtrees, then I'm the meeting point (LCA)

This intuitive approach leads to an elegant and efficient solution that processes the tree in a single pass while maintaining all necessary information to make the correct decision at each level.