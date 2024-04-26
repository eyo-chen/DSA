# Problem Explanation

The key idea to solve this problem is <br>
***For any given node, we need to answer: Are you balanced?*** <br>
To answer that problem, we need to know the following information: <br>
1. Is the left subtree balanced?
2. Is the right subtree balanced?
3. Is the abs(left subtree height - right subtree height) <= 1?

If any of the above conditions is false, then the current node is not balanced. <br>
So we can immediately tell the parent node that I'm not balanced. <br>

Above is the key point to solve this problem recursively. <br>

Now, the problem is how to design the recursive function. <br>
Because we have to know three information for each node, but the function can only return one value in c++. <br>
We can just make the recursive function return the height(int) of the current node. <br>
We invoke the recursive function on left and right subtree to get the height<br>
Then, we can know if the current node is balanced or not. <br>
If the current node is balanced, we return the height(max(left, right) + 1) of the current node. <br>
If the current node is not balanced, we return -1. <br>
-1 means the current node is not balanced. <br>
So, whenever we get -1 from the left or right subtree, we can immediately return -1 to the parent node. <br>
Also, we return 0 to the parent node if the current node is a leaf node. (base case)<br>

Note that
```c++
int left = helper(root->left);
if (left < 0) return -1;
```
The condition right after the recursive function is very important. <br>
If the left subtree is not balanced, we can immediately return -1 to the parent node. <br>
We don't need to check the right subtree. <br>

# Complexity Analysis
## Time Complexity O(N)
- We visit each node once. So the time complexity is O(N).

## Space Complexity O(H)
- The space complexity is O(H), where H is the height of the tree.
- The space complexity is O(logN) for a balanced tree.
- The space complexity is O(N) for a skewed tree.