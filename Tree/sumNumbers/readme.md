# Problem Explanation

This problem is not that hard to solve as long as we understand the problem correctly<br>
It's similar to minDepth problem<br>
Because both of them ask us to "do something" for each leaf node<br>
leaf node is the key word<br>

For example,<br>
```
    3
   / \
  9  20
    /  \
   x    7
```
We can't calculate (3,20) because 20 is not a leaf node<br>
And this is the key point to solve this problem<br>

## Recursion
The idea is straightforward<br>
Along with the recursion, we just need to keep track of the current path<br>
If we reach the leaf node, we can calculate the sum of the path<br>

The logic is<br>
1. If the node is nullptr, return 0<br>
2. If the node is leaf node, return the sum of the path<br>
   - We can tell if the node is leaf node by checking if the right and left child is 0<br>
   - Because we return 0 at the base case<br>
3. Otherwise, we need to calculate the sum of the left and right subtree<br>

Let's walk through the example<br>
```
    3
   / \
  9   1
     /  \
    4    2
```
callstack 1
- root = 3
- path = "3"
- go left

callstack 2
- root = 9
- path = "3->9"
- it's a leaf node, return 39

callstack 1
- root = 3
- path = "3"
- left = 39
- go right

callstack 2
- root = 1
- path = "3->1"
- go left

callstack 3
- root = 4
- path = "3->1->4"
- it's a leaf node, return 314

callstack 4
- root = 1
- path = "3->1"
- left = 314
- go right

callstack 3
- root = 2
- path = "3->1->2"
- it's a leaf node, return 312

callstack 2
- root = 1
- path = "3->1"
- left = 314
- right = 312
- return 314 + 312 = 326

callstack 1
- root = 3
- path = "3"
- left = 39
- right = 326
- return 39 + 326 = 365

Done!!!<br>

In the first solution, we use string to keep track of the path<br>
At the leaf node, we simply iterate the string and calculate the sum<br>
But we can optimize this process<br>
Intead of using string, we can use integer to keep track of the path<br>
```
    3
   / \
  9   1
     /  \
    4    2
```
When we're at the root node, the path sum is 3<br>
When we're at node 9, the path sum is 3x10 + 9 = 39<br>
When we're at node 1, the path sum is 3x10 + 1 = 31<br>
When we're at node 4, the path sum is 31x10 + 4 = 314<br>
When we're at node 2, the path sum is 31x10 + 2 = 312<br>

Can you see the pattern?<br>
The path sum is the (previous path sum * 10) + (current node value)<br>

### Complexity Analysis
#### Time Complexity O(n)
- At worst, we have to visit all the nodes in the tree
- Note that it may take O(n ^ 2) if we use string to keep track of the path

#### Space Complexity O(n)
- At worst, the recursion callstack may take O(n) space
- And also we may use string to keep track of the path

## Iterative
In iterative solution, we use queue to do the breadth first search<br>
The idea is similar to the recursion<br>
But now we use struct to keep track the path information<br>
(string path or int sum)<br>

The logic is<br>
1. If the root is nullptr, return 0<br>
2. Initialize the queue with the root node and the path sum<br>
3. While the queue is not empty
   - Pop the front element
   - If it's a leaf node, calculate the sum
   - Otherwise, push the left and right child to the queue only if it's not nullptr

Let's walk through the example<br>
```
    3
   / \
  9   1
     /  \
    4    2
```
We use `int sum` as the path information<br>
queue = [{3,3}]<br>
- pop 3
- push 9, 1

queue = [{9,39}, {1,31}]<br>
- pop 9
- 9 is leaf node, calculate the sum (result += 39)
- pop 1
- push 4, 2

queue = [{4,314}, {2,312}]<br>
- pop 4
- 4 is leaf node, calculate the sum (result += 314)
- pop 2
- 2 is leaf node, calculate the sum (result += 312)

### Complexity Analysis
#### Time Complexity O(n)
- At worst, we have to visit all the nodes in the tree

#### Space Complexity O(n)
- At worst, the queue may take O(n) space


