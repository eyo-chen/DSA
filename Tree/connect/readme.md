# Problem Explanation

This problem is not that hard if question doesn't ask for implementing the solution in O(1) space complexity.<br>
Note that this problem assume that the tree is perfect balanced tree.<br>

We'll walk through the three different solutions for this problem.

1. Using Queue
 - Still solve the problem if it's not perfect balanced tree.
 - Use O(n) space complexity.

2. Using Recursion
  - Assume the tree is perfect balanced tree.
  - Use O(n) space complexity, but the problem said stack doesn't count as extra space.

3. Using Iteration with O(1) space complexity
  - Assume the tree is perfect balanced tree.
  - Use O(1) space complexity.


## Using Queue
This is the most general solution for this problem.<br>
Even though it costs O(n) space complexity, it can solve the problem for any kind of tree.<br>

The idea is simple.<br>
What we're asked is to "do something" for each level of the tree.<br>
So, using BFS is the most natural way to solve this problem.<br>

We just simply using queue.<br>
For each level, we wire-up the next pointer of each node.<br>

The small variation is using `preNode`<br>
For each level, we use `prevNode->next = curNode` to wire-up the next pointer.<br>
For the first node of each level, there's no `prevNode`.<br>
So we have if-statements to handle this case.<br>
```c++
if (preNode != nullptr) preNode->next = curNode;
```
After each iteration, we just make sure to update `preNode`<br>

Suppose the tree is following
```
    1
   / \
  2   3
 / \ / \ 
4  5 6  7
```
First while-loop
- queue: [1]
- push 2, 3 to the queue

Second while-loop
- queue: [2, 3]
- First for-loop
  - curNode = 2
  - preNode = nullptr
  - push 4, 5 to the queue
  - preNode is nullptr, so no wiring-up
  - update preNode to 2
- Second for-loop
  - curNode = 3
  - preNode = 2
  - push 6, 7 to the queue
  - wire-up 2->3

Third while-loop
- queue: [4, 5, 6, 7]
- First for-loop
  - curNode = 4
  - preNode = nullptr
  - push nothing to the queue
  - preNode is nullptr, so no wiring-up
  - update preNode to 4
- Second for-loop
  - curNode = 5
  - preNode = 4
  - push nothing to the queue
  - wire-up 4->5
- Third for-loop
  - curNode = 6
  - preNode = 5
  - push nothing to the queue
  - wire-up 5->6
- Fourth for-loop
  - curNode = 7
  - preNode = 6
  - push nothing to the queue
  - wire-up 6->7

### Complexity Analysis
#### Time Complexity O(n)
- We visit each node once.
- So, the time complexity is O(n)

#### Space Complexity O(n)
- We use queue to store the nodes.

## Using Recursion
When dealing with recursion, we need to think about three different things.
1. What's the base case?
2. How to wire-up the next pointer?
3. How to recursively call the function?

For the base case, it's simple.<br>
If the root is nullptr, we just return nullptr.<br>

For the wiring-up the next pointer, we need to think about two different cases.<br>
1. How to wire-up the next pointer of the left child?
2. How to wire-up the next pointer of the right child?

For the left child, it's simple.<br>
We just wire-up the left child's next pointer to the right child.<br>
```c++
root->left->next = root->right;
```
```
    1
   / \
  2 -> 3
```

For the right child, it's a bit tricky.<br>
We need to wire-up the right child's next pointer to the root's next's left child.<br>
```c++
root->right->next = root->next->left;
```
```
  2 -> 3
 / \   / \
4  5->6   7
```
`root->right` is 5, `root->next` is 3, and `root->next->left` is 6.<br>

For this case, we need to make sure that `root->next` is not nullptr.<br>
Suppose we're at 3, `3->next->left` will cause the segmentation fault.<br>
So that's why we need to make sure that `root->next` is not nullptr.<br>
```c++
if (root->right != nullptr && root->next != nullptr)
```


For the recursive call, it's simple.<br>
We just call the function for the left child and the right child.<br>
```c++
root->left = connect(root->left);
root->right = connect(root->right);
```
It simply means "go to connect the left child tree and the right child tree".

Suppose the tree is following
```
    1
   / \
  2   3
 / \ / \ 
4  5 6  7
```
First callstack
- root = 1
- root->left = 2
  - wire-up 2->3
- root->next = nullptr
 - no wiring-up for right child
- call connect(2)

Second callstack
- root = 2
- root->left = 4
  - wire-up 4->5
- root->next = 3
  - wire-up 5->6(2->right->next = 2->next->left)
- call connect(4)

Third callstack
- root = 4
- root->left = nullptr
- root->right = nullptr

Second callstack
- root = 2
- root->left = 4
  - wire-up 4->5
- root->next = 3
  - wire-up 5->6(2->right->next = 2->next->left)
- call connect(4)
- call connect(5)

Third callstack
- root = 5
- root->left = nullptr
- root->right = nullptr

Second callstack
- root = 2
- root->left = 4
  - wire-up 4->5
- root->next = 3
  - wire-up 5->6(2->right->next = 2->next->left)
- call connect(4)
- call connect(5)
- pop from the callstack

First callstack
- root = 1
- root->left = 2
  - wire-up 2->3
- root->next = nullptr
 - no wiring-up for right child
- call connect(2)
- call connect(3)

Second callstack
- root = 3
- root->left = 6
  - wire-up 6->7
- root->next = nullptr
  - no wiring-up for right child
- call connect(6)

Third callstack
- root = 6
- root->left = nullptr
- root->right = nullptr

Second callstack
- root = 3
- root->left = 6
  - wire-up 6->7
- root->next = nullptr
  - no wiring-up for right child
- call connect(6)
- call connect(7)

Third callstack
- root = 7
- root->left = nullptr
- root->right = nullptr

Second callstack
- root = 3
- root->left = 6
  - wire-up 6->7
- root->next = nullptr
  - no wiring-up for right child
- call connect(6)
- call connect(7)
- pop from the callstack

First callstack
- root = 1
- root->left = 2
  - wire-up 2->3
- root->next = nullptr
 - no wiring-up for right child
- call connect(2)
- call connect(3)
- pop from the callstack

Note that this approach is only valid for the perfect balanced tree.<br>

### Complexity Analysis
#### Time Complexity O(n)
- We visit each node once.
- So, the time complexity is O(n)

#### Space Complexity O(n)
- We use callstack to implement the recursion.


## Using Iteration with O(1) space complexity
Although this is the most efficient solution, it's a bit tricky.<br>

The idea is<br>
We just go deep to the leftmost node.<br>
And then, we wire-up the next pointer for each level.<br>
For example,
```
    1
   / \
  2   3
 / \ / \
4  5 6  7
```
We start from 1.<br>
Go to 2, and wire-up 2->3.<br>
Go to 4, and wire-up 4->5->6->7.<br>

We need two while-loops.<br>
First while-loop is for each level.<br>
Second while-loop is for each node in the level.<br>

We need two pointers.<br>
One is `leftMostNode` to keep track of the leftmost node of each level.<br>
One is `curNode` to keep track of the current node of each level.<br>

How to wire-up the next pointer of left child?<br>
It's simple.<br>
```c++
curNode->left->next = curNode->right;
```

How to wire-up the next pointer of right child?<br>
It's a bit tricky.(as same as the recursion)<br>
```c++
curNode->right->next = curNode->next->left;
```
But, we need to make sure that `curNode->next` is not nullptr.<br>
```c++
if (curNode->next != nullptr)
```

Let's directly walk through the example.<br>
Suppose the tree is following
```
    1
   / \
  2   3
 / \ / \
4  5 6  7
```
First outer while-loop
- leftMostNode = 1
- curNode = 1
- First inner while-loop
  - wire-up 2->3
  - no need to wire-up 3->nullptr
  - curNode = nullptr (1->next)
  - break the inner while-loop
- leftMostNode = 2 (1->left)

Second outer while-loop
- leftMostNode = 2
- curNode = 2
- First inner while-loop
  - wire-up 4->5
  - wire-up 5->6
  - curNode = 3 (2->next)
- Second inner while-loop
  - wire-up 6->7
  - no need to wire-up 7->nullptr
  - curNode = nullptr (3->next)
  - break the inner while-loop
- leftMostNode = 4 (2->left)

Third outer while-loop
- leftMostNode = 4
- because 4->left is nullptr, break the outer while-loop

### Complexity Analysis
#### Time Complexity O(n)
- We visit each node once.

#### Space Complexity O(1)
- We don't use any extra space.
