# Problem Explanation
Consider the following tree:
```
      4
    /   \
   2     7
  / \   / \
 1   3 6   9
```
after inverting the tree, it will look like:
```
      4
    /   \
   7     2
  / \   / \
 9   6 3   1
```

How can we invert a binary tree?<br>
The core idea is to swap the left and right child of each node ***FROM THE BOTTOM TO THE TOP***.<br>
***FROM THE BOTTOM TO THE TOP*** is very important to note, because if we swap the left and right child of each node from the top to the bottom, we will lose the original left and right child of the node.<br>

Use above tree as example,<br>
We first swap the leaf nodes (1 <-> 3, 6 <-> 9)<br>
```
      4
    /   \
   2     7
  / \   / \
 3   1 9   6
```

Then we swap the parent nodes subtree (2 <-> 7)<br>
```
      4
    /   \
   7     2
  / \   / \
 9   6 3   1
```
Done!

# Complexity Analysis
## Time Complexity O(N)
- Where N is the number of nodes in the tree.<br>
- We visit each node once.<br>

## Space Complexity O(H)
- Where H is the height of the tree.<br>