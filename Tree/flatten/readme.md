# Problem Explanation

There're two approaches to solve this problem:
1. Using vector to store the pre-order node values and then re-construct the tree.
2. Build the tree recursively.

## Approach 1
This is the most intuitive way to solve this problem.<br>

We first traverse the tree in pre-order and store the node values in a vector. Then we re-construct the tree using the vector.

Just look the code to understand the idea.<br>

### Complexity Analysis
#### Time Complexity(O(n))
- For `genPreOrderNodes`, we need to traverse the tree in pre-order, which takes O(n) time.
- For the for-loop, it also takes O(n) time.

#### Space Complexity(O(n))
- For `genPreOrderNodes`, we need to store the node values in a vector, which takes O(n) space.
- For storing the nodes in vector, it also takes O(n) space.

## Approach 2
The second approach is more efficient, but it's a little bit tricky.<br>
Consider the following tree
```
      1
     / \
    2   5 
   / \   \
  3   4   6
```
And let's focus on this subtree only
```
    2
   / \
  3   4
```
How can we make this tree to be linked-list like structure?
```
    2
     \
      3
       \
        4
```
We have three steps to do this:
1. 3.right = 4
```
    2
   / \
  3   4
   \
    4
```
2. 2.right = 3
```
    2
   / \
  3   3
   \   \
    4   4
```
3. 2.left = null
```
    2
   / \
  x   3
       \
        4
```

Let's imagine that 3 and 4 already have been converted to linked-list like structure.
```
      2
    /   \
   3     4
    \     \
     x     x
      \     \
       x     x
        \     \
         x     x
```
How should we connect 2 with 3 and 4?<br>
Recall the three steps above, we can do the similar thing to connect 2 with 3 and 4.
1. 3.right = 4
2. 2.right = 3
3. 2.left = null

Let's examine each steps, and see how to change it a little bit to make it work for the whole tree.
1. 3.right = 4
This is obvious wrong<br>
What we want is the last node of 3's right subtree to point to 4.<br>
The correct way is<br>
`the_last_node_of_3_right_subtree.right = 4`<br>
which is<br>
`the_tail_of_left_subtree.right = the_root.right`

2. 2.right = 3
This is almost correct<br>
But let's clarify it a little bit<br>
What we want is the root of right subtree to be the left child of the root.<br>
`the_root.right = the_root.left`

3. 2.left = null<br>
This is 100% correct<br>
We want to make sure the left child of the root to be null.<br>

Now, the problem is what shoud we return?<br>
```
    x
     \
      x
       \
        x 
```
When there's right child tail, we should return the right child tail.<br>


```
     x
    /
   x
  /
 x
```
When there's no right child tail, we should return the left child tail.<br>

```
    x
```
When there's no left child tail, we should return the root itself.<br>

Let's walk through the process of converting the tree to linked-list like structure.
```
      1
     / \
    2   5 
   / \   \
  3   4   6
```
callstack 1
- root = 1
- go to left subtree

callstack 2
- root = 2
- go to left subtree

callstack 3
- root = 3
- left tree is null
- right tree is null
- return root(3)

callstack 2
- root = 2
- leftTail = 3
- go to right subtree

callstack 3
- root = 4
- left tree is null
- right tree is null
- return root(4)

callstack 2
- root = 2
- leftTail = 3
- rightTail = 4
- leftTail is not null, so do the wiring operation
- after wiring,
```
    2
     \
      3
       \
        4
```
- return rightTail(4)

callstack 1
- root = 1
- leftTail = 4
- go to right subtree

callstack 2
- root = 5
- leftTail = null
- go to right subtree

callstack 3
- root = 6
- left tree is null
- right tree is null
- return root(6)

callstack 2
- root = 5
- leftTail = null
- rightTail = 6
- leftTail is null, so do nothing
- return rightTail(6)

callstack 1
- root = 1
- leftTail = 4
- rightTail = 6
- leftTail is not null, so do the wiring operation
- after wiring,
```
      1
       \
        2
         \
          3
           \
            4
             \
              5
               \
                6
```
Done!

### Complexity Analysis
#### Time Complexity(O(n))
- We visit each node once, so the time complexity is O(n).

#### Space Complexity(O(h))
- The space complexity is O(h), where h is the height of the tree.