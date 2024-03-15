# Problem Explanation

The core idea of solving this problem is truely understanding the definition of binary search tree (BST)

For any give node, there are only four possible scenarios
1. The value of node is equal to `p` or `q`
   - We immediately know the current node is the common ancestor of `p` and `q`
2. The value of node is in the range of `p` and `q`
   - The current node is the common ancestor of `p` and `q`
3. The value of node is smaller than `p` and `q`
   - The common ancestor of `p` and `q` is in the left subtree of the node
4. The value of node is greater than `p` and `q`
   - The common ancestor of `p` and `q` is in the right subtree of the node

Let's see the example
```
             6
          /     \
        2        8
      /   \    /   \
      0    4  7     9
          / \
         3   5
```
Suppose we're currently at node 6<br>

If p = 6 and q = 8<br>
It's the first case, the common ancestor of `p` and `q` is 6<br>

If p = 2 and q = 8<br>
It's the second case, the common ancestor of `p` and `q` is 6<br>
Why? <br>
Because 6 is in the range of [2, 8]<br>

If p = 2 and q = 4<br>
It's the third case, the common ancestor of `p` and `q` is in the left subtree of 6<br>
Why? <br>
Because both 2 and 4 are smaller than 6<br>

If p = 7 and q = 9<br>
It's the fourth case, the common ancestor of `p` and `q` is in the right subtree of 6<br>
Why? <br>
Because both 7 and 9 are greater than 6<br>

The above four cases are the core idea of solving this problem<br>

# Complexity Analysis
## Time Complexity: O(log(n))
- In every case, we're able to reduce the search space by half
- Simply put, we always go down one path of the tree
- We can also say the time complexity is O(h), where h is the height of the tree

## Space Complexity: O(log(n))
- The space complexity is O(h) due to the recursive stack