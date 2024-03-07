# Problem Explanation

The problem becomes much easier if we can see the pattern to solve the problem<br>
Consider the following input<br>
<pre>
          1
        /   \
       2a    2b
     /  \   /   \
    3a  4a  4b   3b
   / \  /\  /\   / \
  5   6 7 8 8 7  6  5
</pre>
Let's see how to find the pattern<br>
At first level(root level), how do we know the next level of tree is symmetric?<br>
We just compare the left and right children of the root node<br>
Okay, let's move to the next level<br>

How can we make sure the third level is symmetric?<br>
2a.left == 2b.right<br>
2a.right == 2b.left<br>

How can we make sure the fourth level is symmetric?<br>
3a.left == 3b.right<br>
3a.right == 3b.left<br>
4a.left == 4b.right<br>
4a.right == 4b.left<br>

Can you see the pattern?<br>
For any given level (except the root level),<br>
- The left child of the left node should be equal to the right child of the right node<br>
- The right child of the left node should be equal to the left child of the right node<br>

For the root level, we just compare the left and right children<br>
Or even more simply, we just throw left and right to the helper function<br>

We can do this recursively or iteratively<br>

## Solution 1: Recursive
Let's start with the base case<br>
- If both nodes are null, return true<br>
- If one of the nodes is null, return false<br>
- If the values of the nodes are not equal, return false<br>

Then, we can recursively check the left and right children of the nodes<br>

Let's walk through the example (see the code for more details)<br>
<pre>
          1
        /   \
       2a    2b
     /  \   /   \
    3a  4a  4b   3b
</pre>

First call stack,<br>
node1 = 2a<br>
node2 = 2b<br>
No base case is met, so we continue to the next call stack<br>
helper(2a.left, 2b.right)<br>

Second call stack,<br>
node1 = 3a<br>
node2 = 3b<br>
No base case is met, so we continue to the next call stack<br>
helper(3a.left, 3b.right)<br>

Third call stack,<br>
node1 = nullptr<br>
node2 = nullptr<br>
Both nodes are nullptr, so we return true<br>

Now, we are back to the second call stack<br>
We have checked the first condition, so we continue to the next condition<br>
helper(3a.right, 3b.left)<br>

Third call stack,<br>
node1 = nullptr<br>
node2 = nullptr<br>
Both nodes are nullptr, so we return true<br>

Now, we are back to the second call stack<br>
We have checked the second condition, so we return true<br>

Now, we are back to the first call stack<br>
We have checked the first condition, so we continue to the next condition<br>
helper(2a.right, 2b.left)<br>
... same as above<br>

### Complexity Analysis
#### Time Complexity: O(n)
- The `helper` function is called once for each pair of nodes in the tree
- In the worst case, `helper` function might be called for every node in the tree, which happens when the tree is both full and symmetric
- Therefore, the time complexity of the solution is O(n), where n is the number of nodes in the binary tree, since each node is visited once

#### Space Complexity: O(log n) for a balanced tree, O(n) for a skewed tree
- The maximum depth of the recursion stack corresponds to the height of the tree
- For a balanced tree, the height is O(log n), leading to a space complexity of O(log n) because of the call stack during the recursion
- However, in the worst case, if the tree is skewed (e.g., each parent has only one child), the height can become O(n), which would also be the space complexity in that case due to the stack space needed for the recursion calls
- But if tree is skewed, won't we immediately return false? so the space complexity is still O(log n)?
  - Yes, if the tree is not symmetric, the function may return early without having to visit all nodes
  - However, we typically consider the worst-case scenario when analyzing space complexity
  - The worst-case scenario for a skewed tree would be if it is symmetric, or nearly symmetric, which would result in the deepest recursive calls going down to the last level of the tree
  - In such a case, the space complexity would indeed be O(n)
  - It is important to note that space complexity analysis often assumes the worst-case scenario to ensure that the algorithm can handle any input up to a certain size
  - Therefore, while the observation is valid in the context of early returns for certain inputs, the worst-case space complexity for a skewed tree remains O(n) due to the potential height of the recursion call stack in a symmetric or nearly symmetric skewed tree scenario


## Solution 2: Iterative
We can either use a stack or a queue to solve the problem iteratively<br>
In the code example, I use stack<br>
The core idea is the same as the recursive solution<br>

We create two stacks<br>
One stack is for the left subtree<br>
The other stack is for the right subtree<br>

We push the left and right children of the root node to the stacks<br>
Then, we pop the nodes from the stacks and compare the values<br>

We'll keep doing this until both the stacks are empty<br>
If one of the stack is nullptr, we'll catch that inside the while loop<br>

Let's walk through the example (see the code for more details)<br>
<pre>
          1
        /   \
       2a    2b
     /  \   /   \
    3a  4a  4b   3b
</pre>

Initially, the left stack has 2a, and the right stack has 2b<br>
stack1 = [2a]<br>
stack2 = [2b]<br>

First iteration,<br>
n1 = 2a<br>
n2 = 2b<br>
We compare the values, and they are equal<br>
stack1.push(2a->left)<br>
stack1.push(2a->right)<br>
stack2.push(2b->right)<br>
stack2.push(2b->left)<br>
stack1 = [3a, 4a]<br>
stack2 = [3b, 4b]<br>

Second iteration,<br>
n1 = 4a<br>
n2 = 4b<br>
We compare the values, and they are equal<br>
stack1 = [3a, nullptr, nullptr]<br>
stack2 = [3b, nullptr, nullptr]<br>

Third iteration,<br>
n1 = nullptr<br>
n2 = nullptr<br>
It's very important at this step<br>
In recursive, we return true at this condition<br>
But in iterative, we continue to the next iteration<br>
Because we haven't done the process<br>
stack1 = [3a, nullptr]<br>
stack2 = [3b, nullptr]<br>

Fourth iteration,<br>
n1 = nullptr<br>
n2 = nullptr<br>
stack1 = [3a]<br>
stack2 = [3b]<br>

Fifth iteration,<br>
n1 = 3a<br>
n2 = 3b<br>
We compare the values, and they are equal<br>
stack1 = [nullptr, nullptr]<br>
stack2 = [nullptr, nullptr]<br>

Now, we just keep popping the nodes from the stacks<br>
Because they all are nullptr, we keep continue<br>

Breaking the while loop, we return true<br>

### Complexity Analysis

#### Time Complexity: O(n)
- Similar to the recursive solution, this iterative approach examines each node in the tree once
- For every node, we're performing a constant number of operations (pushing and popping from the stack, and checking values)
- Therefore, the time complexity remains O(n), where n is the number of nodes in the binary tree


#### Space Complexity: O(n)
- In terms of space complexity, the algorithm uses two stacks to keep track of the nodes to visit
- At most, the stacks will hold all the nodes at the largest width of the tree
- This occurs at the level with the most nodes, which, for a perfectly balanced binary tree, is at the last level with up to n/2 nodes (half of the total nodes in the tree)
- Therefore, the space complexity for a balanced tree is O(n) since in the worst-case scenario, the last level of the tree can contain at most half of the nodes