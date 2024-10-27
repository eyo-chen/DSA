# Problem Explanation

The main point of solving this problem is to ***NOT TRY TO MANIPULATE THE MEMORY ADDRESS*** of the deleted node.<br>
Just change the value of the deleted node, and then delete the replaced node.<br>

Consider the following tree<br>
```
                                       12
                            5                      15
                      3         7            13         17
                  1                9             14          20
                                8     11                  18
```
If we want to delete 15, we have to do the following steps:<br>
1. Find the node with the value 15.
2. Find the node to replace the deleted node. In this case, the node with the value 17.
3. Change the value of the node with the value 15 to 17.
4. Delete the node with the value 17.

Let's walk through the above steps:<br>
1. Find the node with the value 15.
```
                                       12
                            5                      15  <--- deleted node
                      3         7            13         17
                  1                9             14           20
                                8     11                  18
```

2. Find the node to replace the deleted node. In this case, the node with the value 17.
```
                                       12
                            5                      15  <--- deleted node
                      3         7            13         17  <--- node to replace the deleted node
                  1                9             14          20
                                8     11                  18
```

3. Change the value of the node with the value 15 to 17.
```
                                       12
                            5                      17  <--- deleted node
                      3         7            13         17  <--- node to replace the deleted node
                  1                9             14          20
                                8     11                  18
```

4. Delete the node with the value 17.
```
                                       12
                            5                      17
                      3         7            13         20
                  1                9             14   18
                                8     11
```

Note that we don't need to move the node 17 up to the position of the node 15. We just need to change the value of the node 15 to 17, and then delete the node 17.<br>
It means that don't try to manipulate the memory address of the deleted node.<br>
Because it's way more complicated and not necessary.<br>

Now, let's come up with the detailed steps to solve this problem<br>
Consider the same tree as above, and the value to delete is 15.<br>
```
                                       12
                            5                      15
                      3         7            13         17
                  1                9              14        20
                                8    11                  18
```
When we're at the root node 12, Do we need to go to the left subtree?<br>
No, because the value to delete is 15, and 15 is greater than 12.<br>
Because of the property of the binary search tree, we can guarantee that node 15 is in the right subtree.<br>
That's the first point to solve this problem.<br>
- If the value to delete is ***less than*** the current node, go to the ***left subtree*** to find the node to delete.
- If the value to delete is ***greater than*** the current node, go to the ***right subtree*** to find the node to delete.

After finding the node to delete, how to delete the node?<br>
There are three cases to consider<br>
1. The node to delete is a leaf node.
2. The node to delete has only one child.
3. The node to delete has two children.

Let's consider the first case.<br>
```
                                       12
                            5                      15
                      3         7            13         17
  deleted node --> 1                9            14          20
                                8      11                18
```
If we want to delete node 1, we should we do?<br>
We simply return nullptr to the parent node(3).<br>
```
                                       12
                            5                      15
                      3         7            13         17
                  x                9             14         20
                                8    11                  18
```
Therefore, when the node to delete is a leaf node, we simply return nullptr to the parent node.<br>

Let's consider the second case.<br>
```
                                       12
                            5                              15
                      3         7  <-- deleted node   13         17
                  1                9                      14          20
                                8     11                           18
```
If we want to delete node 7, what should we do?<br>
We should return the node 9 to the parent node(5).<br>
```
                                       12
                            5                      15
                      3         x            13         17
                  1                9            14          20
                                8    11                  18
```
```
                                       12
                            5                      15
                      3         9            13         17
                  1           8    11           14          20
                                                        18
```
If we simply return the node 9 to the parent node(5), the tree is still a binary search tree.<br>
Therefore, when the node to delete has only one child, we simply return the non-nullptr child node to the parent node.<br>

Let's consider the third case.<br>
```
                                       12
                            5                       15  <--- deleted node
                      3         7             13         17
                  1                9             14          20
                                8     11                  18
```
If we want to delete node 15, what should we do?<br>
We have two options to replace the deleted node.<br>
1. Find the maximum value node in the left subtree.(In this case, the node with the value 14)
2. Find the minimum value node in the right subtree.(In this case, the node with the value 17)
For both of the options, we can make sure that the tree is still a binary search tree.<br>

We choose the second option to replace the deleted node.<br>
```
                                       12
                            5                       15  <--- deleted node
                      3         7             13         17  <--- node to replace the deleted node
                  1                9             14         20
                                8     11                 18
```

Now, we change the value of the node 15 to 17.<br>
```
                                       12
                            5                       17  <--- deleted node
                      3         7             13         17  <--- node to replace the deleted node
                  1                9             14         20
                                8     11                 18
```

Now, we try to delete the node 17.<br>
```
                                       12
                            5                       17  
                      3         7             13         17  <--- deleted node
                  1                9             14         20
                                8     11                 18
```
We simply recursively call the deleteNode function with the node 17, and repeat the same process.<br>
Because the root node is already at the node 17, we immediately find the node to delete.<br>
And because the node 17 only has one child(left subtree), we can simply return the 20 to the parent node(17).<br>
```
                                       12
                            5                       17  
                      3         7             13         20  
                  1                9             14    18         
                                8     11                    
```
After deleting the node 17, the tree is still a binary search tree.<br>


Let's summarize the above steps.<br>
1. Find the node to delete.
  - If the value to delete is less than the current node, go to the left subtree to find the node to delete.
  - If the value to delete is greater than the current node, go to the right subtree to find the node to delete.
2. If the node to delete is a leaf node, return nullptr to the parent node.
3. If the node to delete has only one child, return the non-nullptr child node to the parent node.
4. If the node to delete has two children, find the maximum value node in the left subtree or the minimum value node in the right subtree.
  - Change the value of the node to delete to the value of the node to replace the deleted node.
  - Delete the replaced node.

Let's finally walk through the above steps with the following tree.<br>
We want to delete the node 12
```
                                       12  <--- delete node
                            5                       15
                      3         7             13         17
                  1                9             14         20
                                8     11                 18
```
Because the value to delete is equal to the current node, we found the node to delete.<br>

If the node to delete has two children, find the maximum value node in the left subtree or the minimum value node in the right subtree.<br>
We choose to find the minimum value node in the right subtree.<br>
```
                                       12  <--- delete node
                            5                             15  
                      3         7             13 <-- replace node   17
                  1                9              14                    20
                                8     11                             18
```

Change the value of the node 12 to 13.<br>
```
                                       13  <--- delete node
                            5                             15  
                      3         7             13 <-- replace node   17
                  1                9              14                    20
                                8     11                             18
```

Repeat the same delete process on the right subtree of the delete node.<br>
```
                                       13  
                            5                       15  <-- root node (the root node of the right subtree of the delete node)
                      3         7             13         17
                  1                9             14          20
                                8     11                  18
```

13 is less then the root node 15, so we go to the left subtree to find the node to delete.<br>
```
                                       13  
                            5                              15  <-- root node (the root node of the right subtree of the delete node)
                      3         7               13 <-- delete node    17
                  1                9               14                    20
                                8     11                              18
```

The node 13 only has one child, so we simply return the node 14 to the parent node(15).<br>
```
                                       13  
                            5                            15  
                      3         7             x                       17
                  1                9              14                        20
                                8     11                                18
```
```
                                       13  
                            5                            15  
                      3         7             14                       17
                  1                9                                        20
                                8     11                                18
```
Done.<br>

We just simply compare the before and after to see if the tree is still a binary search tree.<br>
Before
```
                                       12  <--- delete node
                            5                       15
                      3         7             13         17
                  1                9             14          20
                                8     11                  18
```
After
```
                                       13
                            5                       15
                      3         7             14         17
                  1                9                         20
                                8     11                  18
```

10/27/2024 updated:<br>
Let's use the recursion idea to think about the problem.<br>
When using recursion, we always think about one idea, and three things.<br>
One idea is that "What problem we're solving for each call stack?".<br>
Three things are <br>
(1) What is the information we need to keep track of for each call stack?<br>
(2) What is the base case for the recursion?<br>
(3) What is the recursive case for the recursion?<br>

First idea:<br>
For each call stack, we're dealing with two cases.<br>
(1) Finding the node to delete.<br>
=> Is the current node the node to delete?<br>
=> If not, go to the left or right subtree to find the node to delete.<br>
(2) Deleting the node.<br>
=> If the node to delete is a leaf node, return nullptr to the parent node, which represents that the node is deleted.<br>
=> If the node to delete has only one child, return the non-nullptr child node to the parent node, which represents that the node is deleted.<br>
=> If the node to delete has two children, find the minimum value node in the right subtree, change the value of the node to delete to the value of the minimum value node, and delete the minimum value node.

Three things:<br>
(1) What is the information we need to keep track of for each call stack?<br>
=> The root node of the subtree that we're currently dealing with, and the value of the node to delete.
(2) What is the base case for the recursion?<br>
=> When the current node is nullptr, we return nullptr to the parent node.
(3) What is the recursive case for the recursion?<br>
=> When the current node is not nullptr, we compare the value of the current node with the value of the node to delete.<br>
=> If the current node is the node to delete, we deal with the three cases mentioned above.<br>
=> If the current node is not the node to delete, we go to the left or right subtree to find the node to delete.

# Complexity Analysis
## Time Complexity O(h)
- In the worst case, when the tree is highly unbalanced, the time complexity of the deleteNode function can be O(n), where n is the number of nodes in the tree.
- This worst-case scenario occurs when you have to traverse the entire height of the tree to find the node to delete.
- However, in average cases and balanced trees, the time complexity tends to be closer to O(log n), where n is the number of nodes.

## Space Complexity O(h)
- The space complexity of the deleteNode function is O(h), where h is the height of the tree.
- This space is due to the recursive calls made during the traversal down the tree.
- In the worst case, when the tree is highly unbalanced, the space complexity can be O(n), where n is the number of nodes (essentially turning into a linear linked list).
- In balanced trees or average cases, the space complexity is typically O(log n), corresponding to the height of a balanced binary search tree.
