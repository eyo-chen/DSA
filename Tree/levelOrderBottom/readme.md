# Problem Explanation

The problem is simple, and we have two different approaches.<br>

1. Using stack to store the nodes, then pop the nodes and append the values to the result list.
2. Find the maximum depth of the tree, then append the values to the result list.

For the first approach,<br>
We will use a stack to store the nodes. We will traverse the tree in a level order traversal and store the nodes in the stack. Then we will pop the nodes and append the values to the result list.<br>
In this way, we will get the level order traversal in reverse order.<br>

For the second approach,<br>
We will find the maximum depth of the tree. Then we will append the values to the result list in a level order traversal. We will start from the maximum depth and go to the root of the tree.
In this way, we will get the level order traversal in reverse order.<br>

# Complexity Analysis
## Time Complexity(O(n))
- For the first approach, we are traversing the tree in a level order traversal, so the time complexity will be O(n).
- For the second approach, we are finding the maximum depth of the tree, and also the second recursion to fill the result list, so the time complexity will be O(n).

## Space Complexity(O(n))
- For the first approach, we are using a stack and queue to store the nodes, so the space complexity will be O(n).
- For the second approach, we use recursion, so the space complexity will be O(n).