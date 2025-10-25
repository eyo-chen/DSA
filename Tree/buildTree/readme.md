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

## Updated At 2025/09/13
It's pretty easy to understand that `preorder` tells us the order of the root, and `inorder` tells us how to split right subtree and left subtree.<br>
However, the actual implementation is actually a little bit tricky<br>
The tricky part is how to correctly find the root using index<br>

Let's see this example<br>
```
             3
        /         \ 
       9          20
     /   \       /   \
   10     5     15    7
```
preorder=[3,9,10,5,20,15,7], inorder=[10,9,5,3,15,20,7]<br>
It's obvious that the 3 is the first root node<br>
Then, we attempt to construct the left subtree. It's obvious that we just increment the index of preorder to find the next root<br>
The tricky part is for the right subtree.<br>
Suppose we're at the first callstack where we've constructed the left subtree<br>
```
index_of_preorder = 0
root = 3 = preorder[0]
left = 9 = preorder[0+1] (suppose the entire left subtree is constructed)
right=?
```
How do we update the index to find the correct root node of right subtree.<br>
The intuition is that we want to ***skip the number of node in the left subtree***<br>
Let's see the diagram again
```
             3
        /         \ 
       9           ?
     /   \       
   10     5

         idx <- current index position
preorder=[3,9,10,5,20,15,7]
```
We know there are 3 nodes in the left subtree, that means we need to update the index to `idx+3+1` to find the root node of right subtree.<br>
The `+1` logic is to skip the root node itself<br>

Now, the question becomes how to know the number of node in the left subtree?<br>
We can utilize inorder array<br>
inorder=[10,9,5,3,15,20,7]<br>
After we finding the root node 3, we can conceptually split the inorder into two halves<br>
- left subtree = [10,9,5]
- right subtree = [15,20,7]
And the index of root node in inorder array is 3.<br>
That means there are three nodes in the left subtree<br>

Dependes on how we split the inorder array, there are different approaches to find the number of node in the left subtree<br>
1. Actually slices the array(the first Go solution)
In this case, we just update the index `idx+1+index_of_root_node`.

2. We only use left, right pointer to split the inorder array(Do not actually slice the array) (The second Go solution)
In this case, we need to update the index `idx+1+index_of_root_node-left`<br>
Why it's differenct?<br>
Note that our goal is to skip the number of node in the left subtree + root node itself.<br>
Because we do not slice the array, so `index_of_root_node` won't really represent the number of node in left subtree.<br>
```
             3
        /         \ 
       9          20
     /   \       /   \
   10     5     15    7
```
- When we want to find the correct index for 20: `idx = idx+3+1-0`
  - idx=0 (starting point)
  - preorder=[3,9,10,5,20,15,7]
  - root_node=preorder[idx]=3
  - inorder=[10,9,5,3,15,20,7]
  - left=0, right=7
  - index_of_root_node=3
  - number of node in left subtree=3
- When we want to find the correct index for 7: `idx+5+1-4`
  - idx=4
  - preorder=[3,9,10,5,20,15,7]
  - root_node=preorder[idx]=20
  - inorder=[10,9,5,3,15,20,7]
  - left=4, right=7
  - index_of_root_node=5
  - number of node in left subtree=1
Note that there's only one node in the left subtree for the tree=[15,20,7]<br>
The `index_of_root_node` is 5, but this 5 is the absolute index in the original inorder array, which is not right<br>
We need to subtract the left value to simulate that we're slicing the array.


## Complexity Analysis
### Time Complexity: O(n)
- At worst, we have to visit all the nodes in the tree

### Space Complexity: O(n)
- Because we use hashmap to store the index of inorder traversal