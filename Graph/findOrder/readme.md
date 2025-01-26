# Problem Explanation

The idea to solve this problem is to use topological sort.<br>
Topological sort is a linear ordering of vertices such that for every directed edge uv from vertex u to vertex v, u comes before v in the ordering.<br>
For example,<br>
```
0 -> 1 -> 3
0 -> 2 -> 3
```
The topological sort of this graph is [0, 1, 2, 3]<br>

There are two ways to solve this problem:
1. Use BFS to find the topological sort
2. Use DFS to find the topological sort

# Solution 1: Use BFS to find the topological sort
It's similar to the typical BFS approach, where we use a queue to traverse the graph.<br>
However, there's one thing one need to keep in mind:<br>
***We have to keep track of the number of dependencies for each course, and only add the course to the queue if it has no dependencies.***<br>

At the beginning, we need to build the adjacency list and the dependency count for each course.<br>
Then, we need to find all the courses that have no dependencies, and add them to the queue.<br>
Which means that we can only start processing the courses that have no dependencies.<br>
After that, we need to traverse the graph using BFS, and add the course to the result list if it has no dependencies.<br>
Finally, we need to check if the result list has the same length as the number of courses, if so, return the result list, otherwise, return an empty list.<br>

This part of code is important<br>
```go
		for _, neighbor := range adj[node] {
			dependency[neighbor]--
			if dependency[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
```
We can only add the course to the queue if it has no dependencies, which means that we can only start processing the courses that have no dependencies.<br>
Let's consider the following example:<br>
```
0 -> 1 -> 3 -> 4 -> 5
0 -> 2 -> 5
```
suppose the current node is 2, and the current builded answer is [0, 1, 2]<br>
So, now we check all the neighbors of node 2, which is node 5<br>
Suppose we directly add node 5 to the queue, then we will have a problem<br>
We might add node 5 to the answer before node 4, which is incorrect<br>
So, we can ***only process node 5 if we finish processing node 4 and node 2***<br>
Every time we process a node, we need to decrease the dependency count of its neighbors, and if the dependency count of a neighbor is 0, we add it to the queue.<br>

# Solution 2: Use DFS to find the topological sort
When using DFS, it's very similar to the Course Schedule I problem, where we need to check if there's a cycle in the graph.<br>
If there's a cycle, we return an empty list, otherwise, we return the order of the topological sort.<br>

There's only one thing we need to keep in mind:<br>
When do we add the course to the answer list?<br>
We can only add the course to the answer list after we finish processing all of its neighbors.<br>
This is because we need to make sure that the course is added to the answer list in the correct order.<br>

Let's suppose we add the course to the answer list before we finish processing all of its neighbors.<br>
```
0 -> 1 -> 3 -> 4 -> 5
0 -> 2 -> 5
```
We first visit the path 0 -> 1 -> 3 -> 4 -> 5<br>
So our current answer is [0, 1, 3, 4, 5]<br>
Then we visit the path 0 -> 2 -> 5<br>
So our current answer is [0, 1, 3, 4, 5, 2]<br>
As we can see, the answer is not correct, because we added the course to the answer list before we finish processing all of its neighbors.<br>

So, the correct way to do it is to add the course to the answer list after we finish processing all of its neighbors.<br>
We first vising the path 0 -> 1 -> 3 -> 4 -> 5<br>
After processing all the neighbors of node 5, we add node 5 to the answer list, ans = [5]<br>
After processing all the neighbors of node 4, we add node 4 to the answer list, ans = [5, 4]<br>
After processing all the neighbors of node 3, we add node 3 to the answer list, ans = [5, 4, 3]<br>
After processing all the neighbors of node 1, we add node 1 to the answer list, ans = [5, 4, 3, 1]<br>

Then, we visit the path 0 -> 2 -> 5<br>
Don't need to visit node 5 because we've already visited it.<br>
After processing all the neighbors of node 2, we add node 2 to the answer list, ans = [5, 4, 3, 1, 2]<br>
After processing all the neighbors of node 0, we add node 0 to the answer list, ans = [5, 4, 3, 1, 2, 0]<br>

Finally, we just reverse the answer list, ans = [0, 2, 1, 3, 4, 5]<br>

That's the core logic of this approach.<br>
Note that the `hasCycle` function does two things:<br>
1. Check if there's a cycle in the graph
2. Build up the answer in reverse order