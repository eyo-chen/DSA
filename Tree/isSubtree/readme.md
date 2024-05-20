# Problem Explanation

The idea to solve this problem is easy.<br>
We traverse each node, and check if the current node is the same as the subtree we are looking for.<br>

For example, if we have the following tree:<br>
```
    3
   / \
  4   5
 / \
1   2
```

And we are looking for the subtree:<br>
```
  4
 / \
1   2
```

We first look at node 3, and ask Is the current node the same as the subtree we are looking for?<br>
If not, we go to the left and right child of node 3, and ask the same question.<br>
If yes, we return True.<br>

# Complexity Analysis
## Time Complexity O(N)
- Where N is the number of nodes in the tree.<br>

## Space Complexity O(H)
- Where H is the height of the tree.<br>