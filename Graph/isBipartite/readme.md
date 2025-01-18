# Problem Explanation

The problem statement is a little bit confusing.<br>
The core idea is that ***we can only divide the graph into two groups, and the node and it's neighbors can't be in the same group.***

Once we understand the problem, the solution is straightforward.<br>
We can use DFS or BFS to traverse the graph and check if the node and it's neighbors are in the same group.<br>
The idea is that we just traverse all the nodes in the graph<br>
If the node is not visited, we start a traversal from that node.<br>
Before we start a traversal, we just give the node a group number.(In this case, we give it 1)<br>
In the traversal logic, the logic is following:
1. If the node is already visited, we check if it's the expected group.
2. If it's not the expected group, we immediately return false.
3. If it's the expected group, we continue the traversal.
4. We visite all the neighbors of the node, and give them the opposite group number.

In this problem, group number is just 1 and -1.(represent two groups)<br>
When we want to give a neighbor a group number, we just multiply the current group number by -1.<br>


# Complexity Analysis
## Time Complexity O(V+E)
- where V is the number of vertices and E is the number of edges.
- For either DFS or BFS, we traverse all the nodes and edges in the graph once.

## Space Complexity O(V)
- We use a hash table to store the visited nodes and their group numbers.
- The space complexity is O(V) because we need to store the visited nodes and their group numbers.
