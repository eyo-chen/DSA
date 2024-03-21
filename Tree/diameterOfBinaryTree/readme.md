# Problem Explanation

The core idea of this problem is to understand what's diameter<br>
The diameter of a binary tree is the length of the longest path between any two nodes in a tree<br>
This path may or may not pass through the root<br>

For example, the simple tree
```
   1
  / \
 2   3
```
The diameter of the above tree is 2, which is the length of the path [2, 1, 3]<br>
(Diamter have to exclude the root node 1)<br>

For example, the diameter of the below tree is 4<br>
```
   1
  / \
 2   3
    / \
   4   5
        \
         6
```
The diameter of the above tree is 4, which is the length of the path [2, 1, 3, 5, 6]<br>
(Diamter have to exclude the root node 1)<br>

But the path may not pass through the root, for example,
```
     1
    / \
   2   3
      / \
     4   5
    /     \
   6       7
  /         \
 8           9
```
The diameter of the above tree is 6, which is the length of the path [8, 6, 4, 3, 5, 7, 9]<br>
(Diamter have to exclude the root node 3)<br>

Now, the problem is How to find the diameter of a binary tree?<br>
In other words, how to find the longest path between any two nodes in a tree?<br>
Let's find out by look at above three examples<br>

For the first example,<br>
We know the longest path is [2, 1, 3], and the length of the path is 2<br>
But which node find this path?<br>
The answer is the node 1<br>
It seems simple, let's move to the second example<br>

For the second example,<br>
We know the longest path is [2, 1, 3, 5, 6], and the length of the path is 4<br>
But which node find this path?<br>
The answer is the node 1<br>
Why?<br>
Because when we're at node 1<br>
***The sum of the longest path from the left subtree and the longest path from the right subtree*** is 4<br>
And this 4 is the best value for each node<br>

For the third example,<br>
We know the longest path is [8, 6, 4, 3, 5, 7, 9], and the length of the path is 6<br>
But which node find this path?<br>
The answer is the node 3<br>
Why?<br>
Because when we're at node 3<br>
***The sum of the longest path from the left subtree and the longest path from the right subtree*** is 6<br>
And this 6 is the best value for each node<br>

Can we find a pattern?<br>
Yes, we can<br>
The pattern is<br>
For each node in the tree, <br>
We want to find the longest path of left subtree<br>
And the longest path of right subtree<br>
Then, sum up both of them<br>
And see is this sum is the best value for each node<br>

Let's summarize the above pattern<br>
1. Find the longest path of the left subtree
2. Find the longest path of the right subtree
3. Sum up both of them
4. And see is this sum is the best value for each node<br>

Let's walk through the process<br>
```
     1
    / \
   2   3
      / \
     4   5
```
Callstack1<br>
At root node 1<br>
(1)Go to the left subtree<br>

Callstack2<br>
At node 2<br>
(1)Go to the left subtree<br>

Callstack3<br>
At null node<br>
Return 0<br>

Back to Callstack2<br>
At node 2<br>
(1)Go to the left subtree: 0<br>
(2)Go to the right subtree<br>

Callstack3<br>
At null node<br>
Return 0<br>

Back to Callstack2<br>
At node 2<br>
(1)Go to the left subtree: 0<br>
(2)Go to the right subtree: 0<br>
(3) Sum up both of them: 0 + 0 = 0<br>
(4) And see is this sum is the best value for each node<br>
The best value for node 2 is 0<br>

Back to Callstack1<br>
At root node 1<br>
(1)Go to the left subtree: 1<br>
(Because at previous callstack, we return 0 + 1)<br>
(2)Go to the right subtree<br>

Callstack2<br>
At node 3<br>
(1)Go to the left subtree<br>

Callstack3<br>
At node 4<br>
(1)Go to the left subtree<br>

Callstack4<br>
At null node<br>
Return 0<br>

Back to Callstack3<br>
At node 4<br>
(1)Go to the left subtree: 0<br>
(2)Go to the right subtree<br>

Callstack4<br>
At null node<br>
Return 0<br>

Back to Callstack3<br>
At node 4<br>
(1)Go to the left subtree: 0<br>
(2)Go to the right subtree: 0<br>
(3) Sum up both of them: 0 + 0 = 0<br>
(4) And see is this sum is the best value for each node<br>
The best value for node 4 is 0<br>

Back to Callstack2<br>
At node 3<br>
(1)Go to the left subtree: 1<br>
(2)Go to the right subtree<br>

Callstack3<br>
At node 5<br>
(1)Go to the left subtree<br>

Callstack4<br>
At null node<br>
Return 0<br>

Back to Callstack3<br>
At node 5<br>
(1)Go to the left subtree: 0<br>
(2)Go to the right subtree<br>

Callstack4<br>
At null node<br>
Return 0<br>

Back to Callstack3<br>
At node 5<br>
(1)Go to the left subtree: 0<br>
(2)Go to the right subtree: 0<br>
(3) Sum up both of them: 0 + 0 = 0<br>
(4) And see is this sum is the best value for each node<br>
The best value for node 5 is 0<br>

Back to Callstack2<br>
At node 3<br>
(1)Go to the left subtree: 1<br>
(2)Go to the right subtree: 1<br>
(3) Sum up both of them: 1 + 1 = 1<br>
(4) And see is this sum is the best value for each node<br>
The best value for node 3 is 2<br>

Back to Callstack1<br>
At root node 1<br>
(1)Go to the left subtree: 1<br>
(2)Go to the right subtree: 2<br>
(3) Sum up both of them: 1 + 2 = 3<br>
(4) And see is this sum is the best value for each node<br>
The best value for node 1 is 3<br>

The best value for each node is<br>
At node1, we know the best value is 3<br>

The diameter of the above tree is 3<br>




