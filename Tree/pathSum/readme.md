# Problem Explanation

The main hard part of this problem is how to handle the base case<br>
And also note that the problem is asking for a path from root to ***leaf***<br>

First, look at the wrong solution, and walk through the process<br>
```
                    5
              /           \
            4               8
          /                /  \
        11                13   4
```
callstack 1
- root = 5
- sum = 20
- path = []
- res = []
- push 5 to path
- go left

callstack 2
- root = 4
- sum = 15 (20 - 5)
- path = [5]
- res = []
- push 4 to path
- go left

callstack 3
- root = 11
- sum = 11 (15 - 4)
- path = [5, 4]
- res = []
- push 11 to path
- go left

callstack 4
- root = null
- sum = 0 (11 - 11)
- path = [5, 4, 11]
- res = []
- because sum is 0, push path to res
- back to callstack 3

callstack 3
- root = 11
- sum = 11 (15 - 4)
- path = [5, 4]
- res = [[5, 4, 11]]
- push 11 to path
- go left
- go right

callstack 4
- root = null
- sum = 0 (11 - 11)
- path = [5, 4, 11]
- res = [[5, 4, 11]]
- because sum is 0, push path to res
- back to callstack 3

callstack 3
- root = 11
- sum = 11
- path = [5, 4]
- res = [[5, 4, 11], [5, 4, 11]]
- push 11 to path
- go left
- go right

Can you see the problem?<br>
We push the two same paths to result because we don't handle the base case correctly<br>
Plus we handle the sum incorrectly<br>
And also we do not handle the leaf node logic<br>

For handling the sum incorrectly,<br>
We always only get the sum of previous callstack<br>
So that's why when we're at nullptr, we get the sum as 0<br>
And then we push the duplicate path to the result<br>

To handle correctly, we need to do three things<br>
1. We need to calculate the sum, and push the val to the path at the very beginning of the function (before the base case)<br>
  - For example, when we're at node 4, we need to calculate the sum as 20 - 4 = 16<br>
  - And push 4 to the path<br>
  - Now, we check if the sum is 0. If it is, we immediately find the path<br>
2. We need to handle the leaf node logic<br>
  - We only check the sum when we're at the leaf node<br>

Let's walk through the correct solution<br>
```
                    5
              /           \
            4               8
          /                /  \
        11                13   4
```
callstack 1
- root = 5
- sum = 20 - 5 = 15
- path = [5]
- res = []
- go left

callstack 2
- root = 4
- sum = 15 - 4 = 11
- path = [5, 4]
- res = []
- go left

callstack 3
- root = 11
- sum = 11 - 11 = 0
- path = [5, 4, 11]
- Because it's a leaf node, check if sum is 0
  - Yes, push path to res
- res = [[5, 4, 11]]
- back to callstack 2

callstack 2
- root = 4
- sum = 15 - 4 = 11
- path = [5, 4]
- res = [[5, 4, 11]]
- go left
- go right

callstack 3
- root = null
- immediately return

Look at the above process<br>
Because we handle the base case correctly, we only push the correct path to the result once<br>


# Complexity Analysis
## Time Complexity: O(n)
- In the worst case, we might need to visit every node in the tree to find all paths that sum up to the target sum

## Space Complexity: O(n + m)
- Therefore, the space complexity is O(n) for the recursion stack and O(m) for the curPath vector, where n is the number of nodes in the tree and m is the number of paths that sum up to the target sum

