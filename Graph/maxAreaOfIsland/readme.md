# Problem Explanation

This problem is quite simple.<br>
All we need to is following:<br>
1. Traverse the grid to find the starting point of the island.
2. Once we find the starting point, we will call a function to find the area of the island.
3. We will keep track of the maximum area of the island.

# Complexity Analysis
## Time Complexity O(N*M)
- where `N` is the number of rows and `M` is the number of columns in the grid.
- We are traversing the grid only once because of `seen` data structure.

## Space Complexity O(N*M)
- where `N` is the number of rows and `M` is the number of columns in the grid.
- We are using `seen` data structure to keep track of the visited nodes.
- And also we are using queue to keep track of the nodes to be visited.