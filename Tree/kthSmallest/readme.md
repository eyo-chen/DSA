# Explanation

The core idea is actually quite straightforward. Since we're dealing with a binary search tree, the nodes are naturally sorted when viewed from left to right.

```
          3
      1      4
         2
```

If we imagine a vertical line sweeping from left to right across the tree, we can see that the values are encountered in sorted order: 1, 2, 3, 4.

Understanding this property gives us two approaches:

## Approach 1: Complete In-Order Traversal
Simply perform an in-order traversal to collect all node values in an array, then return the element at index `k-1`.

## Approach 2: Early Termination with Counter
We still use in-order traversal, but maintain a counter `k` throughout the process. 

Recall that in-order traversal follows this pattern:
- Traverse left subtree
- Process current node
- Traverse right subtree

In the first approach, "process current node" means storing the value in an array. In this optimized approach, we decrement our counter `k` instead.

When `k` reaches 0, we know we've found the kth smallest element and can return immediately without visiting the remaining nodes.