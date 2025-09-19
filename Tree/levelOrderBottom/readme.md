# Problem Explanation

The problem is simple, and we have two different approaches.<br>

1. Using stack to store the nodes, then pop the nodes and append the values to the result list.
2. Find the maximum depth of the tree, then append the values to the result list.

For the first approach,<br>
We will use a stack to store the nodes. We will traverse the tree in a level order traversal and store the nodes in the stack. Then we will pop the nodes and append the values to the result list.<br>
In this way, we will get the level order traversal in reverse order.<br>

For the second approach,<br>
We will find the maximum depth of the tree. Then we will append the values to the result list in a level order traversal. We will start from the maximum depth and go to the root of the tree.
In this way, we will get the level order traversal in reverse order.<br>

# Complexity Analysis
## Time Complexity(O(n))
- For the first approach, we are traversing the tree in a level order traversal, so the time complexity will be O(n).
- For the second approach, we are finding the maximum depth of the tree, and also the second recursion to fill the result list, so the time complexity will be O(n).

## Space Complexity(O(n))
- For the first approach, we are using a stack and queue to store the nodes, so the space complexity will be O(n).
- For the second approach, we use recursion, so the space complexity will be O(n).

# Updated at 2025/09/19

## Regular Level Order Traversal
The idea of this approach is very straightforward<br>
We simply use regular level order traversal, then add each current level nodes at the beginning of the slice<br>
```
                3
          9          20
                15        7
```
First level: [3]<br>
-> Add to the beginning of the result: [[3]]<br>
Second level: [9,20]<br>
-> Add to the beginning of the result: [[3],[9,20]]<br>
Third level: [15,7]
-> Add to the beginning of the result: [[15,7],[9,20],[3]]<br>

## Using Stack
The idea is that we use regular level order traversal<br>
For each level nodes, we just add to stack<br>
In the end, we just pop the value from the stack, and add it to the final result<br>
```
                3
          9          20
                15        7
```
After the level order traversal, the stack becomes: [[3],[9,20],[15,7]]<br>
Then, pop the last element from the stack, and add it to the final result<br>
result: [[15,7],[9,20],[3]]

## Count the Length of Tree
The idea of this approach is to first count the max depth of the tree<br>
Once we know that, we can initialize the final result with max depth<br>
Then, we just use DFS to traverse the tree with level param, add each node to the correct position<br>
Note that we want the answer in reverse order(from bottom to top), so we directly pass (level - 1)<br>
```
                3
          9          20
                15        7
```
We know the max depth of tree is 3<br>
Initialize the final result: [[],[],[]]<br>

First Level:<br>
- node: 3
- level: 2
- result: [[],[],[3]]
Append the node value into 2th position

Second Level: 
- node: 9
- level: 1
- result: [[],[9],[3]]
Append the node value into 1th position

Second Level: 
- node: 20
- level: 1
- result: [[],[9,20],[3]]
Append the node value into 1th position

Third Level: 
- node: 15
- level: 0
- result: [[15],[9,20],[3]]
Append the node value into 0th position

Third Level: 
- node: 7
- level: 0
- result: [[15,7],[9,20],[3]]
Append the node value into 0th position