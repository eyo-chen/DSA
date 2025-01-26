# Problem Explanation

The key to approach this problem is to figure out the different scenarios for each node.<br>

For any given node, there are only four possible scenarios
1. The node is equal to `p` or `q`
2. `p` and `q` are in my children(might in my left subtree or right subtree)
3. Only my left subtree has `p` or `q`
4. Only my right subtree has `p` or `q`

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
1. Suppose we're at node 4, p = 4 and q = 5
   - When we're at node 4, we know that node 4 is equal to `p`
   - There are only two cases
     - (1) ***node 4 is the common ancestor of `p` and `q`***
     - (2) ***node 4's parent is the common ancestor of `p` and `q`***
   - This is the key point, because in either case, we just direcly return the node 4
   - It's like saying, "Hey, my caller, I've found either `p` or `q`."
   - Then node 2 will return node 4 because it can't find target node in its left subtree
   - Then node 6 will return node 4 because it can't find target node in its right subtree
   - Then the caller will return node 4

2. Suppose we're at node 4, p = 3 and q = 5
   - When we're at node 4, my left and right child node tell me, "Hey, we both found `p` or `q`"
   - That means, I'm the common ancestor of `p` and `q`
   - So I return myself

3. Suppose we're at node 2, p = 4 and q = 5
   - When we're at node 2, my right child node tell me, "Hey, I found `p` or `q`"
   - However, my left child node tell me, "Hey, I didn't find `p` or `q`"
   - That means, the common ancestor has to be in my right subtree, just return my right child node

That's the core idea of this problem<br>
1. If I'm equal to `nil`, `p` or `q`, return myself(base case)
2. Go to my left and right subtree, and I'll do the further operation once I get the result from my children
3. If my left and right child both found `p` or `q`, return myself
4. If only my left child found `p` or `q`, return my left child
5. If only my right child found `p` or `q`, return my right child

# Complexity Analysis
## Time Complexity O(n)
- The algorithm needs to visit each node in the binary tree exactly once to find the lowest common ancestor

## Space Complexity O(h)
- The space complexity is O(h) where h is the height of the binary tree
- This is because the maximum number of recursive calls on the call stack is equal to the height of the tree
- In the worst case, the tree is skewed and the height is O(n), so the space complexity is O(n)
