# Problem Explanation

This problem seems easy at first, but it's actually a little tricky

## Wrong Solution
At first, we might think we can use the following approach  
For each node, we basically just ask left and right subtree "What's the max path sum you got?"  
Once we get the answer from left and right subtree, we can compare different cases, and return the max value  
We might implement the code like this  
```go
func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxPathSum(root.Left)
    right := maxPathSum(root.Right)

    all := left + right + root.Val
    rootWithLeft := left + root.Val
    rootWithRight := right + root.Val

    return max(all, rootWithLeft, rootWithRight, root.Val, left, right)
}
```

However, this approach has a big problem, which is the answer we got is not the max path sum. It's just the max sum value in the tree node.  
Note that we're looking for the max path sum, which means that the max value has to be a PATH in the tree  
Let's see the following example: [1,-2,-3,1,3,-2,null]  
```
                1
        -2             -3
    1       3       -2
```
The wrong solution will return 4, but the actual answer is 3.  
The 4 comes from 1 + 3.  
Can you see the problem?  
1 + 3 is not a valid path!!!  

That's the problem  
We can't only return the global optimal solution to our parents  

## Key Insight: Two Different Values

The core insight is that at each node, we need to track **TWO different things**:

1. **Global Maximum**: Best path sum that passes THROUGH this node (for updating our final answer)
2. **Returnable Maximum**: Best path sum STARTING from this node going DOWN (what we return to parent)

These two values are different because:
- **Global maximum** can consider paths that go left → node → right (through the node)
- **Returnable maximum** can only consider paths that go in ONE direction (either left OR right, plus the node)

## Why We Need max(0, ...) Pattern

Before diving into the correct solution, there's an important optimization:

```go
left := max(0, helper(root.Left, ans))   // Ignore negative contributions
right := max(0, helper(root.Right, ans)) // Ignore negative contributions
```

We use `max(0, ...)` because:
- If a subtree gives us a negative path sum, we're better off ignoring it entirely
- A path with just the current node is better than current node + negative subtree
- This automatically handles the case where we might want just the single node

## Correct Solution
Let's see how to come up with the correct solution  
We use the same example: [1,-2,-3,1,3,-2,null]  
```
                1
        -2*             -3
    1       3       -2
```
Suppose we're at the position of -2(*), after applying the `max(0, ...)` pattern:
- left = max(0, 1) = 1
- right = max(0, 3) = 3

Now we have the following candidate values:

### For Global Maximum (what we track for final answer):
- `root + left + right`: -2 + 1 + 3 = 2 (path through this node)
- Previous global maximum

### For Returnable Maximum (what we return to parent):
- `root + left`: -2 + 1 = -1
- `root + right`: -2 + 3 = 1  
- `root` alone: -2

We return `max(root + left, root + right, root)` = max(-1, 1, -2) = 1

Let's see why we can't return other values:

**Why we can't return `left` (1) only:**  
If we return left(1) only, that means the path is skipping -2, we can't say a path is 1 → 1(root).

**Why we can't return `right` (3) only:**  
If we return right(3) only, that means the path is skipping -2, we can't say a path is 3 → 1(root).

**Why we can't return `left + root + right` (1-2+3):**  
If we return left+root+right(1-2+3), that means the path goes 1 → -2 → 3.  
But when the parent tries to extend this path, it would create an invalid tree path because you can't continue from both ends of 1 → -2 → 3.

So, the candidate values we can consider for returning to parent are `root+left`, `root+right`, and `root` alone.  
And we just return the max value.

The global optimal max path sum considers `root+left+right` because this represents a complete path that passes through the current node.


## Final Example Walkthrough

For the tree [1,-2,-3,1,3,-2,null]:
```
                1
        -2             -3
    1       3       -2
```

**Processing order (post-order):**
1. Node 1 (leftmost): ans = max(-∞, 0+0+1) = 1, returns 1
2. Node 3: ans = max(1, 0+0+3) = 3, returns 3  
3. Node -2: ans = max(3, 1+3+(-2)) = max(3, 2) = 3, returns max(1,3)+(-2) = 1
4. Node -2 (rightmost): ans = max(3, 0+0+(-2)) = 3, returns max(0,-2) = -2
5. Node -3: ans = max(3, max(0,-2)+0+(-3)) = 3, returns max(0,0)+(-3) = -3
6. Node 1 (root): ans = max(3, 1+max(0,-3)+1) = max(3, 2) = 3

**Final answer: 3** (representing the path 1 → 3)

## Complexity Analysis
### Time Complexity O(N)
- We visit each node once, so the time complexity is O(N)

### Space Complexity O(H)
- The space complexity is O(H), where H is the height of the tree. This is the space used by the recursion stack.