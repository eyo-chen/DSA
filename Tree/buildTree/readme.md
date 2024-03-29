# Problem Explanation

Let's walk through the thought process to solve this problem<br>

```
    3
   / \
  9  20
    /  \
   15   7
```
First look at the preorder traversal<br>
preorder traversal is [3,9,20,15,7]<br>
What does this traversal guarantee?<br>
It guarantees that ***the first element is the root of the tree***<br>
Because preorder is do, left, right<br>
But there's one problem<br>
We don't know how to construct the left and right subtree<br>
The tree may look like this<br>
```
    3
     \
      9
       \
       20
        \ 
        15
         \
          7
```

Then look at the inorder traversal<br>
inorder traversal is [9,3,15,20,7]<br>
What does this traversal guarantee?<br>
It guarantee ***how to construct the left and right subtree when given the node***<br>
Because inorder is left, do, right<br>
If we know that the root is 3<br>
We know that left subtree is [9] and right subtree is [15,20,7]<br>

Let's summarize the above two points<br>
1. Preorder traversal tells us the root of the tree
2. Inorder traversal tells us how to construct the left and right subtree when given the root

Let's see the example,
preorder = [3,9,20,15,7]<br>
inorder =  [9,3,15,20,7]<br>
We know that the root is 3<br>
So we can first construct the tree as following<br>
```
    3
  /   \ 
[9] [20,15,7]
```
Then we know that the left subtree is [9] and right subtree is [20,15,7]<br>
But now the next problem is how to keep finding the root node for each subtree<br>

This is the problem bother me when I try to solve this problem<br>
Because even though I can construct the first level of the tree<br>
I can't find the pattern to construct the entire tree recursively<br>

The key to solve this problem is that<br>
***preorder traversal tell us the order of the root of each subtree***<br>
Let's walk through<br>
preorder = [3,9,20,15,7]<br>
At first, we know that 3(index 0) is the root node<br>
So we construct the following tree<br>
```
    3
  /   \ 
[9] [20,15,7]
```
If we look at the left subtree first<br>
The root node is appraently 9(index 1)<br>
Now, we can go check right subtree<br>
The root node is 20(index 2)<br>
So we can construct the right subtree as following<br>
```
    3
  /   \ 
[9]   20
     /   \
   [15]  [7]
```
Again, check left first<br>
The root node is 15(index 3)<br>
Then check right<br>
The root node is 7(index 4)<br>

The key point is that<br>
As long as we follow the pattern of left first,<br>
We know that the order of root node of each subtree is the same as preorder traversal<br>

## Complexity Analysis
### Time Complexity: O(n)
- At worst, we have to visit all the nodes in the tree

### Space Complexity: O(n)
- Because we use hashmap to store the index of inorder traversal