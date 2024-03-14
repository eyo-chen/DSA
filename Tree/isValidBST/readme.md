# Problem Explanation

The core idea to solve this problem is truely understand the definition of a binary search tree<br>
A binary search tree must satisfy the following conditions<br>
1. The ***left subtree*** of a node contains only nodes with keys ***less than*** the node's key
2. The ***right subtree*** of a node contains only nodes with keys ***greater than*** the node's key
3. Both the left and right subtrees must also be binary search trees

Let's take a look at some examples<br>
```
    2
   / \
  1   3
```
The above tree is a binary search tree<br>
Because left(1) < root(2) < right(3)<br>

```
    5
   / \
  1   2
     / \
    3   6
```
The above tree is not a binary search tree<br>
Because right(2) < root(5)<br>
Right node should be greater than the root node<br>

```
    10
   /  \
  5   15
     /  \
    6   20
```
The above tree is not a binary search tree<br>
At first, it looks like a binary search tree<br>
But it's not<br>
The gotcha is the node(6)<br>
If we only focus on the right subtree(6,15,20), it's a binary search tree<br>
But the root node(10) is greater than the right node(6)<br>
Recall the second point of the definition<br>
The ***right subtree*** of a node contains only nodes with keys ***greater than*** the node's key<br>
Any node in the right subtree should be greater than the root node<br>
So this tree is not a binary search tree<br>

```
    10
   /  \
  5   15
     /  \
    12  20
```
The above tree is a binary search tree<br>


After understanding the definition of a binary search tree, the problem becomes much easier<br>
We just traverse the whole tree, and check if the current node is satisfying the definition<br>

How can we know if the current node is satisfying the definition?<br>
Let's walk through the thought process, and use the following tree as an example<br>
```
    10
   /  \
  5   15
     /  \
    12  20
```
For root node(10), is it satisfying the definition?<br>
Yes, because root node has no limited value, it can be any value<br>

For left node(5), is it satisfying the definition?<br>
Yes, because it's less than the root node(10)<br>

For right node(15), is it satisfying the definition?<br>
Yes, because it's greater than the root node(10)<br>

For left node(12), is it satisfying the definition?<br>
Yes, because it's greater than the root node(10), and also less than the right node(15)<br>
Here is the key point<br>
The current node has to be in a range<br>
The range is (10, 15)

For right node(20), is it satisfying the definition?<br>
Yes, because it's greater than the root node(10), and also greater than the left node(15)<br>

Can you see the pattern?<br>
For any given node, it has to be in a range<br>
The range is (min, max)<br>
At root node, the range is (-inf, inf), because root node can by any value<br>
If we go left, the range becomes (min, root)<br>
We update the upper bound to the root node<br>
Because the left node has to be less than the root node<br>
If we go right, the range becomes (root, max)<br>
We update the lower bound to the root node<br>
Because the right node has to be greater than the root node<br>
And this is the core idea to solve the problem<br>

Let's walk through the process
```
    10(-inf, inf)
   /  \
  5   15
     /  \
    12  20
```
At root node(10), the range is (-inf, inf)<br>
We go left<br>

```
            10(-inf, inf)
           /  \
(-inf, 10) 5   15
              /  \
             12  20
```
At left node(5), the range is (-inf, 10)<br>
There's no left or right node, so we go back to the root node<br>
We go right<br>

```
              10(-inf, inf)
             /  \
(-inf, 10)  5   15 (10, inf)
               /  \
              12  20
```
At right node(15), the range is (10, inf)<br>
We go left<br>

```
              10(-inf, inf)
             /  \
(-inf, 10)  5   15 (10, inf)
               /  \
    (10, 15)  12  20
```
At left node(12), the range is (10, 15)<br>
There's no left or right node, so we go back to the root node<br>
We go right<br>

```
              10(-inf, inf)
             /  \
(-inf, 10)  5   15 (10, inf)
               /  \
    (10, 15)  12  20 (15, inf)
```
At right node(20), the range is (15, inf)<br>
There's no left or right node, so we go back to the root node<br>
We're done<br>


# Complexity Analysis
## Time Complexity: O(N)
- We visit every node once

## Space Complexity: O(h)
- The space complexity is O(h) because we're using the call stack