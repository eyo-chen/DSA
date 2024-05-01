# Problem Explanation

This problem is very similar to the connect problem.<br>
The only difference is that the tree is not perfect balanced tree.<br>

In the connect problem, we come up with three different solutions.<br>
Let's examine each of them and see if they can be applied to this problem.<br>

## Using Queue
This is the most general solution for this problem.<br>
It uses BFS to iterate through the tree.<br>
It can solve the problem for any kind of tree.<br>
But it costs O(n) space complexity.<br>

The code is as same as the connect problem.<br>

### Complexity Analysis
#### Time Complexity(O(n))
- We visit each node once.

#### Space Complexity(O(n))
- We use a queue to implement the BFS.

## Using Recursion
We can use recursion to solve this problem.<br>
But we have to improve the solution a bit.<br>

Imagine a tree is following
```
    1
   / \
  2   3
 / \   \
4   5   6
```
When we visit the node 2, we can't connect 5 and 6.<br>
In the first connect problem, we can simply say `2->right->next = 2->next->left`.<br>
But in this case, we can't do that.<br>

So, we have to come up with a solution to find the next node to connect.<br>
We can simply design a `findNextNode` helper function.<br>
The idea of this problem is as follows:
- It takes a node as an argument.
- It returns the next existing node in the next level.

Imagine the tree is following
```
                 1
              /     \
            2         3
          /   \         \
        4       5         6
       /                   \
      7                     8
```
Suppose we're at node 4.<br>
We can simply pass the node 5 to the `findNextNode` function.<br>
Then, the function will return the node 8.<br>

So, `findNextNode` returns the next existing node in the next level.<br>

The detail implementation is as follows:
- If the node is nullptr, return nullptr.
- If the node has left child, return the left child.
- If the node has right child, return the right child.
- Go to the next node and repeat the process.

By using this helper function, we can solve this problem using recursion.<br>


One final thing to note when using recursion is that we need to wire-up the right child first.<br>
If we wire-up the left child first, we can't find the next node to connect.<br>
Imaing the following tree
```
                                           2
                                    /            \
                                  1                3
                                /     \         /     \
                              0         7      9       1
                             /         / \             / \
                            2         1   0           8   8
                                         /              
                                        7               
```
Say we're at node 7.<br>
We want to connect it's right child to the next node. (0 -> 8)<br>
Because we wire-up the left child first, 9 and 1 are NOT connected at this point.<br>
So we can't find the next node to connect.<br>

So, we need to wire-up the right child first.<br>
Then, we can find the next node to connect.<br>

Let's summarize the solution:
- When root is nullptr, return nullptr.(base case)
- When left child exists,
  - If right child exists, wire-up the left child's next pointer to the right child.
  - If right child doesn't exist, call `findNextNode` on the next node to find the next existing node to wire-up.
- When right child exists,
  - call `findNextNode` on the next node to find the next existing node to wire-up.
- Call the reucrsive function for the right child.
- Call the recursive function for the left child.

Let's use this example to walk through the code.<br>
```
    1
   / \
  2   3
 / \   \
4   5   6
```
First callstack
- root = 1
- root->left = 2
  - wire-up 2->3
- root->right = 3
  - wire-up 3->nullptr
- call connect(3)

Second callstack
- root = 3
- root->left = nullptr
- root->right = 6
  - wire-up 6->nullptr
- pop from the callstack

First callstack
- root = 1
- root->left = 2
  - wire-up 2->3
- root->right = 3
  - wire-up 3->nullptr
- call connect(3)
- call connect(2)

Second callstack
- root = 2
- root->left = 4
  - wire-up 4->5
- root->right = 5
  - wire-up 5->6
- pop from the callstack(skip further callstack)

First callstack
- root = 1
- root->left = 2
  - wire-up 2->3
- root->right = 3
  - wire-up 3->nullptr
- call connect(3)
- call connect(2)
- pop from the callstack

### Complexity Analysis
#### Time Complexity O(n)
- We visit each node once.
- For each node, it performs a constant amount of work, except for the calls to findNextNode, which could potentially traverse up the tree.
- However, the findNextNode function, while called multiple times, doesn't visit any node more than once due to the nature of how the next pointers are set up (it only follows the next pointers, and since we connect right nodes before left nodes, there is no way it would visit a node more than once for a given depth).
- Therefore, the overall time complexity is O(N), where N is the number of nodes in the tree.

#### Space Complexity O(H)
- The space complexity of the solution is determined by the height of the tree due to the recursive nature of the connect method.
- In the worst case (when the tree is completely unbalanced), the height of the tree can be N, which would make the space complexity O(N) due to the call stack.
- For a balanced tree, the height would be log(N), and the space complexity would be O(log(N)).



## Using Iteration with O(1) space complexity
The overall concept of this approach is similar to the connect problem.<br>
However, there're some differences.<br>

First, as we discussed in the recursion approach, we need to use the `findNextNode` helper function to find the next existing node in the next level.<br>
It's basically the same as the recursion approach.<br>
- If the left subtree exists,
  - If the right subtree exists, wire-up the left subtree's next pointer to the right subtree.
  - If the right subtree doesn't exist, call `findNextNode` on the next node to find the next existing node to wire-up.
- If the right subtree exists,
  - call `findNextNode` on the next node to find the next existing node to wire-up.

The second difference is how to update `leftMostNode`.<br>
In the first problem, we can always assume that there's a left subtree on every node.<br>
However, in this problem, we can't assume that.<br>
Suppose the following tree
```
              1
            /   \
          2       3
            \       \
              5       7
```
At the third level, the `leftMostNode` is 5.<br>
How to find the `leftMostNode` for each level?<br>
We can leverage the `findNextNode` helper function.<br>
We simply pass the current `leftMostNode` to the `findNextNode` function.<br>
Then, the function will return the next `leftMostNode`.<br>
Becauses `findNextNode` returns the next existing node in the next level.<br>

And because of that, the inner while-loop condition is different too.<br>
```c++
while (leftMostNode != nullptr)
```
Because `leftMostNode->left` might be nullptr, so we just check the `leftMostNode`.<br>

Let's use the following example to walk through the code.<br>
```
              1
            /   \
          2       3
            \       \
              5       7
```
Initially, `leftMostNode` is 1.<br>

First outer while-loop
- `leftMostNode` is 1.
- `curNode` is 1.
- In the first inner while-loop
  - `curNode` is 1.
  - `curNode->left` is 2.
  - `curNode->right` is 3.
  - wire-up 2->3.
  - update `curNode` to nullptr
- break the inner while-loop
- update `leftMostNode` to 2.

Second outer while-loop
- `leftMostNode` is 2.
- `curNode` is 2.
- In the first inner while-loop
  - `curNode` is 2.
  - `curNode->left` is nullptr.
  - `curNode->right` is 5.
  - wire-up 5->7.
  - update `curNode` to 3.
- In the second inner while-loop
  - `curNode` is 3.
  - `curNode->left` is nullptr.
  - `curNode->right` is 7.
  - update `curNode` to nullptr.
- break the inner while-loop
- update `leftMostNode` to 5.

Third outer while-loop
- `leftMostNode` is 5.
- `curNode` is 5.
- In the first inner while-loop
  - `curNode` is 5.
  - `curNode->left` is nullptr.
  - `curNode->right` is nullptr.
  - update `curNode` to 7.
- In the second inner while-loop
  - `curNode` is 7.
  - `curNode->left` is nullptr.
  - `curNode->right` is nullptr.
  - update `curNode` to nullptr.
- break the inner while-loop
- update `leftMostNode` to nullptr.
- break the outer while-loop.

### Complexity Analysis
#### Time Complexity O(n)
- We visit each node once.

#### Space Complexity O(1)
- We don't use any extra space.