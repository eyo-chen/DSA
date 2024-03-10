# Problem Explanation

The problem is not that hard<br>
But has to be careful about the problem description<br>
The path has to be from the root to the leaf<br>
If targetSum is 5, and the tree looks like following
```
    1
   / \
  2   4
 / \ / \
4  5 6  7
```
It may look like the path 1->4 is a valid path<br>
But it's not what the problem is asking for<br>
We want a path from the root to the leaf<br>

The idea is very simple<br>
We just need to depth-first search the tree<br>
We decrement the targetSum by the value of the current node along with the process<br>
If we reach the leaf node and the targetSum is 0, we return True<br>

Let's walk through the example<br>
targetSum = 8<br>
```
    1
   / \
  2   4
 / \ / \
4  5 6  7
```

We start with the root node(1)<br>
targetSum = 8 - 1 = 7<br>
It's not 0<br>
So we continue to the left child(2)<br>

targetSum = 7 - 2 = 5<br>
It's not 0<br>
So we continue to the left child(4)<br>

targetSum = 5 - 4 = 1<br>
It's not 0<br>
So we continue to the left child(nullptr)<br>

It's a null node<br>
We return False<br>
Back to previous call stack<br>

Back to node(4)<br>
We've explored the left child<br>
Keep exploring the right child<br>

It's a null node<br>
We return False<br>
Back to previous call stack<br>

Back to node(2)<br>
We've explored the left child<br>
Keep exploring the right child<br>

targetSum = 5 - 5 = 0<br>
It's 0<br>
And it's also a leaf node<br>
We return True<br>

Back to the root node<br>
Now, we don't need to explore the right child<br>
Because we've already found the path<br>

## Complexity Analysis
### Time Complexity: O(n)
- At worst, we have to visit all the nodes in the tree

### Space Complexity: O(h) or O(log n)
- The space complexity is O(h) where h is the height of the tree


