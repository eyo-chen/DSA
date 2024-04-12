# Problem Explanation

The key idea of this problem is to truly understand the problem<br>
The problem said
```
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.
```

```
        3
       / \
      9  20
         /  \
        15   7
```
The minimum depth of the tree is 2, because the shortest path is 3->9

```
        2
       / \
      x   3
           \
            4
             \
              5
```
The minimum depth of the tree is 4, because the shortest path is 2->3->4->5<br>
It's not 2 because 2 is not a leaf node<br>

## Using Queue
Using queue to find the minimum depth is the most straightforward way to solve this problem

The idea is
1. Put the root node into the queue as starting point
2. Going into while-loop, check the size of the queue
3. Iterate over the current size of the queue
4. Pop the first element of the queue
5. Check if the left and right child of the node is null, if it is null, then we can return the depth of the node
   - Because it means we find the leaf node, and the minimum depth of tree
6. If the left and right child of the node is not null, then we can add the left and right child of the node to the queue

The key here is that we need a for-loop inside while-loop<br>
Why is that?<br>
The main reason is that it's the only way to keep track of the depth of each level<br>
Without this, there's no way to know what's the depth of the current node<br>
For example,<br>
```
        3
      /   \
     9    20
    / \   / \
   15  7  3  4
```
At first, queue is [3], and we try to use `height` to keep track of the depth of the node<br>
Then we iterate over the queue<br>
Inside first while-loop, we pop the first element of the queue, which is 3<br>
Then we check the left and right child of the node, the qeueu is [9, 20]<br>
Now, we increment the height by 1, and we know that the height of the node is 1<br>

Then we go to the next iteration of the while-loop<br>
Inside the second while-loop, we pop the first element of the queue, which is 9<br>
Then we check the left and right child of the node, the queue is [20, 15, 7]<br>
Now, we increment the height by 1, and we know that the height of the node is 2<br>

Then we go to the next iteration of the while-loop<br>
Inside the third while-loop, we pop the first element of the queue, which is 20<br>
Then we check the left and right child of the node, the queue is [15, 7, 3, 4]<br>
Now, we increment the height by 1, and we know that the height of the node is 3<br>

Here is wrong, because the height of the node 20 is 2, not 3<br>
That's why we can't accurately keep track of the depth of the node<br>


Whenever we're in the beginning of the while-loop,<br>
The queue at this state represents all the nodes at the same level<br>
Let's see the example below<br>
```
        3
      /   \
     9    20
    / \   / \
   15  7  3  4
```
The queue at the beginning of the while-loop is [3]<br>
It means we only have one node 3 at this level<br>
We know the we need only one for-loop<br>
We add the left and right child of the node to the queue, which is [9, 20]<br>
Then we go to the next iteration of the while-loop<br>

The queue at this state is [9, 20]<br>
It means we have two nodes 9 and 20 at this level<br>
We know that we need two for-loop<br>
Inside for-loop, we add the left and right child of the node to the queue in order, which is [15, 7, 3, 4]<br>
Then we go to the next iteration of the while-loop<br>

The queue at this state is [15, 7, 3, 4]<br>
It means we have four nodes 15, 7, 3, 4 at this level<br>
We know that we need four for-loop<br>

And so on<br>

Along with the process, as long as we find the first leaf node, we can return the depth of the node<br>

### Complexity Analysis
#### Time Complexity O(n)
- We need to traverse all the nodes in the tree

#### Space Complexity O(n)
- In the worst case, the queue can contain all the nodes in the tree

## Using Queue with Struct
Recall from the previous explanation, we need a for-loop inside while-loop to keep track of the depth of the node<br>
But the problem is that we can't accurately keep track of the depth of the node<br>

What if we can keep track of the depth of the node in the queue?<br>
Then we don't need to use for-loop inside while-loop<br>

The core idea to achieve this is to use struct<br>
We can create a struct that contains the node and the depth of the node<br>
So we know the depth for each node<br>

The idea is
1. Construct a struct that contains the node and the depth of the node
2. Put the struct root node into the queue as starting point
3. Going into while-loop
4. Pop the first element of the queue
5. Check if the left and right child of the node is null, if it is null, then we can return the depth of the node
    - Because it means we find the leaf node, and the minimum depth of tree
6. If the left and right child of the node is not null, then we can add the left and right child of the node to the queue

The core logic is similar to the previous one<br>
The key difference here is that we don't need for-loop inside while-loop<br>
Because we can use struct to give us the information of the depth of the node<br>

For example, <br>
```
        3
      /   \
     9    20
    / \   / \
   15  7  3  4
```
The queue at the beginning of the while-loop is [{3,1}]<br>
It means we only have one node 3 at this level, and the depth of the node is 1<br>
We add the left and right child of the node to the queue, which is [{9,2}, {20,2}]<br>
Then we go to the next iteration of the while-loop<br>

The queue at this state is [{9,2}, {20,2}]<br>
We pop the first element of the queue, which is {9,2}<br>
It means the node is 9, and the depth of the node is 2<br>
Then we add the left and right child of the node to the queue, which is [{20,2}, {15,3}, {7,3}]<br>
Then we go to the next iteration of the while-loop<br>

The queue at this state is [{20,2}, {15,3}, {7,3}]<br>
We pop the first element of the queue, which is {20,2}<br>
It means the node is 20, and the depth of the node is 2<br>
Then we add the left and right child of the node to the queue, which is [{15,3}, {7,3}, {3,3}, {4,3}]<br>
Then we go to the next iteration of the while-loop<br>

So on<br>

### Complexity Analysis
#### Time Complexity O(n)
- We need to traverse all the nodes in the tree

#### Space Complexity O(n)
- In the worst case, the queue can contain all the nodes in the tree

## Using Recursion
Using recursion to find the minimum depth is another way to solve this problem<br>

There are four different senarios to consider<br>
1. Both left and right child of the root is null
```
        3
       / \
      x   x
```
In this case, we simply return 1 because the root is the leaf node<br>
`return 1`

2. The left child of the root is null
```
        3
       / \
      x   20
           ...
```
In this case, we can't consider the depth of the left node<br>
Because there's no way left node can be the leaf node<br>
So we just simply return the minimum depth of the right child of the root plus 1<br>
`return rightDepth + 1`

3. The right child of the root is null
```
        3
       / \
      9   x
    ...
```
In this case, we can't consider the depth of the right node<br>
Because there's no way right node can be the leaf node<br>
So we just simply return the minimum depth of the left child of the root plus 1<br>
`return leftDepth + 1`

4. Both left and right child of the root is not null
```
        3
       / \
      9   20
    ...   ...
```
In this case, we need to consider both left and right child of the root<br>
So we return the minimum depth of the left and right child of the root plus 1<br>
`return min(leftDepth, rightDepth) + 1`

Let's walk through the example<br>
```
        3
       / \
      9  20
        /  \
       15   7
```
callstack 1<br>
- root = 3
- go left

callstack 2<br>
- root = 9
- left is null, right is null
- return 1

callstack 1<br>
- root = 3
- left = 1
- go right

callstack 2<br>
- root = 20
- go left

callstack 3<br>
- root = 15
- left is null, right is null
- return 1

callstack 2<br>
- root = 20
- left = 1
- go right

callstack 3<br>
- root = 7
- left is null, right is null
- return 1

callstack 2<br>
- root = 20
- left = 1
- right = 1
- both left and right is not null
- return min(1,1) + 1 = 2

callstack 1<br>
- root = 3
- left = 1
- right = 2
- both left and right is not null
- return min(1,2) + 1 = 2


Note that there are short version of the code<br>
But I prefer to stick this one because it's more readable<br>

### Complexity Analysis
#### Time Complexity O(n)
- We need to traverse all the nodes in the tree
- Note that even though the time complexity is the same as using queue. It's less efficient than using queue
- Imaging the tree is following
```
        3
       /  \
      4    10
          /   \
         x     20
              /   \
             x     30
                  /   \
                 x     40
                        /   \
                       x     50
```
- When using recursion, we need to traverse all the nodes in the tree
- But when using queue, we can return the depth of the node as soon as we find the leaf node

#### Space Complexity O(n)
- In the worst case, the callstack can go as deep as the height of the tree


